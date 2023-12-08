package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, World!")
}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Sort(RidesSlice(rides))
	notRideMax := make([]int64, n+1)
	rideIndex := 0
	i := 1
outLoop:
	for ; i <= n; i++ {
		notRideMax[i] = notRideMax[i-1]
		for rides[rideIndex][1] <= i {
			rideStartIndex := rides[rideIndex][0]
			rideStopIndex := rides[rideIndex][1]
			gain := (int64)(rideStopIndex - rides[rideIndex][0] + rides[rideIndex][2])
			thisRideResult := gain + notRideMax[rideStartIndex]
			notRideMax[rideStopIndex] = max(notRideMax[rideStopIndex], thisRideResult)
			rideIndex++
			if rideIndex >= len(rides) {
				i++
				break outLoop
			}
		}
	}
	return notRideMax[i-1]
}

type RidesSlice [][]int

func (p RidesSlice) Len() int           { return len(p) }
func (p RidesSlice) Less(i, j int) bool { return p[i][1] < p[j][1] }
func (p RidesSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
