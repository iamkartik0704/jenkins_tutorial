package main

import "fmt"

func Greeting() string {
	return "Hello from Go Application for Jenkins CI/CD!"
}

func main() {
	fmt.Println(Greeting())
}