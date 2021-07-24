package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 20, "number of Goroutines")
	flag.Parse()
	count := *n
	fmt.Printf("Goint to create %d goroutines.\n", count)

	var waitGroup sync.WaitGroup
	/*
		type WaitGroup struct {
			noCopy noCopy
			state1 [12]byte
			sema uint32
		}
	*/

	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%3d", x)
		}(i)
	}
	waitGroup.Add(1) //deadlock(데드락 : sync.Done()이 호출되기를 끝없이 기다리는 상태, 교착상태) 유도

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...\n")

}
