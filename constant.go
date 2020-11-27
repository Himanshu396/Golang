package main

import "fmt"

//const a int32 = 120

func main() {

	const a = 42
	var b int16 = 27
	fmt.Printf("%v,%T\n", a, a)
	fmt.Printf("%v,%T\n", a+b, a+b)
}
