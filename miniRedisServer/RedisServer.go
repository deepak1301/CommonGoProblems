package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var items = make(map[string]string)
	r := bufio.NewReader(os.Stdin)
	var comm string

	for strings.ToLower(comm) != "exit" {
		fmt.Println("Enter the Command to Proceed : ")
		text, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		words := strings.Fields(text)
		switch strings.ToLower(string(words[0])) {
		default:
			fmt.Println("Unknown Command")
		case "ping":
			fmt.Println("PONG")
		case "set":
			if len(words) != 3 {
				fmt.Println("ERR wrong number of arguments for command")
				return
			}
			items[string(words[1])] = words[2]
			fmt.Println("Value Inserted Successfully ! !")
		case "get":
			if len(words) != 2 {
				fmt.Println("ERR wrong number of arguments for command")
				return
			}
			val, ok := items[string(words[1])]
			if !ok {
				fmt.Println("The value is not present in the Database ")
			} else {
				fmt.Println("The value requested is : " + val)
			}

		case "del":
			if len(words) != 2 {
				fmt.Println("ERR wrong number of arguments for command")
				return
			}
			_, ok := items[string(words[1])]
			if !ok {
				fmt.Println("The value is not present in the Database ")
			} else {
				delete(items, string(words[1]))
				fmt.Println("Value Deleted Successfully ! !")
			}
		case "count":

			val := len(items)
			if val == 0 {
				fmt.Println("No Elements are present in Database ")
			} else {
				fmt.Println("Total number of elements present are : ", val)
			}
		case "exit":
			comm = "exit"
			fmt.Println("Good Bye :) ")
		}
	}
}
