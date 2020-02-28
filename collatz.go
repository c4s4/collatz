package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func collatz(n int, l []int) []int {
	l = append(l, n)
	if n == 1 {
		return l
	}
	if n%2 == 0 {
		return collatz(n/2, l)
	} else {
		return collatz(n*3+1, l)
	}
}

func max(l []int) int {
	m := 0
	for _, i := range l {
		if i > m {
			m = i
		}
	}
	return m
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
		list := collatz(n, nil)
		strs := make([]string, len(list))
		for i, v := range list {
			strs[i] = strconv.Itoa(v)
		}
		fmt.Printf(strings.Join(strs, " "))
		fmt.Printf(" (%d) [%d]\n", len(list)-1, max(list))
	}
}
