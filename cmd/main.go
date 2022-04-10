package main

import (
	"fmt"
	"github.com/NaysayerWasTaken/go-sudoku/m/v2/pkg/generation"
)

func main() {
	g := generation.Grid{}
	g.Init()
	fmt.Println(g)
}
