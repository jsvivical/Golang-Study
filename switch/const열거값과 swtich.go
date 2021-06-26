package main

import "fmt"

type ColorType int

const (
	Red    ColorType = iota
	Blue   ColorType = iota
	Green  ColorType = iota
	Yellow ColorType = iota
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"

	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("my favorite color is " + colorToString(getMyFavoriteColor()))
}
