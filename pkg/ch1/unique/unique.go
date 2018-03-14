package unique

// IsUnique asdf
func IsUnique(s string) bool {
	// initialize map of existant values
	existant := make(map[byte]int)

	// loop through the string entering each character
	for i := 0; i < len(s); i++ {
		c := s[i]
		if existant[c] > 0 { // we've found a duplicate
			return false
		}

		existant[c] = 1
	}

	// there are no duplicates
	return true

}

// IsUniqueWithArray asdf
func IsUniqueWithArray(s string) bool {
	// initialize an array of existant values
	var existant [128]int // extended ASCII

	for i := 0; i < len(s); i++ {
		r := rune(s[i]) // get the rune

		if existant[r] > 0 { // we've found a duplicate
			return false
		}

		existant[r]++ // mark that the character exists
	}
	// there are no duplicates
	return true
}

// IsUniqueWithNoDataStructures asdf
func IsUniqueWithNoDataStructures(s string) bool {
	// search through array for duplicate element
	for i := 0; i < len(s); i++ {
		// search remainder of array for duplicate
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}
