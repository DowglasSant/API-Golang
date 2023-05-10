package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	database *sql.DB
}

func NewUsersRepository(database *sql.DB) *users {
	return &users{database}
}

func (repository users) Insert(user models.User) (uint64, error) {
	statement, erro := repository.database.Prepare(
		"insert into users (user_name, nick, email, user_password) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, nil
	}

	lastIdInsert, erro := result.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	return uint64(lastIdInsert), nil
}
