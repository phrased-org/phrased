package generator

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/phrased-org/phrased/wordlists"
)

func PickRandomElement(wordlist wordlists.Wordlist) string {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(wordlist.Words))))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return wordlist.Words[n]
}

func Generate(length uint32, id string) (string, error) {
	var result []string
	var wordlist, err = wordlists.FindWordlist(id)
	if err != nil {
		return "", err
	}

	for i := uint32(0); i < length; i++ {
		result = append(result, PickRandomElement(wordlist))
	}

	return strings.Join(result, " "), nil
}
