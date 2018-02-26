package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {

	num_args := len(os.Args)

	if num_args < 3 {
		fmt.Println("Please enter the correct arguments: filename and ngram_len.")
		os.Exit(3)
	}

	filename := os.Args[1]
	ngram_len, err := strconv.Atoi(os.Args[2]) //convert string argument into int

	if ngram_len < 2 || ngram_len > 5 {
		fmt.Println("Please enter an ngram_len of size 2-5.")
		os.Exit(3)
	}

	// fmt.Println(num_args)
	// fmt.Println(filename)
	// fmt.Println(ngram_len)

	inputBytes, err := ioutil.ReadFile(filename) // just pass the file name

	//print errors from user parsing
	if err != nil {
		fmt.Print(err)
	}

	str := string(inputBytes)         // convert content to a 'string'
	lowercase := strings.ToLower(str) //convert string to all lowercase

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")

	noPunctuation := reg.ReplaceAllString(lowercase, " ")

	// fmt.Println(noPunctuation) // print the content without punctuation

	ngramMap := make(map[string]int) //create a map of strings to values

	splitWords := (strings.Split(noPunctuation, " ")) //split content of file into indivudual words

	//iterate over all words in file
	for i := range splitWords {
		processWord(splitWords[i], ngramMap, ngram_len)

	}

	// 	frequentNGrams(40, ngramMap) //shows the 40 most frequent nGrams
	// 	allNGrams(ngramMap)           //shows all nGrams
	nGramAboveThresh(10, ngramMap) //shows the nGrams if they have more than input number of appearances
}

func processWord(word string, wordMap map[string]int, nGramType int) {
	//fmt.Println(word)
	length := len(word)
	//fmt.Println(length)

	switch nGramType {

	case 2:
		{ //bigrams
			for i := 0; i < length-1; i += 2 {
				// fmt.Println(word[i:i+2])
				addGramToMap(word[i:i+2], wordMap)
			}
			for i := 1; i < length-1; i += 2 {
				// fmt.Println(word[i:i+2])
				addGramToMap(word[i:i+2], wordMap)
			}
		}

	case 3:
		{ //trigrams
			for i := 0; i < length-2; i += 3 {
				// fmt.Println(word[i:i+3])
				addGramToMap(word[i:i+3], wordMap)
			}
			for i := 1; i < length-2; i += 3 {
				// fmt.Println(word[i:i+3])
				addGramToMap(word[i:i+3], wordMap)
			}
			for i := 2; i < length-2; i += 3 {
				addGramToMap(word[i:i+3], wordMap)
				// fmt.Println(word[i:i+3])
			}
		}

	case 4:
		{ //4-grams
			for i := 0; i < length-3; i += 4 {
				// fmt.Println(word[i:i+4])
				addGramToMap(word[i:i+4], wordMap)
			}
			for i := 1; i < length-3; i += 4 {
				// fmt.Println(word[i:i+4])
				addGramToMap(word[i:i+4], wordMap)
			}
			for i := 2; i < length-3; i += 4 {
				// fmt.Println(word[i:i+4])
				addGramToMap(word[i:i+4], wordMap)
			}
			for i := 3; i < length-3; i += 4 {
				// fmt.Println(word[i:i+4])
				addGramToMap(word[i:i+4], wordMap)
			}
		}

	case 5:
		{ //5-grams
			for i := 0; i < length-4; i += 5 {
				// fmt.Println(word[i:i+5])
				addGramToMap(word[i:i+5], wordMap)
			}
			for i := 1; i < length-4; i += 5 {
				// fmt.Println(word[i:i+5])
				addGramToMap(word[i:i+5], wordMap)
			}
			for i := 2; i < length-4; i += 5 {
				// fmt.Println(word[i:i+5])
				addGramToMap(word[i:i+5], wordMap)
			}
			for i := 3; i < length-4; i += 5 {
				// fmt.Println(word[i:i+5])
				addGramToMap(word[i:i+5], wordMap)
			}
			for i := 4; i < length-4; i += 5 {
				// fmt.Println(word[i:i+5])
				addGramToMap(word[i:i+5], wordMap)
			}
		}

	default:
		{
			fmt.Println("Please enter an ngram_len of size 2-5.")
			os.Exit(3)
		}
	}
}

func addGramToMap(gram string, wordMap map[string]int) {
	//test if wordMap contains gram already
	value, ans := wordMap[gram]
	//if found, increment gram count
	if ans == true {
		wordMap[gram] = value + 1
	} else {
		//if not found, add to wordMap
		wordMap[gram] = 1
	}
}

//returns the most frequnt nGrams. Number of nGrams shown controlled by numGrams
func frequentNGrams(numGrams int, wordMap map[string]int) {

	//store the values in a slice
	var vals []int
	for key, value := range wordMap {
		key = key
		vals = append(vals, value)
	}

	sort.Ints(vals) //sorts the slice of vals (number of occurrences in input text)

	thresh := vals[len(wordMap)-1-numGrams] //find threshold value

	fmt.Println("thresh=", thresh)

	for key, value := range wordMap {
		//if nGram is above threshold value
		if value > thresh {
			fmt.Println("nGram:\t", key, "\tnumAppearances:\t", value)
		}
	}
}

func allNGrams(wordMap map[string]int) {
	for key, value := range wordMap {
		fmt.Println("nGram:\t", key, "\tnumAppearances:\t", value)
	}
}

func nGramAboveThresh(thresh int, wordMap map[string]int) {
	for key, value := range wordMap {
		//if nGram is above threshold value
		if value > thresh {
			fmt.Println("nGram:\t", key, "\tnumAppearances:\t", value)
		}
	}
}
