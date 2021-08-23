package main

import (
	fib "dynamicProgramming/fibonacci"
	"dynamicProgramming/game"
	gridtravel "dynamicProgramming/gridTravel"
	sum "dynamicProgramming/sum"
	wc "dynamicProgramming/wordConstruct"
	"fmt"
)

func main() {
	/* memo := map[string]bool{}
	fmt.Println(wc.CanConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, memo))
	memo = map[string]bool{}
	fmt.Println(wc.CanConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, memo))
	fmt.Println(memo)
	memoCount := map[string]int{}
	fmt.Println(wc.CountConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, memoCount))
	fmt.Println(memoCount)
	memoCount = map[string]int{}
	fmt.Println(wc.CountConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, memoCount))
	fmt.Println(memoCount) */

	memoAll := map[string][][]string{}
	fmt.Println(wc.AllConstructRec("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}, memoAll))
	/*memoAll = map[string][][]string{}
	fmt.Println(wc.AllConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, memoAll)) */
	//fmt.Println(memoAll)
	fmt.Println(fib.FibIter(100))
	/* memoFib := map[int]int{}
	fmt.Println(fib.FibRec(100, memoFib)) */
	fmt.Println(gridtravel.GridTravelIter(18, 18))
	fmt.Println(sum.CanSumIter(300, []int{7, 14}))
	/* memoSum := map[int]bool{}
	fmt.Println(sum.CanSumRec(300, []int{7, 14}, memoSum)) */
	fmt.Println(sum.HowSumIter(7, []int{5, 3, 4}))
	memoSum := map[int][][]int{}
	fmt.Println(sum.HowSumRec(7, []int{5, 3, 4}, memoSum))
	memoSum = map[int][][]int{}
	fmt.Println(sum.HowSumRec(8, []int{2, 3, 5}, memoSum))
	fmt.Println(sum.CanSumTwo([]int{0, 0, -5, 30212}, []int{-10, 40, -3, 9}, -8))
	memoBestSum := map[int][]int{}
	fmt.Println(sum.BestSumRec(8, []int{2, 3, 5}, memoBestSum))
	memoBestSum = map[int][]int{}
	fmt.Println(sum.BestSumRec(8, []int{4, 3, 1}, memoBestSum))
	fmt.Println(sum.BestSumIter(8, []int{4, 3, 1}))
	fmt.Println(sum.BestSumIter(100, []int{1, 2, 5, 25}))
	fmt.Println(game.MaxDifference(4, []int{-1, 100, 4, -5}))
}
