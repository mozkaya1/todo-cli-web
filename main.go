package main

import (
	"github.com/mozkaya1/todo-cli/internal"
	// "github.com/mozkaya1/todo-cli/server"
	"github.com/mozkaya1/todo-cli/storage"
)

func main() {
	t := internal.Todos{}
	s := storage.Storage[internal.Todos]{Filename: "todo.json"}
	s.LoadFile(&t)
	// server.RunServer(t)
	F := internal.CmdParse()
	F.Execute(&t)
	s.SaveFile(t)
}
