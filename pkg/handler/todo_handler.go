// Copyright (c) 2021 Jackie Nguyen. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package handler handles the logic of the handle
// functions that will be passed to the routers.
package handler

import (
	"encoding/json"
	"net/http"
	"restful-api-golang/pkg/data"
	"restful-api-golang/pkg/tdo"
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

func CreateTodo(writer http.ResponseWriter, request *http.Request) {
	var newTodo tdo.Todo
	if err := json.NewDecoder(request.Body).Decode(&newTodo); err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Inavalid body"})
		return
	}

	newTodo.ID = generateId(data.Todos)
	data.Todos = append(data.Todos, newTodo)

	responseWithJson(writer, http.StatusCreated, newTodo)
}

func UpdateTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	var updateTodo tdo.Todo
	if err := json.NewDecoder(request.Body).Decode(&updateTodo); err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateTodo.ID = id

	for i, todo := range data.Todos {
		if todo.ID == id {
			data.Todos[i] = updateTodo
			responseWithJson(writer, http.StatusOK, updateTodo)
			return
		}
	}

	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}

func DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	for i, todo := range data.Todos {
		if todo.ID == id {
			data.Todos = append(data.Todos[:i], data.Todos[i+1:]...)
			responseWithJson(writer, http.StatusOK, map[string]string{"message": "Todo was deleted"})
			return
		}
	}

	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}

func generateId(todos []tdo.Todo) int {
	var maxId int
	for _, todo := range todos {
		if todo.ID > maxId {
			maxId = todo.ID
		}
	}

	return maxId + 1
}
