package generator

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/phrased-org/phrased/wordlists"
	"math"
)

func PickRandomElement(wordlist wordlists.Wordlist) string {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(wordlist.Words))))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return wordlist.Words[n]
}

func Generate(strength uint32, id string) (string, error) {
	var result []string
	var wordlist, err = wordlists.FindWordlist(id)
	if err != nil {
		return "", err
	}
	var entropyOfSingleWord = math.Log2(float64(len(wordlist.Words)))
	var entropyOfRegularPasswordChar = math.Log2(32 + 32 + 10 + 5)
	var requiredEntropy = (entropyOfRegularPasswordChar * 11) + (float64(strength-1) * entropyOfRegularPasswordChar)

	for (float64(len(result)) * entropyOfSingleWord) < requiredEntropy {
		result = append(result, PickRandomElement(wordlist))
	}

	return strings.Join(result, " "), nil
}
