package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/phrased-org/phrased/args_parser"
	"github.com/phrased-org/phrased/generator"
	"github.com/phrased-org/phrased/wordlists"
	"os"
)

func parseLength(args []string) uint32 {
	if len(args) != 1 {
		return 6
	}
	if s, err := strconv.ParseUint(args[0], 10, 32); err == nil {
		return uint32(s)
	}
	return 6
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	var passphrase string
	var defaultWordlist wordlists.Wordlist
	var phrasedArgs args_parser.PhrasedArgs
	var err error

	defaultWordlist, err = wordlists.RandomWordlist()
	handleError(err)

	flag.Parse()
	phrasedArgs, err = args_parser.Parse(defaultWordlist.Key, flag.Args())
	handleError(err)

	for i := uint32(0); i < phrasedArgs.Amount; i++ {
		passphrase, err = generator.Generate(phrasedArgs.Strength, phrasedArgs.WordlistKey)
		handleError(err)
		fmt.Println(passphrase)
	}
}
