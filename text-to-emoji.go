package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("dic.list")
	checkErr(err)
	fmt.Print(string(data))
}

// in this project errors are not important
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
