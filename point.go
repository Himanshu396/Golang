package main

import "fmt"

func main(){
    var p *int
    if p!= nil{
        fmt.Println("P is having a value: ", *p)
    } else {
        fmt.Println("Watchout for nil values")
    }
}