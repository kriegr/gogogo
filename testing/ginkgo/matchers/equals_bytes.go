package matchers

import (
	"fmt"

	"github.com/onsi/gomega/types"
	"github.com/rotisserie/eris"
)

func EqualsBytes(expected []byte) types.GomegaMatcher {
	return &equalsBytesMatcher{
		expected: expected,
	}
}

type equalsBytesMatcher struct {
	expected    []byte
	actualBytes []byte
	failedIndex int
}

func (matcher *equalsBytesMatcher) Match(actual interface{}) (bool, error) {
	var ok bool
	matcher.actualBytes, ok = actual.([]byte)
	if !ok {
		return false, eris.New("actual needs to be a byte slice.")
	}

	for idx, byte := range matcher.actualBytes {
		if byte != matcher.expected[idx] {
			matcher.failedIndex = idx
			return false, nil
		}
	}
	return true, nil
}

func (m *equalsBytesMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf(
		"Expected byte at index %d to equal %#02x, but it's %#02x",
		m.failedIndex, m.expected[m.failedIndex], m.actualBytes[m.failedIndex],
	)
}

func (m *equalsBytesMatcher) NegatedFailureMessage(actual interface{}) string {
	return "Expected bytes not to be equal"
}
