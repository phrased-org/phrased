package main

import (
	"flag"
	"fmt"
	"strconv"

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

func main() {
	var wordlistType string
	var phraseLength uint32
	var passphrase string
	var defaultWordlist wordlists.Wordlist
	var err error

	defaultWordlist, err = wordlists.RandomWordlist()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	flag.StringVar(&wordlistType, "wordlist", defaultWordlist.Key, "which wordlist to use")
	flag.Parse()
	phraseLength = parseLength(flag.Args())

	passphrase, err = generator.Generate(phraseLength, wordlistType)
	if err == nil {
		fmt.Println(passphrase)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
