package main

import (
	"fmt"
	"github.com/kr/pretty"
)
type item struct {
	name  string
	peso  int
	valor int
}

const (
	bagCapacity = 5
	itemNumber  = 4
)

func main() {
	objects := []item{
		{"", 0, 0},
		{"lápis", 4, 500},
		{"caneta", 2, 400},
		{"caderno", 1, 300},
		{"borracha", 3, 450},
	}
	var bag [itemNumber + 1][bagCapacity + 1]int
	for i := 0; i <= itemNumber; i++ {
		bag[i][0] = 0
	}
	for b := 1; b <= bagCapacity; b++ {
		bag[0][b] = 0
		for i := 1; i <= itemNumber; i++ {
			s := bag[i-1][b]
			if objects[i].peso <= b {
				sum := bag[i-1][b-objects[i].peso] + objects[i].valor
				if s < sum {
					s = sum
				}
			}
			bag[i][b] = s
		}
	}
	pretty.Println(bag)
	i := itemNumber
	b := bagCapacity
	aux := bag[i][b]

	for aux != 0 {
		if bag[i][b] != bag[i-1][b] {
			fmt.Println(objects[i].name)
			b -= objects[i].peso
		}
		i--
		aux = bag[i][b]
	}
}
