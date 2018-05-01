package assert

import (
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
	message :=
		makeMultilineStringOfMessages(messages) + "\n" +
			"Expected <" + reflect.TypeOf(expected).String() + ">'" +
			expected.(string) + "',\n" +
			"But got <" + reflect.TypeOf(actual).String() + ">'" +
			actual.(string) + "'\n"

	t.Errorf(message)
}

func makeMultilineStringOfMessages(messages []string) string {
	messagesMultilineString := "\n"
	for _, message := range messages {
		messagesMultilineString += string(message) + "\n"
	}
	return messagesMultilineString
}
