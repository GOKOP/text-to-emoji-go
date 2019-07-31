package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"errors"
)

func main() {
	dictionary := createDictionary( loadFileToString("dic.list") )

	arg, err := getArgument()
	checkErr(err)

	fmt.Println(toEmoji(arg, dictionary))
}

// don't panic, just close gracefully with message of the error
func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func loadFileToString(path string) string {
	data, err := ioutil.ReadFile("dic.list")
	checkErr(err)
	return string(data)
}

func createDictionary(raw_data string) map[string]string {
	dictionary := make(map[string]string)

	for _, line := range strings.Split(raw_data, "\n") {
		var entry = strings.Split(line, " ")

		if len(entry) >= 2 {
			dictionary[ entry[0] ] = entry[1];
		}
	}

	return dictionary
}

func findMaxKeyLen(dict map[string]string) int {
	maxLen := 0

	for key, _ := range dict {
		if len(key) > maxLen {
			maxLen = len(key)
		}
	}

	return maxLen
}

func getArgument() (string, error) {
	args := os.Args;
	if(len(args) <= 1) {
		return "", errors.New("You must provide an argument.")
	} else {
		return args[1], nil
	}
}

func toEmoji(original string, dictionary map[string]string) string {
	maxLen    := findMaxKeyLen(dictionary)
	converted := ""

	for len(original) > 0 {
		curLen := maxLen
		emoji  := ""

		for curLen > 0 {
			if curLen > len(original) {
				curLen = len(original)
			}

			snippet    := original[:curLen-1]
			match, err := matchSnippet(snippet, dictionary)
			emoji       = match

			if err == nil {
				break;
			}

			curLen -= 1
		}

		if emoji == "" {
			converted += strings.ToUpper( string(original[0]) )
			original = original[1:]
		} else {
			converted += emoji
			original = original[curLen-1:]
		}
	}

	return converted
}

func matchSnippet(snippet string, dictionary map[string]string) (string, error) {
	for key, value := range dictionary {
		if snippet == key {
			return value, nil
		}
	}
	return "", errors.New("No match")
}
