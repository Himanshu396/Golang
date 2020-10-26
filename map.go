package main

import "fmt"

func main() {
	score := make(map[string]int)
	score["Akash"] = 89
	fmt.Println(score)

	score["john"] = 34
	score["tom"] = 23
	score["leo"] = 56
	score["jordan"] = 78
	fmt.Println(score)

	gettomScore := score["tom"]
	fmt.Println(gettomScore)

}