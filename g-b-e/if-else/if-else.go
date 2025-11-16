package main

import "fmt"

func main() {

	i := 7

	if i%2 == 0 {
		fmt.Printf("%v is even\n", i)
	} else {
		fmt.Printf("%v is odd\n", i)
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 are even")
	}

	if num := 10; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
