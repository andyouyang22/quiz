# Quiz #

Andrew Ouyang

## Question ##

Given a list of words find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list.

## Solution ##

Make two passes on the input file:

1. store all words in a data structure
2. determine which words are compound

In the first pass, all words are stored in a space-efficient data structure. This solution uses the `github.com/tchap/go-patricia/patricia` trie.

In the second pass, every word is checked to determine if it is compound with respect to the words stored in the trie. A word is compound if

* a substring of the word starting at the beginning of the word is itself a word
* the remainder of the word is either itself a word or another compound word

Note that this recursive check can be skipped if the word is not longer than the current longest compound word.

## Usage ##

The code for the solution is in `quiz/compound.go`. To install the trie dependency, run

	go get -u github.com/tchap/go-patricia/patricia

See https://github.com/tchap/go-patricia for more details. Then, from inside the `quiz` folder, run

	go build compound.go
	./compound [filename]

where filename is the name of the file containing the input words. If no file is passed in, `word.list` will be used by default.
