package permutation

// IsPermutation asdf
func IsPermutation(a string, b string) bool {
	// declare an array to store the rune difference between the two strings
	var diff [128]int

	for _, r := range a {
		diff[rune(r)]++ // add to the diff
	}

	for _, r := range b {
		diff[rune(r)]-- // subtract from the diff
	}

	return allZero(diff[:])

}

func allZero(a []int) bool {
	for _, v := range a {
		if v != 0 {
			return false
		}
	}

	return true
}
