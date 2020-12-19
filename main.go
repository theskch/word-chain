package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/briandowns/spinner"
	"github.com/thatisuday/commando"
	"github.com/theskch/word-chain/chain"
	"github.com/theskch/word-chain/dictionary"
)

func main() {
	commando.
		SetExecutableName("word-chain").
		SetVersion("1.0.1").
		SetDescription("This is a tool that builds a chain of words, starting from one particular word and ending with another.\nSuccessive entries in the chain must all be real words and each can differ from previous word by just one letter.")
	commando.
		Register(nil).
		AddFlag("path,p", "path to the dictionary of words to use", commando.String, "dictionary/dictionary.txt").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			path, err := flags["path"].GetString()
			if err != nil {
				fmt.Printf("invalid path flag")
				return
			}

			dict, err := dictionary.NewTextDictionary(path)
			if err != nil {
				fmt.Printf("%s", err)
				return
			}

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

			s := spinner.New(spinner.CharSets[35], 200*time.Millisecond)
			fmt.Println("\nSearching: ")
			s.Start()
			chain, err := chain.FindShortestChain(firstWord, secondWord, dict.GetWords(utf8.RuneCountInString(firstWord)))
			s.Stop()
			if err != nil {
				fmt.Printf("faild to create chain, reason: %s", err)
			} else if chain.Length() == 0 {
				fmt.Printf("not possible to create chain with the current dictionary")
			} else {
				fmt.Printf("%s", chain.String())
			}

		})

	commando.Parse(nil)
}
