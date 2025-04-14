package main

import (
	"github.com/mozkaya1/todo-cli-web/internal"
	"github.com/mozkaya1/todo-cli-web/storage"
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
