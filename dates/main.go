package main

import (
	"fmt"
	"time"
)

func main() {
	date, _ := time.Parse(time.RFC3339, "2020-06-03T16:38:39.696-07:00")
	fmt.Printf("Parsed date: %s\n", date)
	fmt.Printf("Unix(): %d\n", date.Unix())
}
