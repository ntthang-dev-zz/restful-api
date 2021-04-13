// Copyright (c) 2021 Jackie Nguyen. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package handler handles the logic of the handle
// functions that will be passed to the routers.
package handler

import (
	"net/http"
	"restful-api-golang/pkg/data"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllTodo(writer http.ResponseWriter, request *http.Request) {
	responseWithJson(writer, http.StatusOK, data.Todos)
}

func GetTodoById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	for _, todo := range data.Todos {
		if todo.ID == id {
			responseWithJson(writer, http.StatusOK, todo)
			return
		}
	}

	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}
