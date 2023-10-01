package main

import "fmt"

func main() {
	var x int
	fmt.Print("Enter a positive integer (less than or equal to 999): ")
	_, err := fmt.Scan(&x)

	if err != nil || x < 1 || x > 999 {
		fmt.Println("Input is not within the valid range (1 to 999).")
		return
	}

	d1 := x / 100       
	d2 := (x / 10) % 10 
	d3 := x % 10        

	fmt.Printf("The first digit is %d\n", d1)
	fmt.Printf("The second digit is %d\n", d2)
	fmt.Printf("The third digit is %d\n", d3)
}
