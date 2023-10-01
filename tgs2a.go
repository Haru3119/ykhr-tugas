package main

import "fmt"

func main() {
    	var x, y int 

	fmt.Print("enter value of x and y:")
	fmt.Scan(&x, &y)

	var fx = (1 / (( 3 * x * x) + 10 )) + (( 10 * y ) + 7)
	fmt.Println(fx) 
	
}
