package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type state struct {
	x int
	y int
	a int
	b int
}

func main() {
	defer out.Flush()
	var t int
	fmt.Fscanf(in, "%d\n", &t)
	for tc := 0; tc < t; tc++ {
		var n, m int
		fmt.Fscanf(in, "%d %d\n", &n, &m)
		var sx, sy, fx, fy int
		fmt.Fscanf(in, "%d %d %d %d\n", &sx, &sy, &fx, &fy)
		var p int
		fmt.Fscanf(in, "%d\n", &p)
		u := make([][]bool, n)
		for i := 0; i < m; i++ {
			u[i] = make([]bool, m)
		}
		for i := 0; i < p; i++ {
			var x1, y1, x2, y2 int
			fmt.Fscanf(in, "%d %d %d %d\n", &x1, &x2, &y1, &y2)
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					u[x][y] = true
				}
			}
		}

		dp := make([][]map[int]map[int]int, n)
		for i := 0; i < n; i++ {
			dp[i] = make([]map[int]map[int]int, m)
			for j := 0; j < m; j++ {
				dp[i][j] = make(map[int]map[int]int)
			}
		}
		dp[sx][sy][0] = make(map[int]int)
		dp[sx][sy][0][0] = 0
		q := []state{{x: sx, y: sy, a: 0, b: 0}}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					a := v.a + i
					b := v.b + j
					x := v.x + a
					y := v.y + b
					if x < 0 || x >= n || y < 0 || y >= m || u[x][y] {
						continue
					}
					if _, ok := dp[x][y][a]; !ok {
						dp[x][y][a] = make(map[int]int)
						dp[x][y][a][b] = dp[v.x][v.y][v.a][v.b] + 1
						q = append(q, state{x: x, y: y, a: a, b: b})
						continue
					}

					if _, ok := dp[x][y][a][b]; !ok {
						dp[x][y][a][b] = dp[v.x][v.y][v.a][v.b] + 1
						q = append(q, state{x: x, y: y, a: a, b: b})
					} else if dp[x][y][a][b] > dp[v.x][v.y][v.a][v.b]+1 {
						dp[x][y][a][b] = dp[v.x][v.y][v.a][v.b] + 1
						q = append(q, state{x: x, y: y, a: a, b: b})
					}
				}
			}
		}

		found := false
		var res int
		for _, v1 := range dp[fx][fy] {
			for _, v2 := range v1 {
				if !found {
					res = v2
					found = true
				} else if v2 < res {
					res = v2
				}
			}
		}
		if found {
			fmt.Fprintf(out, "Optimal solution takes %d hops.\n", res)
		} else {
			fmt.Fprintln(out, "No solution.")
		}
	}
}
