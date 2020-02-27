package main

import (
	"fmt"
	"os"
	"strconv"
)

func collatz(n int) {
	fmt.Print(n)
	if n == 1 {
		return
	}
	if n%2 == 0 {
		fmt.Print(" ")
		collatz(n / 2)
	} else {
		fmt.Print(" ")
		collatz(n*3 + 1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must pass positive integer(s) on command line")
		os.Exit(1)
	}
	for _, s := range os.Args[1:] {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("%s is not an integer\n", os.Args[0])
			os.Exit(1)
		}
		if n <= 0 {
			fmt.Printf("%d is not a positive integer", n)
			os.Exit(1)
		}
		collatz(n)
		fmt.Println()
	}
}
