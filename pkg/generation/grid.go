package generation

import (
	"math/rand"
	"time"
)

type Grid struct {
	Content [9][9]int
	Solved  bool
}

func (g *Grid) Init() {
	k, n := 1, 1

	for i := 0; i < 9; i++ {
		k = n
		for j := 0; j < 9; j++ {
			if k <= 9 {
				g.Content[i][j] = k
				k++
			} else {
				k = 1
				g.Content[i][j] = k
				k++
			}
		}
		n = k + 3
		if k == 10 {
			n = 4
		}
		if n > 9 {
			n = n%9 + 1
		}
	}
	g.randomize(0)
}

func (g *Grid) randomize(c int) {
	var k1, k2 int
	max, min := 1, 0

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		k1 = rand.Perm(max - min + 1)[1] + min

		k2 = rand.Perm(max - min + 1)[1] + min

		for k1 == k2 {
			max += 3
			min += 3
		}

		if c == 1 {
			permRow(k1, k2, g)
		} else if c == 0 {
			permCol(k1, k2, g)
		}
	}

}

func permRow(k1 int, k2 int, g *Grid) {
	var temp int
	for i := 0; i < 9; i++ {
		temp = g.Content[k1][i]
		g.Content[k1][i] = g.Content[k2][i]
		g.Content[k2][i] = temp
	}
}

func permCol(k1 int, k2 int, g *Grid) {
	var temp int
	for i := 0; i < 9; i++ {
		temp = g.Content[i][k1]
		g.Content[i][k1] = g.Content[i][k2]
		g.Content[i][k2] = temp
	}
}
