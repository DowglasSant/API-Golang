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

func (repository posts) ShowPosts(userID uint64) ([]models.Post, error) {
	rows, erro := repository.database.Query(
		`select distinct p.*, u.nick from posts p inner join users u on u.id = p.author_id
		 inner join followers f on p.author_id = f.user_id
		 where u.id = ? or f.follower_id = ?
		 order by 1 desc`, userID, userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if erro = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedTime,
			&post.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository posts) UpdatePost(postID uint64, post models.Post) error {
	statement, erro := repository.database.Prepare(
		"update posts set title = ?, content = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(post.Title, post.Content, postID); erro != nil {
		return erro
	}

	return nil
}

func (repository posts) DeletePost(postID uint64) error {
	statement, erro := repository.database.Prepare(
		"delete from posts where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postID); erro != nil {
		return erro
	}

	return nil
}

func (repository posts) FindPostsByUser(userID uint64) ([]models.Post, error) {
	rows, erro := repository.database.Query(`
		select p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		where p.author_id = ?`, userID,
	)
	if erro != nil {
		return nil, erro
	}

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if erro = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedTime,
			&post.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository posts) Like(postId uint64) error {
	statement, erro := repository.database.Prepare("update posts set likes = likes + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postId); erro != nil {
		return erro
	}

	return nil
}

func (repository posts) Unlike(postId uint64) error {
	statement, erro := repository.database.Prepare(`
		update posts set likes =
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE 0
		END
		where id = ?`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postId); erro != nil {
		return erro
	}

	return nil
}
