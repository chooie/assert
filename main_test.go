package assert_test

import (
	"strings"
	"testing"

	"github.com/chooie/assert"
)

func TestEqual(t *testing.T) {
	assert.Equal(t, "yo", "yo", "Supports strings")
	assert.Equal(t, 1, 1, "Supports numbers")
	assert.Equal(t, map[string]int{"a": 1}, map[string]int{"a": 1})
}
func TestMakeEqualTestMessageWithNoExtraMessages(t *testing.T) {
	actual := assert.MakeEqualTestMessage(1, 2, []string{})
	expected := "\n\nExpected <int>2,\nBut got <int>1\n"
	errorIfUnequalStrings(t, actual, expected)
}

func TestMakeEqualTestMessageWithExtraMessages(t *testing.T) {
	actual := assert.MakeEqualTestMessage("hi", "bye", []string{"Some message"})
	expected := "\nSome message\n\nExpected <string>bye,\nBut got <string>hi\n"
	errorIfUnequalStrings(t, actual, expected)
}

func errorIfUnequalStrings(t *testing.T, string1, string2 string) {
	if strings.Trim(string1, "") != strings.Trim(string2, "") {
		t.Error("\nExpected:\n" + string1)
		t.Error("\nTo equal:\n" + string2)
	}
}
