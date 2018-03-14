package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/phrased-org/phrased/args_parser"
	"github.com/phrased-org/phrased/generator"
	"github.com/phrased-org/phrased/wordlists"
	"os"
	"strings"
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

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			`Usage of %s: [wordlist] [strength] [count]

Phrased will generate [count] different passphrases, each of a strength of
[strength]. By specifiying a [wordlist], phrased will only pick words from
this given wordlist.


Example calls:

To generate one passphrase of default strength from a random wordlist run:
$ phrased

To generate one passphrased of a strenght of 6 run:
$ phrased 6

To generate 4 passphrases of a strength of 6 run:
$ phrased 6 4

To generate a passphrase from the "literature-en" wordlist run:
$ phrased literature-en

To generate a passphrase from the "literature-en" wordlist of strength 8 run:
$ phrased literature-en 8


Possible wordlist are:
`, os.Args[0])
		var allWordlists, e = wordlists.Wordlists()
		handleError(e)
		for _, wordlist := range allWordlists {
			fmt.Printf("%-22s  [%s] '%s' \n", wordlist.Key, strings.Join(wordlist.Languages, ", "), wordlist.Name)
		}
		flag.PrintDefaults()
	}
	flag.Parse()

	phrasedArgs, err = args_parser.Parse(defaultWordlist.Key, flag.Args())
	handleError(err)

	for i := uint32(0); i < phrasedArgs.Amount; i++ {
		passphrase, err = generator.Generate(phrasedArgs.Strength, phrasedArgs.WordlistKey)
		handleError(err)
		fmt.Println(passphrase)
	}
}
