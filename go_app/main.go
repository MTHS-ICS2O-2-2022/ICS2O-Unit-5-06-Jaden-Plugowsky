// Created by: Jaden Plugowsky
// Created on: May 2023
//
// This program takes two user-given numbers and multiplies them together using only addition and loops

package main

import (
	"fmt"
	"math"
)

var firstNumber float64
var secondNumber float64
var counter float64
var answer float64
var moduloFirst float64
var moduloSecond float64

func main() {
	// This function takes two user-given numbers and multiplies them together using only addition and loops
	// Input
	fmt.Println("This program takes two user-given numbers and multiplies them together using only addition and loops.")
	fmt.Println("\nNote: This does not work with decimals.")
	fmt.Println("\nEnter the first number: ")
	fmt.Scanln(&firstNumber)
	fmt.Println("\nEnter the second number: ")
	fmt.Scanln(&secondNumber)

	// Process
	moduloFirst = math.Mod(firstNumber, 1) // Used to help check if the numbers are decimals or not. If this number is different than zero, the number is a decimal.
	moduloSecond = math.Mod(secondNumber, 1)

	if moduloFirst == 0 && moduloSecond == 0 { // Checks if either number is a decimal, if not, it runs the code.
		if secondNumber >= 0 {
			// for all positive or firstNumber negative mixes
			for counter < secondNumber {
				answer += firstNumber
				counter++
			}
		} else {
			// for everything else: all negative or secondNumber negative mixes
			for counter > secondNumber {
				answer -= firstNumber
				counter--
			}
		}
	} else {
		fmt.Println("\nOne or both of your inputs is not a whole number.")
		fmt.Println("Please try again with whole numbers.")
	}

	// Output
	if moduloFirst == 0 && moduloSecond == 0 {
		fmt.Println("\nThe answer is:", answer)
	}

	fmt.Println("\nDone.")
}
