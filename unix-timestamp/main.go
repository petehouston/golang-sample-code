package main

import (
	"fmt"
	"time"
)

func main() {
	fromTime := time.Now()

	ts := fromTime.Unix()

	fmt.Printf("Unix timestamp: %d\n", ts)
}
