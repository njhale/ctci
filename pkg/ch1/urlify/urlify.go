package urlify

// URLify asdf
func URLify(phrase []byte) []byte {

	// get the number of spaces
	for i := 0; i < len(phrase); i++ {
		b := phrase[i]
		if rune(b) == ' ' {
			end := phrase[i+1:]
			phrase = append(phrase[:i], "%20"...)
			phrase = append(phrase, end...)
			i += 3
		}
	}

	return phrase
}
