package main

import "fmt"


func main() {
    var n int
    fmt.Print("Enter the number of items: ")
    fmt.Scanln(&n)
    fib1 := 1
    fib2 := 1
    

    fmt.Println(fib1)
    fmt.Println(fib2)
    
    for i:=0; i<n; i++{
        fib3 := fib1 + fib2
        fmt.Println(fib3)
        fib1 = fib2
        fib2 = fib3
        
        
    }  
}
