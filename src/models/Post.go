package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID          uint64    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	AuthorID    uint64    `json:"authorId,omitempty"`
	AuthorNick  string    `json:"authorNick,omitempty"`
	Likes       uint64    `json:"likes"`
	CreatedTime time.Time `json:"createdTime,omitempty"`
}

func (post *Post) Prepare() error {
	if erro := post.validate(); erro != nil {
		return erro
	}

	post.Format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("the title is required")
	}

	if post.Content == "" {
		return errors.New("the content is required")
	}

	return nil
}

func (post *Post) Format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
