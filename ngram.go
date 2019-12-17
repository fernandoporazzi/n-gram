package ngram

// Ngram expects the n-gram value(1, 2, 3 and so on) and the word
func NGram(n int, value string) []string {
        nGrams := []string{}
	var index int
	
	index = len(value) - n + 1
	
	if index < 1 {
		return nGrams
	}

	for i := index - 1; i >= 0; i-- {
		nGrams = append(nGrams, value[i:i+n])
	}
	
	for i, j := 0, len(nGrams) - 1; i < j; i, j = i + 1, j - 1 {
        	nGrams[i], nGrams[j] = nGrams[j], nGrams[i]
    	}
	
	return nGrams
}
