package main

import (
	"strings"
	"unicode"

	"github.com/extism/go-pdk"
)

type CreatePost struct {
	Name             string  `json:"name"`
	Body             *string `json:"body,omitempty"`
	Community_id     int32   `json:"community_id"`
	Url              *string `json:"url,omitempty"`
	Alt_text         *string `json:"alt_text,omitempty"`
	Honeypot         *string `json:"honeypot,omitempty"`
	Nsfw             *bool   `json:"nsfw,omitempty"`
	Language_id      *int32  `json:"language_id,omitempty"`
	Custom_thumbnail *string `json:"custom_thumbnail,omitempty"`
}

// func main() {
// 	input := "THIS is a TEST message with SOME ALL CAPS words to CHECK the FUNCTIONALITY."
// 	thresholdRatio := 0.5 // 50% ratio
// 	result := checkAndConvertMessage(input, thresholdRatio)
// 	fmt.Println(result)
// }

//export api_before_post_post
func api_before_post_post() int32 {
	params := CreatePost{}
	// use json input helper, which automatically unmarshals the plugin input into your struct
	err := pdk.InputJSON(&params)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	*params.Body = toProperCase(*params.Body)

	// use json output helper, which automatically marshals your struct to the plugin output
	err = pdk.OutputJSON(params)
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	return 0
}

// Function to check if a word is in all caps
func isAllCaps(word string) bool {
	for _, r := range word {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Function to check if a word is a proper noun
func isProperNoun(word string) bool {
	return unicode.IsUpper(rune(word[0]))
}

// Function to convert a string to proper case
func toProperCase(s string) string {
	// Split the string into sentences based on punctuation
	var sentences []string
	var currentSentence strings.Builder
	for _, r := range s {
		currentSentence.WriteRune(r)
		if r == '.' || r == '?' || r == '!' {
			sentences = append(sentences, currentSentence.String())
			currentSentence.Reset()
		}
	}
	if currentSentence.Len() > 0 {
		sentences = append(sentences, currentSentence.String())
	}

	for i, sentence := range sentences {
		words := strings.Fields(sentence)
		if len(words) == 0 {
			continue
		}

		// Capitalize the first letter of the first word of the sentence
		words[0] = strings.Title(strings.ToLower(words[0]))

		for j, word := range words[1:] {
			if isAllCaps(word) {
				words[j+1] = strings.ToLower(word)
			} else if isProperNoun(word) {
				words[j+1] = word
			} else {
				words[j+1] = strings.ToLower(word)
			}
		}

		// Join the words back into a sentence
		sentences[i] = strings.Join(words, " ")
	}

	// Join the sentences back into a single string with proper punctuation
	return strings.Join(sentences, " ")
}

func main() {}
