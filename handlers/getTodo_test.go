package handlers_test

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goproj/handlers"
	"goproj/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodo(t *testing.T) {
	id := AddNewTodo()

	tests := map[string]struct {
		id           string
		expectedCode int
		expected string
	}{
		"should return 200": {
			id:           id,
			expectedCode: 200,
			expected: "learning Golang",
		},
		"should return 200 and modified count 0": {
			id:           id,
			expectedCode: 200,
		},
		"should return 400": {
			id:           "abc",
			expectedCode: 400,
		},
		"should return 404": {
			id: "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest("DELETE", "/todos/"+test.id, nil)
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todos/{id}", handlers.DeleteTodo(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				todo := models.TodoDelete{}
				_ = json.Unmarshal([]byte(rec.Body.String()), &todo)
				assert.Equal(t, test.Tod, todo.DeletedCount)
			}

			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}
	// CleanUp
	_, _ = client.Delete(id)
}
type people struct {
	name string
}

func (p people) Name() {

}

var smth Func