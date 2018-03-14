package wordlists

import (
	"encoding/json"
	"errors"
	"github.com/phrased-org/phrased/wordlists/wordlists_json"
	"crypto/rand"
	"math/big"
)

type Wordlist struct {
	Key    		string 		`json:"key"`
	Name 		string 		`json:"name"`
	Languages   []string 	`json:"languages"`
	Words	   	[]string 	`json:"words"`
}


func Wordlists() ([]Wordlist, error) {
	var str = wordlists_json.WordlistsJson
	var result []Wordlist
	var err = json.Unmarshal([]byte(str), &result)
	return result, err
}

func RandomWordlist() (Wordlist, error) {
	var wordlists, err = Wordlists()
	if err != nil {
		return Wordlist{}, err
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(wordlists))))
	if err != nil {
		return Wordlist{}, err
	}

	n := nBig.Int64()
	return wordlists[n], nil
}

func FindWordlist(key string) (Wordlist, error){
	var wordlists, err = Wordlists()
	if err != nil {
		return Wordlist{}, err
	}

	for _, wordlist := range wordlists {
		if wordlist.Key == key {
			return wordlist, nil
		}
	}

	return Wordlist{}, errors.New("Could not find wordlist: " + key)
}
