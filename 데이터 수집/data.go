package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reade := csv.NewReader(f)

}
