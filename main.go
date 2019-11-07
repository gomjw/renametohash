package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	// search := flag.String("search", "regex", "Regex to search for")
	files, err := ioutil.ReadDir(".")

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		hasher := sha1.New()
		content, _ := os.Open(file.Name())

		if _, err := io.Copy(hasher, content); err != nil {
			log.Fatal(err)
		}

		content.Close()

		fmt.Printf("%s -> %s\n", file.Name(), hex.EncodeToString(hasher.Sum(nil)))
		err := os.Rename(file.Name(), hex.EncodeToString(hasher.Sum(nil))+path.Ext(file.Name()))

		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
