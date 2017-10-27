package main

import (
        "fmt"
        "flag"
        "strconv"

        "github.com/phrase-yourself/phrased/phrased"
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
        var wordlistType string;
        var phraseLength uint32;

        flag.StringVar(&wordlistType, "wordlist", "eff", "which wordlist to use")
        flag.Parse()
        phraseLength = parseLength(flag.Args())

        var passphrase, err = phrased.Generate(phraseLength, wordlistType)
        if err == nil {
                fmt.Println(passphrase)
        } else {
                fmt.Println(err)
        }
}
