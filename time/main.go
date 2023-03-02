package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("Now: %s\n", time.Now().Format(time.RFC3339))
	os.Exit(0)
}
