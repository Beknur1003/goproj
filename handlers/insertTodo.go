package handlers

import (
	"encoding/json"
	"goproj/database"
	"goproj/models"
	"io/ioutil"
	"net/http"
)

func InsertTodo(db database.TodoInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := models.Todo{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = json.Unmarshal(body, &todo)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		_, err = db.Insert(todo)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}



		WriteResponse(w, http.StatusOK, res)
	}
}

func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
