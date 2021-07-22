package main

import "fmt"

type aStructure struct {
	name   string
	height int
	weight int
}

func main() {
	ps := new(aStructure)
	sP := new([]aStructure)

	ps.name = "park jin sol"
	ps.height = 100
	ps.weight = 100

	*sP = append(*sP, *ps)

	fmt.Println(ps)
	fmt.Println(sP)

}
