package repositories

import (
	"api/src/models"
	"database/sql"
)

type posts struct {
	database *sql.DB
}

func NewPostsRepository(database *sql.DB) *posts {
	return &posts{database}
}

func (repository posts) CreatePost(post models.Post) (uint64, error) {
	statement, erro := repository.database.Prepare(
		"insert into posts (title, content, author_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(post.Title, post.Content, post.AuthorID)
	if erro != nil {
		return 0, erro
	}

	lastIdInsert, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIdInsert), nil
}

func (repository posts) FindById(postID uint64) (models.Post, error) {
	row, erro := repository.database.Query(`
		select p.*, u.nick from
		posts p inner join users u
		on u.id = p.author_id where p.id = ?`,
		postID,
	)
	if erro != nil {
		return models.Post{}, erro
	}
	defer row.Close()

	var post models.Post

	if row.Next() {
		if erro = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedTime,
			&post.AuthorNick,
		); erro != nil {
			return models.Post{}, erro
		}
	}

	return post, nil
}
