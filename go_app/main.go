// // Created by: Dominic M.
// // Created on: April 2023
// //
// // This program takes your Subway order and calculates the price.
// package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// )

// func trimRightSpace(stringFormatted string) string {
// 	return strings.TrimRightFunc(stringFormatted, unicode.IsSpace)
// }

// func main() {
// 	const tax float64 = 1.13
// 	const sixInchPrice float64 = 1
// 	const twelveInchPrice float64 = 1.75
// 	const costSteak float64 = 7
// 	const costHam float64 = 5
// 	const costChicken float64 = 5.50
// 	const costTurkey float64 = 6
// 	var meatNumber int
// 	var length int
// 	var cost float64
// 	var price float64
// 	var lengthPrice float64

// 	// input
// 	fmt.Println("This program takes your Subway order and calculates the price.")
// 	fmt.Println()
// 	fmt.Print("What type of meat would you like in your sub? Please enter the corresponding number: ")
// 	fmt.Println()
// 	fmt.Print("1. Steak")
// 	fmt.Println()
// 	fmt.Print("2. Ham")
// 	fmt.Println()
// 	fmt.Print("3. Chicken")
// 	fmt.Println()
// 	fmt.Print("4. Turkey")
// 	fmt.Println()
// 	fmt.Scanln(&meatNumber)
// 	fmt.Println()
// 	fmt.Print("What length would you like your sub? Please enter the corresponding number: ")
// 	fmt.Println()
// 	fmt.Print("1. Six-inch")
// 	fmt.Println()
// 	fmt.Print("2. Footlong")
// 	fmt.Println()
// 	fmt.Scanln(&length)

// 	// process
// 	if length == 1 {
// 		lengthPrice = sixInchPrice
// 	} else {
// 		lengthPrice = twelveInchPrice
// 	}

// 	if meatNumber == 1 {
// 		cost = costSteak
// 	} else if meatNumber == 2 {
// 		cost = costHam
// 	} else if meatNumber == 3 {
// 		cost = costChicken
// 	} else {
// 		cost = costTurkey
// 	}

// 	price = (cost * lengthPrice) * tax

// 	// output
// 	priceFormatted := fmt.Sprintf("%.2f", price)
// 	stringFormatted := trimRightSpace(`Your total is $`)

// 	fmt.Println()
// 	fmt.Printf("%s", stringFormatted)
// 	fmt.Print(priceFormatted, " Thank you for eating at Subway!")
// 	fmt.Println()

// 	fmt.Println("\nDone.")
// }
