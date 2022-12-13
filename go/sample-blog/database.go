package blog

import (
	"database/sql"
	_ "embed"
	"log"
	"os"
)

//go:embed sql/schema.sql
var queryCreateTables string

//go:embed sql/testdata.sql
var queryInsertTestData string

func (s *Server) initDB(dbpath string) error {
	// if fileExists(dbpath) {
	// 	log.Println("Using existing DB", dbpath)
	// } else {
	// 	log.Println("Creating DB", dbpath)
	// }

	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return err
	}

	// This always creates new tables from scratch
	log.Println("Creating DB tables...")
	_, err = db.Exec(queryCreateTables)
	if err != nil {
		return err
	}
	log.Println("Inserting test data...")
	_, err = db.Exec(string(queryInsertTestData))
	if err != nil {
		return err
	}
	log.Println("done.")

	s.db = db
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	b := !os.IsNotExist(err)
	return b
}
