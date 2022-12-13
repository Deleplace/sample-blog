package blog_test

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	blog "github.com/Deleplace/sample-blog/go/sample-blog"
)

var s = blog.NewServer()

var host string

func init() {
	go func() {
		err := s.Start() // Will run until the tests are finished
		log.Fatal(err)
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "http://localhost"
	}
	host = addr + ":" + port
}

func TestHomepage(t *testing.T) {
	payload := getAsString(host+"/", t)
	if needle := "Blog app"; !strings.Contains(payload, needle) {
		t.Errorf("Couldn't find expected fragment %q in page payload %q", needle, prefix(payload, 80))
	}
}

func TestCSS(t *testing.T) {
	payload := getAsString(host+"/static/css/style.css", t)
	if needle := ".post > header h1"; !strings.Contains(payload, needle) {
		t.Errorf("Couldn't find expected fragment %q in page payload %q", needle, prefix(payload, 80))
	}
}

func getAsString(url string, t *testing.T) (payload string) {
	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	return string(body)
}

func prefix(s string, n int) string {
	runes := []rune(s)
	if len(runes) < n {
		return s
	}
	return string(runes[:n]) + "..."
}
