package internal

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title          string
	Created        string
	Completed      bool
	Completed_time string
	Editing        bool
}

type Todos []Todo

func (T *Todos) Add(title string) {
	todo := Todo{
		Title:          title,
		Created:        time.Now().Format(time.RFC1123),
		Completed:      false,
		Completed_time: "",
	}
	*T = append(*T, todo)
}

func (T *Todos) ValidateIndex(Index int) error {

	if Index < 0 || Index >= len(*T) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}
	return nil

}

func (T *Todos) Delete(Index int) error {
	if err := T.ValidateIndex(Index); err != nil {
		return err

	}
	*T = slices.Delete(*T, Index, Index+1)
	return nil
}

func (T *Todos) Toggle(Index int) error {
	if err := T.ValidateIndex(Index); err != nil {
		return err
	}
	t := *T
	iscompleted := t[Index].Completed
	if !iscompleted {
		completion_time := time.Now().Format(time.RFC1123)
		t[Index].Completed_time = completion_time
	}
	t[Index].Completed = !iscompleted

	return nil
}

func (T *Todos) Edit(Index int, title string) error {
	if err := T.ValidateIndex(Index); err != nil {
		return err
	}
	t := *T
	t[Index].Title = title
	return nil
}

func (T *Todos) Print() error {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created Time", "completed_time")

	for index, t := range *T {
		completed := "⭕"
		completed_at := ""
		if t.Completed {
			completed = "✅"
			completed_at = t.Completed_time
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.Created, completed_at)
	}
	table.Render()
	return nil
}
