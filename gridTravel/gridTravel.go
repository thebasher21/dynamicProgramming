package gridtravel

func GridTravelIter(m, n int) int {

	table := make([][]int, m+1)
	//table[1][1] = 1
	for i := range table {
		table[i] = make([]int, n+1)
	}
	table[1][1] = 1
	for i := range table {
		for j := range table[i] {
			current := table[i][j]
			if j+1 <= n {
				table[i][j+1] += current
			}
			if i+1 <= m {
				table[i+1][j] += current
			}
		}
	}
	return table[m][n]
}
