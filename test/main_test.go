package test

import (
	"fmt"
	"github.com/NaysayerWasTaken/go-sudoku/m/v2/pkg/generation"
	"testing"
)

func TestGenerate(t *testing.T) {
	g := generation.Grid{}
	g.Init()
	fmt.Println(g)
}
