package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tchap/go-patricia/patricia"
)

// iter is a wrapper struct for bufio.Reader. This makes creating a new Reader and
// iterating through the file more syntactically-friendly.
type iter struct {
	*bufio.Reader
}

// newIter returns a new iter object for the input file name. It returns an error if
// there is a problem opening the file.
func newIter(fname string) (*iter, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	return &iter{
		bufio.NewReader(f),
	}, nil
}

// next returns the next word in the iter's file.
func (i *iter) next() (string, error) {
	word, err := i.ReadString('\n')
	return strings.TrimRight(word, "\n "), err
}

// compound determines if the input string is compound based on the words stored in
// the input Trie.
func compound(s string, trie *patricia.Trie) bool {
	// Iterate to the second-to-last letter; if no complete word has been found by
	// this point, the word cannot be compound
	for i := 1; i < len(s)-1; i++ {
		// Examine progressively longer substrings starting from the first letter
		// to determine if the input string contains any complete words
		substr := []byte(s[:i])

		// If this substring is a complete word
		if trie.Get(substr) != nil {
			// If the remainder of the string forms a compound word or a complete
			// word, the origin string is a compound word
			if compound(s[i:], trie) || trie.Get([]byte(s[i:])) != nil {
				return true
			}
		}
	}
	return false
}

// longestCompoundWord returns the longest compound word that can be formed using
// the words in the input file. An error is returned if there is a problem opening
// the file. Note that if there is a tie for longest word, only one word will be
// returned.
func longestCompoundWord(fname string) (string, error) {
	trie := patricia.NewTrie()
	words, err := newIter(fname)
	if err != nil {
		return "", err
	}

	// Iterate through the input file and insert all words into the Trie
	word, err := words.next()
	for ; err == nil; word, err = words.next() {
		trie.Insert([]byte(word), struct{}{})
	}

	longest := ""
	words, err = newIter(fname)
	if err != nil {
		return "", err
	}

	// Iterate through the file again and determine if each word is compound
	word, err = words.next()
	for ; err == nil; word, err = words.next() {

		// Only consider words longer than the current longest compound word
		if len(word) > len(longest) && compound(word, trie) {
			longest = word
		}
	}

	return longest, nil
}

func main() {
	var fname string
	args := os.Args
	if len(args) < 2 {
		fname = "word.list"
	} else {
		fname = args[1]
	}
	result, err := longestCompoundWord(fname)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The longest compound string in %s is %q\n", fname, result)
}
