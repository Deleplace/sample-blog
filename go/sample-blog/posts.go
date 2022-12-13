package blog

import (
	"context"
	_ "embed"
	"time"
)

type Post struct {
	Title   string
	Author  string
	Created time.Time
	Body    string
}

//go:embed sql/select_posts.sql
var querySelectPosts string

func (s *Server) posts(ctx context.Context) ([]Post, error) {
	rows, err := s.db.QueryContext(ctx, querySelectPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var p Post
		err = rows.Scan(&p.Author, &p.Created, &p.Title, &p.Body)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}
