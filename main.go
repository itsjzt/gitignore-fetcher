package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("Please enter the language name!")
	}

	language := os.Args[1]
	fmt.Printf("fetching gitignore for %s\n", language)
	gitignoreURL := "https://raw.githubusercontent.com/github/gitignore/master/" + language + ".gitignore"
	resp, err := http.Get(gitignoreURL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Writing .gitignore")
	WriteToFile(".gitignore", string(body))
}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
