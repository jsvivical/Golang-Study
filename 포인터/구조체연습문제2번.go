package main

import "fmt"

type Actor struct {
	name  string
	hp    int
	speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	p := new(Actor)
	p.name = name
	p.hp = hp
	p.speed = speed

	return p
}

func main() {
	var p *Actor

	p = NewActor("박진솔", 100, 100.29)
	fmt.Println(*p)
}
