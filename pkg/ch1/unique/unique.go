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
