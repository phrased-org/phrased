package generator

import (
        "strings"
        "errors"
        "crypto/rand"
        "math/big"

        "github.com/phrase-yourself/phrased/wordlists"
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
  var list = wordlists.Wordlists[name]
  var err error = nil
  if list == nil {
      err = errors.New("I am really sad :(")
  }

  return list, err
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
