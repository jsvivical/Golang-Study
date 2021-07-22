package main

import "fmt"

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	var s1 aStructure //선언
	p1 := aStructure{"fmt", 12, -2}
	p2 := aStructure{
		person: "fmt",
		height: 12,
		weight: -2,
	}
	structures := make([]aStructure, 5)

	structures = append(structures, p1)
	structures = append(structures, p2)
	structures = append(structures, aStructure{"new", 10, 10})

	fmt.Println(s1)
	for i, v := range structures {
		fmt.Printf("[%d] %v\n", i, v)
	}

}
