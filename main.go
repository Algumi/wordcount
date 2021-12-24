package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	fmt.Println(count(src))
}

// count returns number of words in the string
func count(src string) int {
	return len(strings.Fields(src))
}

// readInput reads source string from command line arguments and returns them
func readInput() (src string, err error) {
	src, err = "", nil
	if len(os.Args) != 2 {
		err = errors.New("incorrect number of cli arguments")
		return
	}
	src = os.Args[1]
	return
}

// fail prints the error and exits
func fail(err error) {
	fmt.Println("wordcount:", err)
	os.Exit(1)
}
