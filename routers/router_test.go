package routers

import (
	"encoding/json"
	"fmt"
	"github.com/acenolaza/rest-api-sample/api/parameters"
	"github.com/acenolaza/rest-api-sample/core/repo/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchSingleTodo(t *testing.T) {
	srv := httptest.NewServer(InitRouter())
	defer srv.Close()

	tt := []struct {
		name   string
		todoId string
		status int
	}{
		{name: "Write presentation", todoId: "1", status: http.StatusOK},
		{name: "Host meetup", todoId: "2", status: http.StatusOK},
		{name: "Invalid todoId (Not an int)", todoId: "x", status: http.StatusBadRequest},
		{name: "Invalid todoId (Does not exists in repo)", todoId: "100", status: http.StatusNotFound},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Send GET request using http
			res, err := http.Get(fmt.Sprintf("%s/v1/todos/%s", srv.URL, tc.todoId))
			if err != nil {
				t.Fatalf("could not send GET request: %v", err)
			}

			if res.StatusCode != tc.status {
				t.Errorf("expected status %v; got %v", tc.status, res.StatusCode)
			}

			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			var todo *models.Todo
			if err := json.Unmarshal(body, &todo); err != nil {
				t.Fatalf("could not deserialize body into todo json object: %v", err)
			}

			if todo.Name != "" && todo.Name != tc.name {
				t.Errorf(`expected todo name to be %s, got %s`, tc.name, todo.Name)
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	srv := httptest.NewServer(InitRouter())
	defer srv.Close()

	tt := []struct {
		name   string
		todoId string
		status int
		count  int
		err    string
	}{
		{name: "Invalid todoId (Not an int)", todoId: "x", status: http.StatusBadRequest, err: `strconv.Atoi: parsing "x": invalid syntax`},
		{name: "Invalid todoId (Does not exists in repo)", status: http.StatusNotFound, todoId: "100", err: "could not find Todo with id of 100 to delete"},
		{name: "Write presentation", todoId: "1", status: http.StatusOK, count: 1},
		{name: "Host meetup", todoId: "2", status: http.StatusOK, count: 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1/todos/%s", srv.URL, tc.todoId), nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			// Send DELETE request using httptest
			res, err := srv.Client().Do(req)
			if err != nil {
				t.Fatalf("could not send DELETE request: %v", err)
			}

			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if res.StatusCode != tc.status {
				t.Errorf("expected status %v; got %v", tc.status, res.StatusCode)
			}

			if tc.err != "" {
				var jErr *parameters.JsonError
				if err := json.Unmarshal(body, &jErr); err != nil {
					t.Fatalf("could not deserialize body into JsonErr object: %v", err)
				}

				if jErr.Text != tc.err {
					t.Errorf("expected error message %q; got %q", tc.err, jErr.Text)
				}
				return
			}

			var todos *[]models.Todo
			if err := json.Unmarshal(body, &todos); err != nil {
				t.Fatalf("could not deserialize body into todo json slice: %v", err)
			}

			if len(*todos) != tc.count {
				t.Errorf("expected todo count is %v, got %v", tc.count, len(*todos))
			}
		})
	}
}
