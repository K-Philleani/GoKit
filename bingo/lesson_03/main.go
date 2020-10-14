package main

import "fmt"

func main() {
	// if
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score >75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	if grade := 90; grade >= 90 {
		fmt.Println("A")
	} else if score >75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// switch
	age := 2
	switch age {
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
		default:
			fmt.Println("unknown")
	}

	for i := 0; i <10; i++ {
		fmt.Println("i:", i)
	}


}
