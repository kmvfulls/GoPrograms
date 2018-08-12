package main

import "fmt"

func main() {

	fmt.Println("This is a looping program")

	for i := 0; i < 10000; i++ {
		fmt.Printf("%3d \t %8b \t %2x %q\n", i, i, i, i)
	}


}
