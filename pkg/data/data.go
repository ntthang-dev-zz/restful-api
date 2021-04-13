// Copyright (c) 2021 Jackie Nguyen. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package data manage data of TODO object in the project.
// This project I did not mention the database connection problem,
// I will create the initial TODO list and manage it through this data.go file.
package data

import "restful-api-golang/pkg/tdo"

var Todos []tdo.Todo

// Function initializes the initial values when the project is started
//  This example is the initialization of a Todo list of 5 elements.
func init() {
	Todos = []tdo.Todo{
		{ID: 1, Name: "Monday", Content: "Learn Maths"},
		{ID: 2, Name: "Tuesday", Content: "Learn Literature"},
		{ID: 3, Name: "Wednesday", Content: "Learn Physics"},
		{ID: 4, Name: "Thursday", Content: "Learn Chemistry"},
		{ID: 5, Name: "Friday", Content: "Learn English"},
	}
}
