package wordlists

import (
        "github.com/phrased-org/phrased/wordlists/eff_short_wordlist2"
        "github.com/phrased-org/phrased/wordlists/diceware_german"
)

var Wordlists = map[string][]string{
  "eff": eff_short_wordlist2.Wordlist,
  "diceware-german": diceware_german.Wordlist,
}
