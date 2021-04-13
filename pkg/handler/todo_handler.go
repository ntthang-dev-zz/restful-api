// Copyright (c) 2021 Jackie Nguyen. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package handler handles the logic of the handle
// functions that will be passed to the routers.
package handler

import (
	"net/http"
	"restful-api-golang/pkg/data"
)

func GetAllTodo(writer http.ResponseWriter, request *http.Request) {
	responseWithJson(writer, http.StatusOK, data.Todos)
}
