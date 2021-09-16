package main

import "fmt"

func main(){
	var n int
	fmt.Scanln(&n)
  fmt.Println("These are Fizzbuz numbers !!!")

	for i:=1; i<=n; i++ {
		if (i % 3 != 0){
			if (i % 5 != 0){
				fmt.Println(i)
			} else {
				fmt.Println("Buzz")
			  }
		} else if ( i % 5 != 0){
			fmt.Println("Fizz")
			   }else {
				    fmt.Println("FizzBuz") 
			    }
	}
}
