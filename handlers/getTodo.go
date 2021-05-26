package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"goproj/database"
	"net/http"
)

func GetTodo(db database.TodoInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		for i, s := range params {
			fmt.Println(s)
		}

		nums := []int{1,2,3,4}

		for i, k := range nums {
			if k == 4 {
				return pami
			}
		}

		res, err := db.Get(id)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}
}