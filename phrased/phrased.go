package phrased

import (
        "strings"
        "errors"
        "crypto/rand"
        "math/big"

        "github.com/phrase-yourself/phrased/wordlists/eff_short_wordlist2"
        "github.com/phrase-yourself/phrased/wordlists/diceware_german"
)

func PickRandomElement(array []string) string {
  nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(array))))
  if err != nil {
    panic(err)
  }
  n := nBig.Int64()
  return array[n]
}

func PickWordlist(name string) ([]string, error) {
  switch name {
    case "eff":
      return eff_short_wordlist2.Wordlist, nil
    case "diceware-german":
      return diceware_german.Wordlist, nil
    default:
      return nil, errors.New("I am really sad :(")
  }
}

func Generate(length uint32, id string) (string, error) {
  var result []string
  var wordlist, err = PickWordlist(id)
  if err != nil {
    return "", err
  }

  for i := uint32(0); i < length; i++ {
    result = append(result, PickRandomElement(wordlist))
  }

  return strings.Join(result, " "), nil
}
