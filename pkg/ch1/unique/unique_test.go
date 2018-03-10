package unique

import "testing"

func TestIsUnique(t *testing.T) {
	s := "abcdefg"
	unique := IsUnique(s)

	if !unique {
		t.Errorf("String %s should be unique", s)
	}

	s = "abcdefgg"
	unique = IsUnique(s)

	if unique {
		t.Errorf("String %s should not be unique", s)
	}

}
