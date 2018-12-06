package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	text, err := ioutil.ReadFile("file.txt")
	check(err)

	if len(text) <= 1 {
		fmt.Println("Input text is NULL ! !")
		fmt.Printf("%x\n", md5.Sum(nil))
	} else {
		fmt.Println("Input text is not NULL ! !")
		data := []byte(text)
		fmt.Printf("%x\n", md5.Sum(data))
	}
}
