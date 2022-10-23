package main

import (
	"fmt"
	"strings"
)

/**
 * You are given a sentence with no spaces and dictionary containing thousands of words.
 * Write an algorithm to reconstruct the sentence by inserting spaces in the appropriate positions.
 * Example: input: "theskyisblue" output: "the sky is blue" input: "thegrassisgreen" output: "the grass is green"
 */

//the green banana is pale
//

// "thesetrees", ['the', 'these', 'trees']  -> "these trees" -> !'the setrees'

type test struct {
	in  string
	exp string
}

func main() {
	tests := []test{
		{in: "theskyisblue", exp: "the sky is blue"},
		{in: "thegrassisgreen", exp: "the grass is green"},
		{in: "grassesswayinthewind", exp: "grasses sway in the wind"},
	}
	dict := newDict([]string{"a", "the", "these", "that", "green", "sky", "blue", "is", "some", "grasses", "grass", "sway", "wind", "in"})
	for _, test := range tests {
		out := fixSentence([]string{}, test.in, dict)
		if strings.Join(out, " ") != test.exp {
			panic(fmt.Sprintf("out(%s) != exp(%s)", out, test.exp))
		}
		fmt.Printf("%v -> %v\n\n", test.in, out)
	}
}

func fixSentence(matched []string, s string, dict Dict) []string {
	if s == "" {
		return matched
	}
	n := 1
	for n <= len(s) {
		w := s[0:n]
		if ok := dict.Find(s[0:n]); ok {
			fmt.Printf("(%v) \"%s\"(%d)=\"%s\" FOUND\n", matched, s, n, w)
			if sentence := fixSentence(append(matched, w), s[n:], dict); sentence != nil {
				return sentence
			}
		} else {
			fmt.Printf("(%v) \"%s\"(%d)=\"%s\" NOT FOUND\n", matched, s, n, w)
		}
		n++
	}
	return nil
}

//create sorted lists of words with N chars each...
func newDict(words []string) Dict {
	d := Dict{
		words: words,
	}
	// for _, word := range words {
	// 	d.words.Add(word)
	// }
	return d
}

type Dict struct {
	words []string // map[int][]string
}

func (d Dict) Find(ww string) bool {
	for _, w := range d.words {
		if w == ww {
			return true
		}
	}
	return false
}
