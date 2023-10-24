package compress

func Count(data []byte) map[byte]int {
	occurrences := make(map[byte]int)
	for _, b := range data {
		occurrences[b]++
	}

	return occurrences
}
