// Copyright (c) 2021 Jackie Nguyen. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package tdo contains TOPO object and its application will work.
package tdo

type Todo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
