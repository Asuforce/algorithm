package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	n, _ := strconv.Atoi(args[0])
	low, _ := strconv.Atoi(args[1])
	high, _ := strconv.Atoi(args[2])
	result := turns(n, low, high)

	fmt.Printf("turns is %d", result)
}

func turns(n int, low int, high int) int {
	var turns = 0

	for high >= low {
		turns++
		var mid = (low + high) / 2
		if mid == n {
			return turns
		} else if mid < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return turns
}
