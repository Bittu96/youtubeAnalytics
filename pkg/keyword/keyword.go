package keyword

import (
	"fmt"
	"regexp"
	"strings"
)

func Extract(s string) ([]string, error) {
	keyWordsV0 := strings.Fields(s)
	if len(keyWordsV0) == 0 {
		fmt.Println("no words")
		return nil, nil
	}

	// filter out special characters
	var keyWordsV1 []string
	for _, w := range keyWordsV0 {
		if len(w) < 3 {
			continue
		}
		w = regexp.MustCompile("\\.|,|;|!|\\?|\\(|\\)|:|\"|\\^'|\\$|<|>|-|\\/|\\*|“|”|‘|’|”").ReplaceAllString(w, "")
		if len(w) < 3 {
			continue
		}
		var alphanumeric = regexp.MustCompile("^[a-zA-Z0-9_]*$")
		// fmt.Printf("<%v> %v\n", w, alphanumeric.MatchString(w))
		if alphanumeric.MatchString(w) {
			keyWordsV1 = append(keyWordsV1, strings.ToLower(w))
		}
	}
	// fmt.Println("keyWordsV1:", keyWordsV1)

	// filter out stop words
	var keyWordsV2 []string
	for _, w := range keyWordsV1 {
		if matcher(w, stopwords) {
			keyWordsV2 = append(keyWordsV2, w)
		}
	}
	// fmt.Println("keyWordsV2:", keyWordsV2)

	// filter out duplicate words
	keyWordsV3 := unique(keyWordsV2)
	fmt.Println("keyWordsV3:", keyWordsV3)

	// filter out duplicate words
	// keyWordsV4 := unique(keyWordsV2)
	// fmt.Println("keyWordsV3:", keyWordsV3)

	return keyWordsV3, nil
}

func matcher(s string, stopwords []string) bool {
	for _, w := range stopwords {
		if w == s {
			return false
		}
	}
	return true
}

func unique(slice []string) []string {
	uniqMap := make(map[string]int)
	for _, v := range slice {
		uniqMap[v]++
	}

	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}
