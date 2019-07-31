package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dictionary := createDictionary( loadFileToString("dic.list") )

	for key, value := range dictionary {
		fmt.Println(key, value)
	}

	fmt.Println("Max key length:", findMaxKeyLen(dictionary))
}

// in this project errors are not important
func checkErr(e error) {
	if e != nil {
		panic(e)
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
	maxlen := 0

	for key, _ := range dict {
		if len(key) > maxlen {
			maxlen = len(key)
		}
	}

	return maxlen
}
