package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var num1 = rand.Intn(100)
	var num2 = rand.Intn(100)

	randomize(num1, num2)
}

func randomize(num1 int, num2 int) {
	var arr [1]int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 0; i++ {
		arr[i] = rand.Intn(100)
	}

	// any number from 0 to 24 will return multiplication
	if arr[0] <= 24 {
		multiplication(num1, num2)
		return
	// any number from 25 and 49 will return division
	} else if arr[0] >= 25 && arr[0] <= 49 {
		// can't divide a number by 0
		if num1 == 0 || num2 == 0 {
			division(num1 + 1, num2 + 1)
		} else {
			division(num1, num2)
		}
	// any number from 50 to 74 will return addition
	} else if arr[0] >= 50 && arr[0] <= 74 {
		addition(num1, num2)
		return
	// any number from 75 to 99 will return subtraction
	} else if arr[0] >= 75 {
		subtraction(num1, num2)
		return
	} else {
		fmt.Println("Error randomizing which question gets asked")
		return
	}
}

// gives an addition question
func addition(num1 int, num2 int) {
	var addRes int

	fmt.Println(num1, "+", num2)

	fmt.Println("Your answer:")
	fmt.Scanln(&addRes)

	if addRes == num1 + num2 {
		fmt.Println("Correct!")
		return
	} else {
		fmt.Println("Sorry, that's incorrect")
		return
	}
}

// gives a subtraction question
func subtraction(num1 int, num2 int) {
	var subRes int

	fmt.Println(num1, "-", num2)

	fmt.Println("Your answer:")
	fmt.Scanln(&subRes)

	if subRes == num1 - num2 {
		fmt.Println("Correct!")
		return
	} else {
		fmt.Println("Sorry, that's incorrect")
		return
	}
}

// gives a multiplication question
func multiplication(num1 int, num2 int) {
	var mulRes int

	fmt.Println(num1, "*", num2)

	fmt.Println("Your answer:")
	fmt.Scanln(&mulRes)

	if mulRes == num1 * num2 {
		fmt.Println("Correct!")
		return
	} else {
		fmt.Println("Sorry, that's incorrect")
		return
	}
}

// gives a division question
func division(num1 int, num2 int) {
	var divRes int

	fmt.Println(num1, "/", num2)

	fmt.Println("Your answer:")
	fmt.Scanln(&divRes)

	if divRes == num1 / num2 {
		fmt.Println("Correct!")
		return
	} else {
		fmt.Println("Sorry, that's incorrect")
		return
	}
}
