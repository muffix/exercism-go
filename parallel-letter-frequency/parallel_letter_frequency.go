// Package letter contains functionality to count the frequency of letters
package letter

// ConcurrentFrequency concurrently determines the frequency of letters in an array of strings
func ConcurrentFrequency(strings []string) FreqMap {
	freq := FreqMap{}

	c := make(chan FreqMap, len(strings))

	for _, stringToMap := range strings {
		go frequencyCoro(stringToMap, c)
	}

	for range strings {
		for letter, count := range <-c {
			freq[letter] += count
		}
	}

	return freq
}

func frequencyCoro(stringToMap string, c chan FreqMap) {
	c <- Frequency(stringToMap)
}
