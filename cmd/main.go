package main

import (
	"fmt"
	blogposts "go-tests/reading_files"
	"log"
	"os"
)

func main() {
	post, err := blogposts.NewPostFromFS(os.DirFS("../posts"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(post)
}
