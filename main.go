package main

import (
	"fmt"
	"strconv"
)

func greet(name, year string) {
	fmt.Println("Hello! My name is " + name + ".")
	fmt.Println("I was created in " + year + ".")
}

func showName() {
	var name string

	fmt.Println("Please, remind your name.")
	fmt.Scan(&name)
	fmt.Println("What a great name you have, " + name + "!")
}

func guessAge() {
	var rem3, rem5, rem7, age int

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}

func count() {
	var num int

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&num)
	for i := 0; i <= num; i++ {
		fmt.Println(strconv.Itoa(i) + " !")
	}
}

func startQuiz() {
	questionOne := func() {
		var userInput int
		question := "Why do we use methods?"
		answerOne := "1. To repeat a statement multiple times."
		answerTwo := "2. To decompose a program into several small subroutines."
		answerThree := "3. To determine the execution time of a program."
		answerFour := "4. To interrupt the execution of a program."

		answers := [4]string{answerOne, answerTwo, answerThree, answerFour}
		fmt.Println(question)
		for _, value := range answers {
			fmt.Println(value)
		}

		for {
			fmt.Scan(&userInput)
			if userInput == 2 {
				break
			} else {
				fmt.Println("Please, try again.")
			}
		}
	}

	fmt.Println("Let's test your programming knowledge.")
	questionOne()
}

func sayGoodbye() {
	fmt.Println("Congratulations, have a nice day!")
}

func main() {
	greet("Alan", "2023")
	showName()
	guessAge()
	count()
	startQuiz()
	sayGoodbye()
}
