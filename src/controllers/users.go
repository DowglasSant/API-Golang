package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("register"); erro != nil {
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
	user.ID, erro = repository.Insert(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func ShowAllUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	database, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)
	users, erro := repository.ShowUsers(nameOrNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
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
	user, erro := repository.ShowUserById(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userIdFromToken, erro := authentication.GetUserIdFromToken(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIdFromToken {
		responses.Erro(w, http.StatusForbidden, errors.New("it is not possible to update a user other than your own"))
		return
	}

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

	if erro = user.Prepare("update"); erro != nil {
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
	if erro = repository.UpdateUser(user, userID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
	}

	userIdFromToken, erro := authentication.GetUserIdFromToken(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIdFromToken {
		responses.Erro(w, http.StatusForbidden, errors.New("it is not possible to delete a user other than your own"))
		return
	}

	database, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)
	if erro = repository.DeleteUser(userID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
