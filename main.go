package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now().Add(1 * time.Hour)
	fmt.Printf("Hello, world! %s\n", date.Format(time.UnixDate))
	if (date.Before(time.Now().Add(5 * time.Second))) {
		fmt.Println("Expired")
	} else {
		fmt.Println("Not expired")
	}
}
