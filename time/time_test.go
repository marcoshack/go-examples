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
