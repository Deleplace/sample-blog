package main

import (
	"log"

	blog "github.com/Deleplace/sample-blog/go/sample-blog"
)

func main() {
	s := blog.NewServer()
	err := s.Start()
	log.Fatal(err)
}
