package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stderr, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run main.go TODO
	// To save the output from os.Stdout into eyes.txt, you can use the following command:
	// go run main.go 1> eyes.txt
	// To save the output from os.Stderr into dice.log, you can use the following command:
	// go run main.go 2> dice.log
}
