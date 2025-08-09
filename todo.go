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
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("index out of range")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	t[index].CompletedAt = nil
	if !isCompleted {

		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	t[index].Title = title
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	table.SetRowLines(false)

	for index, todo := range *todos {
		// completed := "❌"

		completed := "No"
		if todo.Completed {
			completed = "Yes"
			// completed = "✅"
		}
		createdAt := todo.CreatedAt.Format("2006-01-02 15:04:05")

		completedAt := "N/A"
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format(time.RFC1123)
		}

		table.AddRow(strconv.Itoa(index), todo.Title, completed, createdAt, completedAt)
	}
	table.Render()
}
