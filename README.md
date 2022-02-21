For a blog post in Hashnode where I give a couple examples of easy projects so that devs can see the difference between their current favorite language and Golang. It's also a quick and fun test for myself and could be for others as well.

---

# What I want to do
Create a simple math quiz that randomly gives you an addition, subtraction, multiplication, or division question.
- generate 2 random numbers (num1, num2)
- have the result (res) be equal to the numbers
- keep all numbers less than or equal to 100
    - May increase limit later, doing this just for simplicity right now
- optionally, have all results be an int so that there aren't any crazy answers like 9.5648...

# My Process
### Starting With Addition
- Create a min and max value
```go
var min = 0
var max = 100
```

- Create the 2 nums with random values. Found in [docs](https://golangdocs.com/generate-random-numbers-in-golang)
```go 
var num1 = rand.Intn(max-min) + min
var num2 = rand.Intn(max-min) + min
```

- Have a variable that's the result of the 2 nums being added
- Print the 2 nums
- Let user input their answer
- If the answer is correct then print `Correct!`. Else print `Sorry, that's incorrect`.
```go
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
```

This is what it looks like so far:
```go
func main() {
	var min = 0
	var max = 100

	var num1 = rand.Intn(max-min) + min
	var num2 = rand.Intn(max-min) + min

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
```
Testing the application.
<img src="/imgs-for-README/addRes1.png">

The initial test works! However, it will always print out `81+87`. We need to randomize the numbers each time.
- Change the logic of num1 and num2
- Found an [article](https://golangbyexample.com/generate-random-number-golang/) that describes how to do this with the time package
    - the docs had something similar but for some reason it couldn't click in my brain until I saw this one
- Delete the `min` and `max` variables and add the Unix randomizer
```go
rand.Seed(time.Now().UnixNano())

var num1 = rand.Intn(100)
var num2 = rand.Intn(100)
```
Testing the application.
<img src="/imgs-for-README/addRes2.png">

It works as intended. If you just want the quiz to be addition only this is good enough, however, I want to include subtraction, multiplication, and division.

Before we can move on, I want to separate the logic for the addition problem into it's own function. If we don't do this before adding in the other logic the main func will look horrendous.
```go
func main() {
	rand.Seed(time.Now().UnixNano())

	var num1 = rand.Intn(100)
	var num2 = rand.Intn(100)

	addition(num1, num2)
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
```
Now that addition is in it's own function, we can start on subtraction.

### Subtraction
The subtraction function is almost the exact same as addition, but swapping `+` for `-`.
```go
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
```
Then in the main func, swap `addition(num1, num2)` for `subtraction(num1, num2)`

Testing `func subtraction`.
<img src="/imgs-for-README/subRes1.png">




---
#### Ideas to implement later
- Have the program ask for an answer that equals the addition problem, subtraction problem, multiplication problem, and division problem all at once.
    - try implementing Go routines to execute all 4 functions simultaneously
    - would be good for an intro to Go routines blog