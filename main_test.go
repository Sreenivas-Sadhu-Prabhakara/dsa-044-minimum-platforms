package main

import "testing"

func TestMinPlatforms(t *testing.T) {
	cases := []struct {
		arr, dep []int
		want     int
	}{
		{[]int{900, 940, 950, 1100, 1500, 1800}, []int{910, 1200, 1120, 1130, 1900, 2000}, 3},
		{[]int{900, 1100, 1235}, []int{1000, 1200, 1240}, 1},
		{[]int{1000, 1000}, []int{1030, 1030}, 2},
		{[]int{900}, []int{910}, 1},
		{[]int{}, []int{}, 0},
		{[]int{900, 910}, []int{910, 920}, 2}, // arrival at the exact departure minute still overlaps
	}
	for _, c := range cases {
		if got := MinPlatforms(c.arr, c.dep); got != c.want {
			t.Errorf("MinPlatforms(%v,%v) = %d, want %d", c.arr, c.dep, got, c.want)
		}
	}
}
