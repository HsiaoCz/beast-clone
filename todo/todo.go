package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreateAt    time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (t *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreateAt:    time.Now(),
	}

	*t = append(*t, todo)
}

func (t *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (t *Todos) delete(index int) error {
	todos := *t
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	*t = append(todos[:index], todos[index+1:]...)
	return nil
}

func (t *Todos) toggle(index int) error {
	todos := *t
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	isCompleted := todos[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		todos[index].CompletedAt = &completionTime
	}
	todos[index].Completed = !isCompleted

	return nil
}
func (t *Todos) edit(index int, title string) error {
	todos := *t
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	todos[index].Title = title

	return nil
}

func (t *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, todo := range *t {
		completed := "❌"
		completedAt := ""

		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreateAt.Format(time.RFC1123), completedAt)

		table.Render()
	}
}
