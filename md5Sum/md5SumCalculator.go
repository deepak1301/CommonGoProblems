package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')

	if len(text) <= 1 {
		fmt.Println("Input text is NULL ! !")
		fmt.Printf("%x\n", md5.Sum(nil))
	} else {
		fmt.Println("Input text is not NULL ! !")
		data := []byte(text)
		fmt.Printf("%x\n", md5.Sum(data))
	}
}
