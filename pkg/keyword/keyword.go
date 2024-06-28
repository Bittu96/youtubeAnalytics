package keyword

import (
	"regexp"
	"strings"
)

func Extract(s string) map[string]struct{} {
	// filter out special characters
	keywordsV1 := filterOutSpecialWords(s)
	if len(keywordsV1) == 0 {
		return nil
	}

	// filter out stop words
	keywordsV2 := filterOutStopWords(keywordsV1)
	if len(keywordsV2) == 0 {
		return nil
	}

	// filter out duplicate keys
	return getUniqueKeys(keywordsV2)
}

func filterOutSpecialWords(text string) (filteredWords []string) {
	special := regexp.MustCompile("\\.|,|;|!|\\?|\\(|\\)|:|\"|\\^'|\\$|<|>|-|\\/|\\*|“|”|‘|’|”")
	text = special.ReplaceAllString(text, " ")

	keywordsV0 := strings.Fields(text)
	if len(keywordsV0) == 0 {
		return nil
	}

	for _, w := range keywordsV0 {
		if len(w) < 3 {
			continue
		}

		alphanumeric := regexp.MustCompile("^[a-zA-Z0-9_]*$")
		if alphanumeric.MatchString(w) {
			filteredWords = append(filteredWords, strings.ToLower(w))
		}
	}

	return
}

func filterOutStopWords(words []string) (filteredWords []string) {
	for _, w := range words {
		if matcher(w, stopwords) {
			filteredWords = append(filteredWords, w)
		}
		// if matcher(w, stopwords) {
		// 	filteredWords = append(filteredWords, w)
		// }
	}

	return
}

func matcher(s string, stopwords []string) bool {
	for _, w := range stopwords {
		if w == s {
			return false
		}
	}

	return true
}

func getUniqueKeys(slice []string) map[string]struct{} {
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	return uniqMap
}
