package main

import (
	"fmt"
	"github.com/andreyvit/diff"
	"github.com/onsi/gomega/types"
)

type ProperDiff struct {
	expected string
}

func BeTheSameStringAs(expected string) types.GomegaMatcher {
	return &ProperDiff{
		expected: expected,
	}
}

func (matcher *ProperDiff) Match(actual interface{}) (success bool, err error) {
	success = matcher.expected == actual
	return
}

func (matcher *ProperDiff) FailureMessage(actual interface{}) (message string) {
	actualString := actual.(string)

	message = fmt.Sprintf("Result not as expected:\n%v", diff.CharacterDiff(matcher.expected, actualString))
	return
}

func (matcher *ProperDiff) NegatedFailureMessage(actual interface{}) (message string) {
	// meh?
	return matcher.FailureMessage(actual)
}
