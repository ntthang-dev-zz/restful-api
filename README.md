# Consuming a restful api with go

You will need Go version 1.11+ installed on your development machine. I should use an already live API that we can easily test to see if it works in our browser. I’ll be using the very popular pokeapi which is an API that exposes all the known information for everything Pokemon related. A bit silly I know but it’s a fully fledged API that follows standard naming conventions and requires no authentication so there is no barrier to entry.

## Introduction

Hello everyone, recently, I have just learned and learned more about the Golang language, today I would like to share some knowledge about this language, and in the next part of this article I will also introduce A simple RESTful API example in Golang using Gorilla Mux.

## Directory structure

Below is the directory structure of the project:

```bash
.
├── cmd/
|   └── main.go
├── pkg/
|   ├── data
|   |   └── data.go
|   ├── dto
|   |   └── todo.go
|   └── handler
|       └── todo_handler.go
├── go.mod
└── go.sum
```