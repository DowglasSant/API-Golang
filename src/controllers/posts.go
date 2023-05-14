package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.GetUserIdFromToken(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Post
	if erro = json.Unmarshal(body, &post); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	post.AuthorID = userID

	if erro = post.Prepare(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	database, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()

	repository := repositories.NewPostsRepository(database)
	post.ID, erro = repository.CreatePost(post)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

func ShowPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, erro := strconv.ParseUint(params["postId"], 10, 64)
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

	repository := repositories.NewPostsRepository(database)
	post, erro := repository.FindById(postID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

func ShowPosts(w http.ResponseWriter, r *http.Request) {}

func UpdatePost(w http.ResponseWriter, r *http.Request) {}

func DeletePost(w http.ResponseWriter, r *http.Request) {}
