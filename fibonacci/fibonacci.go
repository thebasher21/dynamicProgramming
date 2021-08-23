package fibonacci

func FibIter(n int) (ans int) {
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 0; i < n; i++ {
		arr[i+1] += arr[i]
		if i+2 <= n {
			arr[i+2] += arr[i]
		}
	}
	ans = arr[n]
	return ans
}

func FibRec(n int, memo map[int]int) int {
	if n == 0 {
		return 0
	} else if n < 2 {
		return 1
	}
	returnObj, isPresent := memo[n]
	if isPresent {
		return returnObj
	} else {
		memo[n] = FibRec(n-1, memo) + FibRec(n-2, memo)
		return memo[n]
	}
}
