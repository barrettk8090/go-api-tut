package main

import "fmt"

func main() {

	i := 1

	// single condition loop
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// initial/condition/after loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// do this N times
	for i := range 3 {
		fmt.Println("range", i)
	}

	// loop without condition - break
	for {
		fmt.Println("loop")
		break
	}

	// continue
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
