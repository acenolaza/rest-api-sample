package repo

import (
	"fmt"
	"github.com/acenolaza/rest-api-sample/core/repo/models"
)

var currentId int

var TodoList models.Todos

// Some seed data
func init() {
	AddTodo(models.NewTodo("Write presentation"))
	AddTodo(models.NewTodo("Host meetup"))
}

func FindTodo(id int) models.Todo {
	for _, t := range TodoList {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return models.Todo{}
}

func AddTodo(t *models.Todo) {
	currentId += 1
	t.Id = currentId
	TodoList = append(TodoList, *t)
}

func RemoveTodo(id int) error {
	for i, t := range TodoList {
		if t.Id == id {
			TodoList = append(TodoList[:i], TodoList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not find Todo with id of %d to delete", id)
}
