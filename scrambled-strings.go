package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	dictionary := flag.String("dictionary", "", "Dictionary file")
	input := flag.String("input", "", "input file")

	flag.Parse()

	dictionaryFile := *dictionary
	inputFile := *input

	dictionaryStrings := readFile(dictionaryFile)
	if checkLimit(dictionaryStrings) {
		log.Fatal("One of the limit is crossed")
	}

	inputStrings := readFile(inputFile)

	for key, inp := range inputStrings {
		count := 0
		for _, dict := range dictionaryStrings {
			count += countWords(dict, inp)
		}

		fmt.Println("Case #", key+1, ":", count)
	}
}

func readFile(dictionaryFile string) (stringSlice []string) {
	file, err := os.Open(dictionaryFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringSlice = append(stringSlice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stringSlice
}

func countWords(dictOrg, inp string) int {
	dictSlice := strings.Split(dictOrg, "")
	wordLen := len(dictSlice)
	regexWord := dictSlice[0]
	for i := 0; i < (wordLen - 2); i++ {
		regexWord += "."
	}

	regexWord += dictSlice[wordLen-1]

	r, _ := regexp.Compile(regexWord)
	regexStrings := r.FindAllString(inp, -1)

	dict := dictOrg[:wordLen-1][1:] // slicing first and last index
	count := 0

	originalForm := make(map[string]bool)

	for _, val := range regexStrings {
		if dictOrg == val {
			if _, ok := originalForm[val]; !ok {
				originalForm[val] = true
				count++
			}
		} else if isScrambled(val, dict) {
			count++
		}
	}

	return count
}

func isScrambled(regexStr string, dictWord string) bool {
	for i := 1; i < len(regexStr)-1; i++ {
		letter := string(regexStr[i])
		index := strings.Index(dictWord, letter)
		if index > -1 {
			// Deleting letter
			dictWord = strings.Replace(dictWord, letter, "", 1)
		} else {
			fmt.Println("Index not found")
			return false
		}
	}

	return true
}

func sliceToMap(sl []string) map[string]bool {
	m := make(map[string]bool)
	for _, val := range sl {
		m[val] = true
	}

	return m
}

func checkLimit(dict []string) bool {
	// checking limit #3
	if len(dict) > 105 {
		return true
	}

	m := make(map[string]bool)

	// checking limit #1
	for _, val := range dict {
		// checking limit #2
		if len(val) < 2 && len(val) > 105 {
			return true
		}
		if _, ok := m[val]; ok {
			return true
		} else {
			m[val] = true
		}
	}

	return false
}
