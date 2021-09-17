package main

import "fmt"

func main(){
	var n int
	fmt.Scanln(&n)
	n2 := 0
	n1 := n
	for n!=0{
		arr := n %10
		n2 = n2*10 + arr 
		n /= 10 
	}
	if(n2 == n1 ){
		fmt.Println("This number is palindrome")
	}else{
		fmt.Println("This number is not Palindrome")
	}
}
