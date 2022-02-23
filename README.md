For a blog post in Hashnode where I give a couple examples of easy projects so that devs can see the difference between their current favorite language and Golang. It's also a quick and fun test for myself and could be for others as well.

---

# What I want to do
Create a simple math quiz that randomly gives you an addition, subtraction, multiplication, or division question.
- generate 2 random numbers (num1, num2)
- have the result (res) be equal to the numbers
- keep all numbers less than or equal to 100
    - May increase limit later, doing this just for simplicity right now

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

Before we can move on, I want to separate the logic for the addition problem into it's own function. If we don't do this before adding in the other logic, the main func will look horrendous.
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
Now that addition is in it's own function, we can start on the other math questions.

### Subtraction, Multiplication, and Division
- The subtraction function is almost the exact same as addition, but swapping `+` for `-`.
- Then in the main func, swap `addition(num1, num2)` for `subtraction(num1, num2)`
- Then for multiplication it will look similar... Blah, blah, blah you get the idea

Testing `func subtraction`.

<img src="/imgs-for-README/subRes1.png">

Testing `func division` and `func multiplication`

<img src="/imgs-for-README/mulAndDivRes1.png">

### Randomizing Which Function Gets Called
- We need a way to randomize which function gets called back into the main function
- When I was looking in the golang [docs](https://golangdocs.com/generate-random-numbers-in-golang) I saw an example of a `for loop` randomizing integers in an array
- Since there are 4 questions, maybe I could create an array of 4 values, loop over those values with a random number, and if the result of a particular index in the array equals a certain number it would return the associated function?
```go
func randomize(num1 int, num2 int) {
	var arr [4]int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 3; i++ {
		arr[i] = rand.Intn(4)
	}

	var res int

	if res == arr[0] {
		multiplication(num1, num2)
		return
	} else if res == arr[1] {
		division(num1, num2)
		return
	} else if res == arr[2] {
		addition(num1, num2)
		return
	} else if res == arr[3] {
		subtraction(num1, num2)
		return
	} 
}
```
- Then call `randomize(num1, num2)` inside of the main func and see if it works
- Annndddd... nothing happened :/
- (A bit of foreshadowing to a paragraph later; this is why you should print to the console when you're in doubt and walk away from the computer for a few mins when you're stuck)
- I'm going to try again but this time remove `return` from the if else statements

Testing `func randomize`

<img src="/imgs-for-README/randRes1.png">

- So it works sometimes but doesn't work other times
- Maybe if I change `arr[i] = rand.Intn(4)` to `arr[i] = rand.Intn(3)` I can get it to work every time?

Testing `func randomize` again

<img src="/imgs-for-README/randRes2.png">

- So it works, however I found an interesting bug
- When I input for the result of `81 / 0` it gave me an error
- That's because a number can't be divided by 0
- I tried this on my IPhone's calculator and it said "Error"
- This can be avoided by changing the parameters of calling the division function to `else if res == arr[1] && num1 != 0 && num2 != 0`
- After taking a 15 min break and eating some ice cream, I realized the way it chooses which question to ask is kind of dumb

### Refactoring Randomization
- For some reason I thought `res` would equal to... honestly, now that I'm sitting back down and typing out my reasoning, I don't know what I thought it was equal to. In reality it's equal to 0, so the first index in the array to equal 0 will be called
- That's a problem because if the array is `[0 0 1 1]`, multiplication will be called because it's first in the stack
- Also, if the array is `[1 1 2 2]` then nothing will be called
	- This is why it didn't work earlier, not because it had `return` in the if else statements
- I still need a way to evenly decide which question will be asked
- Since there are 4 questions, each should have about a 25% chance of being asked
- What if, instead of having an array of four values that could equal anything from 0-3, the array only has one value that could equal anything from 0-99?
	- if the value is anywhere from 0-24 it would return multiplication
	- 25-49 would be division
	- 50-74 would be addition
	- 75-99 would be subtraction
- And in case num1 or num2 is equal to 0, it'll return addition
	- It won't be a perfect 25% chance for division and addition, but it's very close, and avoids errors
```go
func randomize(num1 int, num2 int) {
	var arr [1]int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 0; i++ {
		arr[i] = rand.Intn(100)
	}

	if arr[0] <= 24 {
		multiplication(num1, num2)
		return
	} else if arr[0] >= 25 && arr[0] <= 49 && num1 != 0 && num2 != 0{
		division(num1, num2)
		return
	} else if arr[0] >= 50 && arr[0] <= 74 {
		addition(num1, num2)
		return
	} else if arr[0] >= 75 {
		subtraction(num1, num2)
		return
	} else {
		addition(num1, num2)
		return
	}
}
```

Testing the result

<img src="/imgs-for-README/finalResult.png">

Finally got it working!!!!



---
#### Ideas to implement later
- Have the program ask for an answer that equals the addition problem, subtraction problem, multiplication problem, and division problem all at once.
    - try implementing Go routines to execute all 4 functions simultaneously