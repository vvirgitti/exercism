package bob // package name must match the package name in bob_test.go
import (
	"strings"
)

const testVersion = 2 // same as targetTestVersion

func Hey(input string) string {

	if input == "" {
		return "Fine, be that way"
	}

	if question(input) {
		return "Sure"
	}

	if yelling(input) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func question (sentence string) bool {
	return strings.Contains(sentence, "?")

}

func yelling (sentence string) bool {
	if strings.Contains(sentence, "!") {
		return true
	}

}