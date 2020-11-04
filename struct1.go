package main  
import (  
   "fmt"  
)  
type person struct {  
   firstName string  
   lastName  string  
   age       int  
}  
func main() {  
   x := person{age: 22, firstName: "Himanshu", lastName: "Kumar", }  
   fmt.Println(x)  
   fmt.Println(x.firstName)  
}  