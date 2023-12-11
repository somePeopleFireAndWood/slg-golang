package main

import (
	"sort"
)

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	sort.Sort(IntSplice(candidates))
	return calcCombination(0, candidates, target)
}

func calcCombination(candidateIndex int, candidates []int, target int) [][]int {
	resultSplice := [][]int{}
	l := len(candidates)
	for i := candidateIndex; i < l; i++ {
		candiate := candidates[i]
		if candiate > target {
			break
		}
		if candiate == target {
			singleCombination := []int{candiate}
			resultSplice = append(resultSplice, singleCombination)
			break
		}
		sonCombinations := calcCombination(i, candidates, target-candiate)
		for _, singleSonCombination := range sonCombinations {
			singleCombination := []int{candiate}
			singleCombination = append(singleCombination, singleSonCombination...)
			resultSplice = append(resultSplice, singleCombination)
		}
	}
	return resultSplice
}

type IntSplice []int

func (h IntSplice) Len() int           { return len(h) }
func (h IntSplice) Less(i, j int) bool { return h[i] < h[j] }
func (h IntSplice) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
