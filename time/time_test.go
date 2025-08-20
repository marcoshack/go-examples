package time

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTime_Epoch(t *testing.T) {
	epoch := time.Now().Unix()
	epochStr := strconv.Itoa(int(epoch))
	fmt.Printf("epoch: %d (%d)", epoch, len(epochStr))
}

func TestTime_Duration(t *testing.T) {
	numberOfSeconds := int64(30)
	duration := time.Duration(numberOfSeconds) * time.Second

	fmt.Printf("duration: %s\n", duration.String())
}
