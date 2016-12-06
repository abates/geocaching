package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var max int
var maxPath []int

func sum(layer []int) int {
	sum := 0
	for _, i := range layer {
		sum += i
	}
	return sum
}

func walk(layers [][]int, row, column int, stack []int) {
	stack = append(stack, layers[row][column])
	if row == len(layers)-1 {
		newMax := sum(stack)
		//fmt.Printf("%d %v\n", newMax, stack)
		if newMax > max {
			max = newMax
			if maxPath == nil {
				maxPath = make([]int, len(stack))
			}
			copy(maxPath, stack)
		}
		return
	}
	walk(layers, row+1, column, stack)
	walk(layers, row+1, column+1, stack)
	stack = stack[0 : len(stack)-1]
}

func main() {
	layers := make([][]int, 0)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		if len(line) == 0 {
			continue
		}

		layer := make([]int, 0)
		for _, s := range strings.Split(string(line), " ") {
			i, _ := strconv.Atoi(s)
			layer = append(layer, i)
		}
		layers = append(layers, layer)
	}

	walk(layers, 0, 0, []int{})
	fmt.Printf("%v\n", maxPath)
}
