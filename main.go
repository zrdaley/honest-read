package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Phrase struct {
	Phrase string 	`yaml:"phrase"`
	Rephrase string `yaml:"rephrase"`
}

func main() {

	caught := []Phrase{}

	phrases, err := readPhrases()
	if err != nil {
		log.Fatalf("Error reading phrases: %v\n", err)
		return
	}

	path := string(os.Args[1])
	inputFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatalf("Error reading file: %v\n", err)
	}
	
	// Attempt to match each word to a phrase
	words := strings.Fields(string(inputFile))
	for i, word := range words {
		matchedPhrases := phrases[word]
		if len(matchedPhrases) == 0 {
			continue
		}

		for _, phrase := range matchedPhrases {
			phraseWords := strings.Fields(phrase.Phrase)
			for _, w := range phraseWords {
				if w != words[i]{
					break
				}
			}
			caught = append(caught, phrase)
		}
	}
	
	
	fmt.Printf("Caught %d phrases\n", len(caught))

	for _, p := range caught {
		fmt.Printf("Change \"%v\" to \"%v\"\n", p.Phrase, p.Rephrase)
	}
}



func readPhrases() (map[string][]Phrase, error) {
	phrases := make(map[string][]Phrase)

	yamlFile, err := ioutil.ReadFile("phrases.yaml")
	if err != nil {
		return phrases, err
	}

	err = yaml.Unmarshal(yamlFile, &phrases)
    if err != nil {
        return phrases, err
	}
	return phrases, nil
}