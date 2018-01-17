package handlers

import (
	"encoding/json"
	"github.com/acenolaza/rest-api-sample/core/repo/models"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	Index(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	t.Log("Index test passed.")
}

func TestFetchAllTodo(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/v1/todos", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	FetchAllTodo(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			t.Fatalf("unexpected error closing the body: %v", err)
		}
	}()

	body, err := ioutil.ReadAll(io.LimitReader(res.Body, 1048576))
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
	if len(body) == 0 {
		t.Error("body was empty")
	}

	var todoList *[]models.Todo
	if err := json.Unmarshal(body, &todoList); err != nil {
		t.Fatalf("could not deserialize body into todo json slice: %v", err)
	}

	if len(*todoList) != 2 {
		t.Errorf("expected todo count is 2, got %v", len(*todoList))
	}

	t.Log("FetchAllTodo test passed.")
}
