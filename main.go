package main

import (
	"fmt"
	"sort"
)

// MinPlatforms returns the fewest railway platforms needed so no train
// waits: sort arrivals and departures separately and sweep — the peak
// number of simultaneously present trains is the answer.
func MinPlatforms(arrivals, departures []int) int {
	arr := append([]int{}, arrivals...)
	dep := append([]int{}, departures...)
	sort.Ints(arr)
	sort.Ints(dep)
	platforms, peak := 0, 0
	i, j := 0, 0
	for i < len(arr) {
		if arr[i] <= dep[j] {
			platforms++
			if platforms > peak {
				peak = platforms
			}
			i++
		} else {
			platforms--
			j++
		}
	}
	return peak
}

func main() {
	fmt.Println(MinPlatforms(
		[]int{900, 940, 950, 1100, 1500, 1800},
		[]int{910, 1200, 1120, 1130, 1900, 2000},
	))
}
