// Package tdo contains TOPO object and its application will work.
package tdo

type Todo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
