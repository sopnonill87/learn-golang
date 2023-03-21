package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// var now time.Time = time.Now()
	// var year int = now.Year()
	// fmt.Printf("%v", year)

	//String replace
	broken := "g# r#cks"
	replacer := strings.NewReplacer("#", "o")
	fmt.Printf("%T\n", replacer)
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	fmt.Print("Enter your grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	if input == "100" {
		fmt.Println("Perfect")
	} else if input >= "60" {
		fmt.Println("You are passed")
	} else {
		fmt.Println("You are failed")
	}

	//Conditionals
	if false {
		fmt.Println("if false")
	} else if true {
		fmt.Println("else if false")

	}

	//Read a file
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
