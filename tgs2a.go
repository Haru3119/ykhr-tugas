package main

import "fmt"

func main() {
    // var x, y int

	// fmt.Print("Enter values of  x and y : ")
	// fmt.Scan(&x, &y)

	// var fx = (1 / ((3 * x * x) + 10 )) + ((10 * y) + 7)
	// fmt.Println(fx)



	
	// Input a positive integer x
	var x int
	fmt.Print("Enter a positive integer (less than or equal to 999): ")
	_, err := fmt.Scan(&x)

	if err != nil || x < 1 || x > 999 {
		fmt.Println("Input is not within the valid range (1 to 999).")
		return
	}

	// Extract the individual digits
	d1 := x / 100       // Extract the hundreds digit
	d2 := (x / 10) % 10 // Extract the tens digit
	d3 := x % 10        // Extract the ones digit

	// Output the result
	fmt.Printf("The first digit is %d\n", d1)
	fmt.Printf("The second digit is %d\n", d2)
	fmt.Printf("The third digit is %d\n", d3)


}