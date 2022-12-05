package blog

import (
	"fmt"
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
	http.HandleFunc("/", s.index)
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving request")
	fmt.Fprintln(w, "Hello from Server")
}
