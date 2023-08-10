package main

import "fmt"

func main() {
	var name string

	fmt.Println("Hello! My name is Alan.")
	fmt.Println("I was created in 2023")
	fmt.Println("Please, remind your name.")

	fmt.Scan(&name)

	fmt.Println("What a great name you have, " + name + "!")
}
