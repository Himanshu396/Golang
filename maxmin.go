package main

import "fmt"

func main() {
	var x int=50
	var y int =100
	var res int
	res=max(x,y)
	fmt.Printf("max value is %d\n",res)
}
func max(num1,num2 int)int{
	var result int 
	if(num1>num2){
		result=num1
	}else{
		result=num2
	}
	return result
}