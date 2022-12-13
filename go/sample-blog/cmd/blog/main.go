package main

import (
	"fmt"
	"log"
	"os"

	blog "github.com/Deleplace/sample-blog/go/sample-blog"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide the SQLite file as argument")
		os.Exit(1)
	}
	dbFilepath := os.Args[1]

	s, err := blog.NewServer(dbFilepath)
	if err != nil {
		log.Fatalln("Creating server:", err)
	}
	err = s.Start()
	log.Fatal(err)
}
