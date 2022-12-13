package blog

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Server struct {
	// TODO: config, state
}

func NewServer() *Server {
	s := &Server{
		// TODO: initialization
	}
	s.initRoutes()
	return s
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
	// TODO: pass dynamic data to the template
	var data struct{}

	err := t.ExecuteTemplate(w, "index.tmpl", data)
	if err != nil {
		log.Println("executing index.tmpl:", err)
	}
}
