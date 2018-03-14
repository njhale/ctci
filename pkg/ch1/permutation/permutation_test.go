package permutation

import "testing"

func TestIsPermutation(t *testing.T) {
	a := "abc"
	b := "cba"

	perm := IsPermutation(a, b)

	if !perm {
		t.Errorf("%s should be a permutation of %s", b, a)
	}

	a = "abhsasd"
	b = "oisapdfjAA"

	perm = IsPermutation(a, b)

	if perm {
		t.Errorf("%s should not be a permutation of %s", b, a)
	}

}
