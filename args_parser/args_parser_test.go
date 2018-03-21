package args_parser

import "testing"

func TestDefaultValues(t *testing.T) {
	var defaultWordlist = "foo"
	var args, err = Parse(defaultWordlist, []string{})

	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	if args.Amount != 1 {
		t.Errorf("Expected amount of 1 by default, got: %d", args.Amount)
	}
	if args.Strength != 1 {
		t.Errorf("Expected strength of 5 by default, got: %d", args.Amount)
	}
	if args.WordlistKey != defaultWordlist {
		t.Errorf("Expected default wordlist, got: %s", args.WordlistKey)
	}
}

func TestZeroIsNotAccepted(t *testing.T) {
	var _, err = Parse("foo", []string{"0"})
	if err == nil {
		t.Errorf("Expected error")
	}

	_, err = Parse("foo", []string{"1", "0"})
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestParsesSingleWordlistString(t *testing.T) {
	var args, err = Parse("foo", []string{"abcde"})

	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	if args.WordlistKey != "abcde" {
		t.Errorf("Expected parsed wordlist, got: %s", args.WordlistKey)
	}
}

func TestFailsOnMultipleWordlistStrings(t *testing.T) {
	var _, err = Parse("foo", []string{"abcde", "xyz"})

	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestParsesSingleWordlistStringWithNumericArgs(t *testing.T) {
	var args, err = Parse("foo", []string{"1", "abcde", "2"})

	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	if args.WordlistKey != "abcde" {
		t.Errorf("Expected parsed wordlist, got: %s", args.WordlistKey)
	}
}

func TestParsesFirstNumberAsStrength(t *testing.T) {
	var args, err = Parse("foo", []string{"1"})

	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	if args.Strength != 1 {
		t.Errorf("Expected parsed strength to be 1, got: %d", args.Strength)
	}
}

func TestParsesFirstNumberAsStrengthAndSecondAsAmount(t *testing.T) {
	var args, err = Parse("foo", []string{"1", "3"})

	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	if args.Strength != 1 {
		t.Errorf("Expected parsed strength to be 1, got: %d", args.Strength)
	}
	if args.Amount != 3 {
		t.Errorf("Expected parsed amount to be 3, got: %d", args.Amount)
	}
}

func TestFailsOnMoreThanTwoNumbers(t *testing.T) {
	var _, err = Parse("foo", []string{"1", "3", "6"})

	if err == nil {
		t.Errorf("Got error: %s", err)
	}
}

func TestParsesFirstNumberAsStrengthAndSecondAsAmountAndFirstStringAsWordlist(t *testing.T) {
	var args, err = Parse("foo", []string{"1", "bar", "3"})

	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	if args.Strength != 1 {
		t.Errorf("Expected parsed strength to be 1, got: %d", args.Strength)
	}
	if args.Amount != 3 {
		t.Errorf("Expected parsed amount to be 3, got: %d", args.Amount)
	}
	if args.WordlistKey != "bar" {
		t.Errorf("Expected parsed wordlist, got: %s", args.WordlistKey)
	}
}
