package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("type something:  ")
	scanner.Scan()
	input := scanner.Text()
	//input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("you typed: %q", input)
}
