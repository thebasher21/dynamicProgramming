package game

type scores struct {
	p1 int
	p2 int
}

func MaxDifference(N int, Tiles []int) int {
	table := make([][]scores, N)
	for i := range table {
		table[i] = make([]scores, N)
	}
	for i := 0; i < N; i++ {
		table[i][i] = scores{p1: Tiles[i], p2: 0}
	}
	for j := 2; j <= len(Tiles); j++ {
		for i := 0; i <= len(Tiles)-j; i++ {
			k := i + j - 1
			if table[i+1][k].p2+Tiles[i] > table[i][k-1].p2+Tiles[k] {
				table[i][k] = scores{
					p1: table[i+1][k].p2 + Tiles[i],
					p2: table[i+1][k].p1,
				}
			} else {
				table[i][k] = scores{
					p1: table[i][k-1].p2 + Tiles[k],
					p2: table[i][k-1].p1,
				}
			}

		}
	}
	return table[0][N-1].p1 - table[0][N-1].p2
}
