package main

import (
	"fmt"
	"strconv"
)

type KeyValue struct {
	Key string
	Value string
}

func main() {
	size := 10
	createWithAppend(size)
	createWithIndex(size)
}

func createWithAppend(size int) {
	fmt.Println(">>> createWithAppend")

	mySlice := make([]KeyValue, 0, size);
	fmt.Println("initial size", len(mySlice))

	// slice was initialized with size=0 and capacity=10, so no default items will be initialized and
	// appending 10 items won't crate a new slice
	for i := 0; i < size; i++ {
		mySlice = append(mySlice, KeyValue{Key: "key" + strconv.Itoa(i), Value: "value" + strconv.Itoa(i)})
	}

	fmt.Println("size after adding", len(mySlice))
	fmt.Println("content", mySlice)
}

func createWithIndex(size int) {
	fmt.Println(">>> createWithIndex")

	mySlice := make([]KeyValue, size);
	fmt.Println("initial size", len(mySlice))
	fmt.Println("empty elements are created, mySlice[0].Key=", mySlice[0].Key)

	
	for i := 0; i < size; i++ {
		mySlice[i] = KeyValue{Key: strconv.Itoa(i), Value: "value of " + strconv.Itoa(i)}
	}

	fmt.Println("size after adding", len(mySlice))
	fmt.Println("content", mySlice)
}