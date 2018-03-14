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

func TestIsUniqueWithArray(t *testing.T) {
	s := "abcdefg"
	unique := IsUniqueWithArray(s)

	if !unique {
		t.Errorf("String %s should be unique", s)
	}

	s = "abcdefgg"
	unique = IsUniqueWithArray(s)

	if unique {
		t.Errorf("String %s should not be unique", s)
	}

}

func TestIsUniqueWithNoDataStructures(t *testing.T) {
	s := "abcdefg"
	unique := IsUniqueWithNoDataStructures(s)

	if !unique {
		t.Errorf("String %s should be unique", s)
	}

	s = "abcdefgg"
	unique = IsUniqueWithNoDataStructures(s)

	if unique {
		t.Errorf("String %s should not be unique", s)
	}

}
