package main
import "fmt"

func max(num1 int,num2 int)int{
    var result int 
    if (num1 >= num2){    
        result = num1
    }else{
        result = num2
    }
    return result
}
func main(){
    fmt.Println(max(20,10))
}
