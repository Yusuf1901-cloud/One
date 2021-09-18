package main

import "fmt"

func main(){
	var n int 
	arr := make([]int, 0)
	fmt.Print("Enter number your arr items: ")
	fmt.Scanln(&n)
	fmt.Println("Now enter elements of your array : ")
	for i:=0; i<n; i++{
		var n0 int
		fmt.Scanln(&n0)
		arr = append(arr, n0)
	}
	fmt.Print("Your array: ")
	fmt.Println(arr)
	fmt.Println(len(arr))
	for i:=0; i<len(arr); i++{
		for j:=i+1; j< len((arr)); j++{
			if arr[j] == arr[i]{
				fmt.Printf("In your arr has duplicate: %v\n", arr[i])

			}
		}

	}
	
}
