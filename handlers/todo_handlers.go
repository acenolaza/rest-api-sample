package handlers

import (
	"net/http"

	"github.com/acenolaza/rest-api-sample/api"
	"github.com/acenolaza/rest-api-sample/core/repo"
	"github.com/acenolaza/rest-api-sample/core/repo/models"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	req := api.Request{Request: r}
	res := api.Response{ResponseWriter: w}

	todo := new(models.Todo)
	if err := req.GetJSONBody(req); err != nil {
		res.RespondWithError(http.StatusUnprocessableEntity, err.Error())
		return
	}

	repo.AddTodo(todo)

	res.RespondWithJSON(http.StatusCreated, todo)
}

func FetchAllTodo(w http.ResponseWriter, r *http.Request) {
	res := api.Response{ResponseWriter: w}

	res.RespondWithJSON(http.StatusOK, repo.TodoList)
}

func FetchSingleTodo(w http.ResponseWriter, r *http.Request) {
	req := api.Request{Request: r}
	res := api.Response{ResponseWriter: w}

	id, err := req.GetVarID()
	if err != nil {
		res.RespondWithError(http.StatusBadRequest, err.Error())
		return
	}

	todo := repo.FindTodo(id)
	if todo.Id > 0 {
		res.RespondWithJSON(http.StatusOK, todo)
		return
	}

	res.RespondWithError(http.StatusNotFound, "Not Found")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	req := api.Request{Request: r}
	res := api.Response{ResponseWriter: w}

	id, err := req.GetVarID()
	if err != nil {
		res.RespondWithError(http.StatusBadRequest, err.Error())
		return
	}

	if err := repo.RemoveTodo(id); err != nil {
		res.RespondWithError(http.StatusNotFound, err.Error())
		return
	}

	res.RespondWithJSON(http.StatusOK, repo.TodoList)
}
