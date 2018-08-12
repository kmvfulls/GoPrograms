package main

import (
	"fmt"
	"github.com/kmvfulls/GoPrograms/04_vars/visi"
)


var b string


func main() {

	a := 10
	b := "go language"
	c := 4.18
	d := true


	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Printf("max val is - %d\n", max(40))

	fmt.Println(visi.MyName)


}




func max (x int) int {

	return 42 + x
}


