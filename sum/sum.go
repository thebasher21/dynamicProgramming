package cansum

func CanSumIter(n int, numBank []int) bool {
	result := make([]bool, n+1)
	result[0] = true
	for _, i := range numBank {
		result[i] = true
		for _, j := range numBank {
			if i+j <= n {
				result[i+j] = true
			}

		}
	}
	return result[n]
}

func CanSumRec(n int, numBank []int, memo map[int]bool) bool {
	returnObj, isPresent := memo[n]
	if isPresent {
		return returnObj
	}
	if n == 0 {
		return true
	} else if n < 0 {
		return false
	}
	for _, num := range numBank {
		newNum := n - num
		memo[n] = CanSumRec(newNum, numBank, memo)
		if memo[n] {
			return true
		}
	}
	memo[n] = false
	return false
}

func HowSumIter(n int, numBank []int) []int {
	result := make([][]int, n+1)
	result[0] = []int{}
	for i := 0; i <= n; i++ {
		if result[i] != nil {
			for _, j := range numBank {
				if i+j <= n {
					result[i+j] = append([]int{j}, result[i+j]...)
				}
			}
		}
	}
	return result[n]
}

func HowSumRec(n int, numBank []int, memo map[int][][]int) [][]int {
	returnObj, isPresent := memo[n]
	if isPresent {
		return returnObj
	}
	if n == 0 {
		return [][]int{{}}
	}
	if n < 0 {
		return nil
	}
	for _, num := range numBank {
		remainder := n - num
		ways := HowSumRec(remainder, numBank, memo)
		for index := range ways {
			memo[n] = append(memo[n], append([]int{num}, ways[index]...))
		}
		if memo[n] != nil {
			return memo[n]
		}

	}
	memo[n] = nil
	return nil
}

func BestSumRec(n int, numBank []int, memo map[int][]int) []int {
	returnObj, isPresent := memo[n]
	if isPresent {
		return returnObj
	}
	if n == 0 {
		return []int{}
	}
	if n < 0 {
		return nil
	}

	var shortestCombo []int

	for _, num := range numBank {
		remainder := n - num
		remainderCombo := BestSumRec(remainder, numBank, memo)
		if remainderCombo != nil {
			combo := append([]int{num}, remainderCombo...)
			if shortestCombo == nil || len(shortestCombo) > len(combo) {
				shortestCombo = combo
			}
		}
	}
	if shortestCombo == nil {
		memo[n] = nil
		return nil
	}
	memo[n] = shortestCombo
	return shortestCombo
}

func BestSumIter(n int, numBank []int) []int {
	result := make([][]int, n+1)
	for i := range result {
		if i == 0 {
			result[i] = []int{}
		} else {
			result[i] = nil
		}
	}
	for i := 0; i <= n; i++ {
		if result[i] != nil {
			for _, j := range numBank {
				if i+j <= n {
					combo := append([]int{j}, result[i]...)
					if result[i+j] == nil || len(result[i+j]) > len(combo) {
						result[i+j] = combo
					}
				}
			}
		}
	}
	return result[n]
}

func CanSumTwo(arr1, arr2 []int, v int) bool {
	seen := map[int]bool{}
	switch len(arr2) > len(arr1) {
	case true:
		for i := range arr1 {
			_, isPresent := seen[arr1[i]]
			if !isPresent {
				seen[arr1[i]] = true
			}
		}
		for i := range arr2 {
			_, isPresent := seen[v-arr2[i]]
			if isPresent {
				return true
			}
		}
	case false:
		for i := range arr2 {
			_, isPresent := seen[arr2[i]]
			if !isPresent {
				seen[arr2[i]] = true
			}
		}
		for i := range arr1 {
			_, isPresent := seen[v-arr1[i]]
			if isPresent {
				return true
			}
		}
	}
	return false
}
