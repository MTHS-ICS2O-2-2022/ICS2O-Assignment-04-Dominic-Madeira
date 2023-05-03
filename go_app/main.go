// Created by: Dominic M.
// Created on: April 2023
//
// This program takes your Subway order and calculates the price.
package main

import "fmt"

func main() {
	var meat int
	var length int
	var price float64

  // input
  fmt.Println("This program takes your Subway order and calculates the price.")
  fmt.Println()
  fmt.Print("What type of meat would you like in your sub? Please enter the corresponding number: ")
	fmt.Println()
  fmt.Print("1. Steak")
	fmt.Println()
  fmt.Print("2. Ham")
	fmt.Println()
  fmt.Print("3. Chicken")
	fmt.Println()
  fmt.Print("4. Turkey")
	fmt.Println()
  fmt.Scanln(&meat)
	fmt.Println()
  fmt.Print("What length would you like your sub? Please enter the corresponding number: ")
	fmt.Println()
  fmt.Print("1. Six-inch")
	fmt.Println()
  fmt.Print("2. Footlong")
	fmt.Println()
  fmt.Scanln(&length)

	// process
	if length == 1 {
		if meat == 1 {
			price = 7.00 * 1.13
		} else if meat == 2 {
			price = 5.00 * 1.13
		} else if meat == 3 {
			price = 5.50 * 1.13
		} else if meat == 4 {
			price = 6.00 * 1.13
		} else {
			fmt.Println("Invalid input.")
		}
	} else if length == 2 {
		if meat == 1 {
			price = (7.00 * 1.75) * 1.13
		} else if meat == 2 {
			price = (5.00 * 1.75) * 1.13
		} else if meat == 3 {
			price = (5.50 * 1.75) * 1.13
		} else if meat == 4 {
			price = (6.00 * 1.75) * 1.13
		} else {
			fmt.Println("Invalid input.")
		}
	} else {
		fmt.Println("Invalid input.")
	}

	// output
	if price > 0{
	priceFormatted := fmt.Sprintf("%.2f", price)
	fmt.Println()
	fmt.Println("Your total is $", priceFormatted, "Thank you for eating at Subway!")
	}

  fmt.Println("\nDone.")
}
