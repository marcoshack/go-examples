package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/marcoshack/go-examples/concurrency"
)

func main() {
	printer := message.NewPrinter(language.English)
	if len(os.Args) < 5 {
		printer.Println("usage: concurrency <numConcurrentReaders> <durationInSeconds> <reloadDataFrequencyInMillisecond> <b[locking]|n[onblocking]")
		os.Exit(1)
	}

	numConcurrentReaders, err := strconv.Atoi(os.Args[1])
	if err != nil || numConcurrentReaders == 0 {
		printer.Printf("invalid number of concurrent readers: '%d'\n", os.Args[1])
		os.Exit(1)
	}

	durationInSeconds, err := strconv.Atoi(os.Args[2])
	if err != nil || durationInSeconds == 0 {
		printer.Printf("invalid duration: '%s'\n", os.Args[2])
		os.Exit(1)
	}

	reloadFrequencyInMillisecond, err := strconv.Atoi(os.Args[3])
	if err != nil || reloadFrequencyInMillisecond == 0 {
		printer.Printf("invalid reload frequency: '%s'\n", os.Args[3])
		os.Exit(1)
	}

	var safeData concurrency.SafeData
	blockingMode := os.Args[4]
	switch blockingMode {
	case "b":
		safeData = concurrency.NewSafeDataBlockingRead(genereateData())
	case "n":
		safeData = concurrency.NewSafeDataNonBlockingRead(genereateData())
	default:
		printer.Printf("invalid blocking mode: '%s'\n", os.Args[4])
		os.Exit(1)
	}

	go reloadDataPeriodically(&reloadDataPeriodicallyInput{
		Frequency:    time.Duration(reloadFrequencyInMillisecond) * time.Millisecond,
		InitialDelay: 0,
		Data:         safeData,
		Printer:      printer,
	})

	var wg sync.WaitGroup
	for i := 1; i <= numConcurrentReaders; i++ {
		wg.Add(1)
		go printDataPeriodically(&printDataPeriodicallyInput{
			RoutineID: i,
			Duration:  time.Duration(durationInSeconds) * time.Second,
			Data:      safeData,
			WaitGroup: &wg,
			Printer:   printer,
		})
	}

	wg.Wait()
}

type printDataPeriodicallyInput struct {
	RoutineID int
	Duration  time.Duration
	Data      concurrency.SafeData
	WaitGroup *sync.WaitGroup
	Printer   *message.Printer
}

func printDataPeriodically(input *printDataPeriodicallyInput) {
	defer input.WaitGroup.Done()
	stopTime := time.Now().Add(input.Duration)
	attr4total := int64(0)
	attr5total := float64(0)
	count := int64(0)
	for {
		count++
		d := input.Data.GetData()
		attr4total += d.Attr4
		attr5total += d.Attr5
		if time.Now().After(stopTime) {
			break
		}
	}

	readsPerSecond := count / int64(input.Duration.Seconds())
	input.Printer.Printf("reader[%03d]: totalReads = %d, avgReads/sec: %d\n", input.RoutineID, count, readsPerSecond)
}

type reloadDataPeriodicallyInput struct {
	Frequency    time.Duration
	InitialDelay time.Duration
	Data         concurrency.SafeData
	Printer      *message.Printer
}

func reloadDataPeriodically(input *reloadDataPeriodicallyInput) {
	if input.InitialDelay > 0 {
		time.Sleep(input.InitialDelay)
	}
	count := 0
	for {
		count++
		input.Printer.Printf("writer: relading data (count = %d, frequency = %s)\n", count, input.Frequency)
		input.Data.SetData(genereateData())
		time.Sleep(input.Frequency)
	}
}

func genereateData() *concurrency.Data {
	return &concurrency.Data{
		Attr1: fmt.Sprintf("value%d", rand.Int()%10),
		Attr2: fmt.Sprintf("value%d", rand.Int()%10),
		Attr3: fmt.Sprintf("value%d", rand.Int()%10),
		Attr4: rand.Int63(),
		Attr5: rand.Float64(),
	}
}
