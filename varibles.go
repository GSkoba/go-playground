package main

import "fmt"

func main() {
	fmt.Println("variables")

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d, e = true, false
	fmt.Println(d, e)
	d = false

	var f int
	fmt.Println(f)

	g := "apple"
	fmt.Println(g)
}
