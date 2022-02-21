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

	addition(num1, num2)
	subtraction(num1, num2)
}

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
