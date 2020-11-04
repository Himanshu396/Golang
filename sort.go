package main  
import (  
    "sort"  
    "fmt"  
)  
func main() {  
  
    intValue := []int{4,40,3,56,36}  
    sort.Ints(intValue)  
    fmt.Println("Ints:   ", intValue)  
  
    floatValue := []float64{12.5, 25.5, 56.5, 81.5}  
    sort.Float64s(floatValue)  
    fmt.Println("floatValue:   ", floatValue)  
  
    stringValue := []string{"ram", "shyam", "ghanshyam"}  
    sort.Strings(stringValue)  
    fmt.Println("Strings:", stringValue)  
  
    str := sort.Float64sAreSorted(floatValue)  
    fmt.Println("Sorted: ", str)  
}  