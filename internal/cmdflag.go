package internal

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Edit   string
	Delete int
	Toggle int
	List   bool
}

func CmdParse() *CmdFlags {
	F := &CmdFlags{}
	flag.StringVar(&F.Add, "Add", "", "Adding new todo with title..")
	flag.StringVar(&F.Edit, "Edit", "", "Editing task with ID. id:New Title")
	flag.IntVar(&F.Delete, "Delete", -1, "Delete task with ID")
	flag.IntVar(&F.Toggle, "Toggle", -1, "Toggle task with ID")
	flag.BoolVar(&F.List, "List", false, "List all Tasks")

	flag.Parse()
	return F

}

func (F *CmdFlags) Execute(T *Todos) {
	switch {
	case F.List:
		{
			T.Print()
		}
	case F.Add != "":
		{
			T.Add(F.Add)
		}
	case F.Edit != "":
		{
			parts := strings.Split(F.Edit, ":")
			part1, err := strconv.Atoi(parts[0])

			if part1 >= 0 && len(parts) == 2 {
				T.Edit(part1, parts[1])
			} else {
				fmt.Println("Invalid Index / Data edit", err)
				os.Exit(1)
			}

		}
	case F.Delete != -1:
		{
			T.Delete(F.Delete)
		}
	case F.Toggle != -1:
		{
			T.Toggle(F.Toggle)
		}
	default:
		fmt.Println("Invalid Command!!")

	}
}
