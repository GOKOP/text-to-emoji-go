package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println(loadFileToString("dic.list"))
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
