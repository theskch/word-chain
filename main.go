package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/theskch/word-chain/chain"
	"github.com/theskch/word-chain/dictionary"
)

func main() {
	path := flag.String("p", "dictionary/dictionary.txt", "path to the dictionary containing all the words")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Word-chain Shell")
	fmt.Println("----------------")
	fmt.Println("first word:")
	fmt.Print("-> ")
	firstWord, _ := reader.ReadString('\n')
	fmt.Println("last word:")
	fmt.Print("-> ")
	secondWord, _ := reader.ReadString('\n')

	firstWord = strings.Replace(firstWord, "\n", "", -1)
	secondWord = strings.Replace(secondWord, "\n", "", -1)

	dict, err := dictionary.NewTextDictionary(*path)
	if err != nil {
		return
	}

	chain, err := chain.FindShortestChain(firstWord, secondWord, dict.GetWords(4))
	if err != nil {
		fmt.Printf("faild to create chain, reason: %s", err)
	} else if chain.Length() == 0 {
		fmt.Printf("not possible to create chain with the current dictionary")
	} else {
		fmt.Printf("%s", chain.String())
	}

}
