package blog

import (
	"database/sql"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Server struct {
	db *sql.DB
}

func NewServer(dbpath string) (*Server, error) {
	s := &Server{}
	err := s.initDB(dbpath)
	if err != nil {
		return nil, err
	}
	s.initRoutes()
	return s, nil
}

//go:embed static
var static embed.FS

//go:embed templates/*.tmpl
var tpls embed.FS

var t = template.Must(template.ParseFS(tpls, "templates/*.tmpl"))

func (s *Server) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func (s *Server) initRoutes() {
	// Serve static resources (CSS, etc.)
	http.Handle("/static/", http.FileServer(http.FS(static)))
	// Serve dynamic, server-side rendered pages
	http.HandleFunc("/", s.index)
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	// TODO: pass dynamic data to the template
	var data struct{}

	err := t.ExecuteTemplate(w, "index.tmpl", data)
	if err != nil {
		log.Println("executing index.tmpl:", err)
	}
}
