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

	for key, value := range dictionary {
		fmt.Println(key, value)
	}

	fmt.Println("Max key length:", findMaxKeyLen(dictionary))

	arg, err := getArgument()
	checkErr(err)
	fmt.Println(arg)
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
