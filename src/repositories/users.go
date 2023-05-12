package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repository users) ShowUsers(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, erro := repository.database.Query(
		"select id, user_name, nick, email, created_time from users where user_name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if erro = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedTime,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) ShowUserById(ID uint64) (models.User, error) {
	rows, erro := repository.database.Query(
		"select id, user_name, nick, email, created_time from users where id = ?",
		ID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if erro = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedTime,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repository users) UpdateUser(user models.User, ID uint64) error {
	statement, erro := repository.database.Prepare(
		"update users set user_name = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

func (repository users) DeleteUser(ID uint64) error {
	statement, erro := repository.database.Prepare(
		"delete from users where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repository users) FindByEmail(email string) (models.User, error) {
	row, erro := repository.database.Query("select id, user_password from users where email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if erro = row.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}
