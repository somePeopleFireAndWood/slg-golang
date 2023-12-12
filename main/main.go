package main

import (
	"sort"
)

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

func secondGreaterElement(nums []int) []int {
	ans := make([]int, len(nums))
	firstStack := []int{}
	secondStack := []int{}
	for i, num := range nums {
		toIndex := len(secondStack) - 1
		for ; toIndex >= 0; toIndex-- {
			if nums[secondStack[toIndex]] < num {
				ans[secondStack[toIndex]] = num
			} else {
				break
			}
		}
		secondStack = secondStack[:toIndex+1]
		toIndex = len(firstStack) - 1
		for ; toIndex >= 0; toIndex-- {
			if nums[firstStack[toIndex]] >= num {
				break
			}
		}
		secondStack = append(secondStack, firstStack[toIndex+1:]...)
		firstStack = firstStack[:toIndex+1]
		firstStack = append(firstStack, i)
	}
	for i := range secondStack {
		ans[secondStack[i]] = -1
	}
	for i := range firstStack {
		ans[firstStack[i]] = -1
	}
	return ans
}
