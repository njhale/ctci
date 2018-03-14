package urlify

import (
	"bytes"
	"testing"
)

func TestURLify(t *testing.T) {
	phrase := "Mr John Smith"
	urlified := URLify([]byte(phrase))
	expected := []byte("Mr%20John%20Smith")

	if bytes.Compare(urlified, expected) != 0 {
		t.Errorf("%s should equal %s", urlified, expected)
	}
}
