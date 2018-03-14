package args_parser

import (
	"errors"
	"regexp"
	"strconv"
)

type PhrasedArgs struct {
	WordlistKey string
	Strength    uint32
	Amount      uint32
}

func parseAsString(arg string) string {
	match, err := regexp.MatchString("([0-9]+)", arg)

	if err == nil && !match {
		return arg
	} else {
		return ""
	}
}

func parseAsInt(arg string) (uint32, error) {
	match, err := regexp.MatchString("([0-9]+)", arg)

	if err == nil && match {
		s, err := strconv.ParseUint(arg, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(s), nil
	} else {
		return 0, errors.New("not a number")
	}
}

func Parse(defaultWordlistKey string, args []string) (PhrasedArgs, error) {
	var defaultStrength uint32 = 5
	var defaultAmount uint32 = 1

	var wordlistKey string
	var strength *uint32 = nil
	var amount *uint32 = nil

	for _, arg := range args {
		parsedWordlist := parseAsString(arg)
		if parsedWordlist != "" {
			if wordlistKey == "" {
				wordlistKey = arg
			} else {
				return PhrasedArgs{}, errors.New("cannot generate passphrase for more than one wordlist")
			}
		}
		parsedNumber, err := parseAsInt(arg)
		if err == nil {
			if strength == nil {
				strength = &parsedNumber
			} else if amount == nil {
				amount = &parsedNumber
			} else {
				return PhrasedArgs{}, errors.New("cannot handle more than two numeric arguments")
			}
		}
	}

	if wordlistKey == "" {
		wordlistKey = defaultWordlistKey
	}
	if strength == nil {
		strength = &defaultStrength
	}
	if amount == nil {
		amount = &defaultAmount
	}

	return PhrasedArgs{Amount: *amount, Strength: *strength, WordlistKey: wordlistKey}, nil
}
