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
	addr := os.Getenv("ADDR") + ":" + port
	log.Printf("Listening on %s\n", addr)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) initRoutes() {
	// Serve static resources (CSS, etc.)
	http.Handle("/static/", http.FileServer(http.FS(static)))
	// Serve dynamic, server-side rendered pages
	http.HandleFunc("/", s.index)
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := s.posts(ctx)
	if err != nil {
		log.Println("Reading posts in DB:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error reading posts from the database")
		return
	}
	data := struct {
		Posts []Post
	}{
		Posts: posts,
	}

	err = t.ExecuteTemplate(w, "index.tmpl", data)
	if err != nil {
		log.Println("executing index.tmpl:", err)
	}
}
