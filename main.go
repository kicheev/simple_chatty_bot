package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello! My name is Alan.")
	fmt.Println("I was created in 2023")
	fmt.Println("Please, remind your name.")

	var name string
	fmt.Scan(&name)

	fmt.Println("What a great name you have, " + name + "!")
	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")

	var remainder3, remainder5, remainder7 int
	var age int
	fmt.Scan(&remainder3, &remainder5, &remainder7)

	age = (remainder3*70 + remainder5*21 + remainder7*15) % 105

	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}
