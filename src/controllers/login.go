package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(body, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	database, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)
	userExistent, erro := repository.FindByEmail(user.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.CheckPassword(userExistent.Password, user.Password); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := authentication.TokenGenerate(userExistent.ID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.Write([]byte(token))
}
