package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	username, _, _ := reader.ReadLine()

	seen := make(map[rune]bool)

	for _, char := range username {
		seen[rune(char)] = true
	}

	if len(seen)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
