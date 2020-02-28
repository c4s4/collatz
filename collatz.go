package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func collatz(n int) *[]int {
	list := make([]int, 0, 10)
	for n > 1 {
		list = append(list, n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}
	list = append(list, n)
	return &list
}

func maxList(list *[]int) int {
	max := 0
	for _, i := range *list {
		if i > max {
			max = i
		}
	}
	return max
}

func listToString(list *[]int) string {
	strs := make([]string, len(*list))
	for i, v := range *list {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, " ")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must pass positive integer(s) on command line")
		os.Exit(1)
	}
	for _, s := range os.Args[1:] {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("%s is not an integer\n", s)
			os.Exit(1)
		}
		if n <= 0 {
			fmt.Printf("%d is not a positive integer\n", n)
			os.Exit(1)
		}
		list := collatz(n)
		fmt.Printf("%s (%d) [%d]\n", listToString(list), len(*list)-1, maxList(list))
	}
}
