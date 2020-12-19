# word-chain

[![Go Report Card](https://goreportcard.com/badge/github.com/theskch/word-chain)](https://goreportcard.com/report/github.com/theskch/word-chain)

Word-chain is a simple CLI tool built to find one of the shortest word chains, in regards to the supplied dictionary. The task was inspired by the http://codekata.com/kata/kata19-word-chains/.<br/> 
When run, the prompt asks for the first word of the chain, and the last word of the chain.
If first and last words of the chain differe in length, error message is shown. If the first word or the last word is not present in the dictionary, error message is shown. You can set the desired text dictionary by passing the -p flag `word-chain -p path/to/dictionary`. Dictionary should contain every word in new line.

## Chain search
Since we are searching for the shortest word chain, Breadth-first search algorithm is used for traversing trough graph data structure. Data structure is created using the words from the dictionary. Two words are connected in graph if they differ only by one character (case-sensitive).
