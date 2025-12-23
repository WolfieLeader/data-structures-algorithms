package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	type Cell struct{ row, col int }

	n, m := len(grid), len(grid[0])
	dirs := [4]Cell{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} // ↑, →, ↓, ←

	bfsFn := func(r, c int) {
		q := make([]Cell, 0, 64)
		q = append(q, Cell{r, c}) // Enqueue
		grid[r][c] = '0'          // mark visited immediately

		// NOTE: LOOP until head reaches end
		for head := 0; head < len(q); head++ {
			cur := q[head]

			// NOTE: LOOP all directions and enqueue (O(1))
			for _, dir := range dirs {
				row := cur.row + dir.row
				col := cur.col + dir.col

				// Bound check
				if row < 0 || row >= n || col < 0 || col >= m {
					continue
				}

				if grid[row][col] != '1' {
					continue
				}

				grid[row][col] = '0' // mark visited
				q = append(q, Cell{row, col})
			}
		}
	}

	islands := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == '1' {
				bfsFn(r, c)
				islands++
			}
		}
	}
	return islands
}
