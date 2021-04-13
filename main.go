// Copyright (c) 2021 Jackie Nguyen. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main is the main file to start the project.
package main

import (
	"net/http"
	"restful-api-golang/pkg/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/todo", handler.GetAllTodo()).Methods(http.MethodGet)
}
