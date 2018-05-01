package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// Equal checks two values are equal and accepts messages to display if an
// exception occurs
func Equal(
	t *testing.T,
	actual,
	expected interface{},
	messages ...string,
) {
	if actual != expected {
		assertEqualError(t, actual, expected, messages)
	}
}

func assertEqualError(
	t *testing.T,
	actual,
	expected interface{},
	messages []string,
) {
	t.Errorf(MakeEqualTestMessage(actual, expected, messages))
}

// MakeEqualTestMessage returns a nicely formatted equal error string
func MakeEqualTestMessage(
	actual interface{},
	expected interface{},
	messages []string,
) string {
	return fmt.Sprintf(
		"%v\nExpected <%v>%v,\nBut got <%v>%v\n",
		makeMultilineStringOfMessages(messages),
		reflect.TypeOf(expected),
		expected,
		reflect.TypeOf(actual),
		actual,
	)
}

func makeMultilineStringOfMessages(messages []string) string {
	messagesMultilineString := "\n"
	for _, message := range messages {
		messagesMultilineString += string(message) + "\n"
	}
	return messagesMultilineString
}
