package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Word-chain Shell")
	fmt.Println("----------------")
	fmt.Print("->")
	firstWord, _ := reader.ReadString('\n')
	fmt.Print("->")
	secondWord, _ := reader.ReadString('\n')

	firstWord = strings.Replace(firstWord, "\n", "", -1)
	secondWord = strings.Replace(secondWord, "\n", "", -1)

	fmt.Printf("first word: %s second word: %s", firstWord, secondWord)
}
