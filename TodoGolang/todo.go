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
	Title     string
	Completed bool
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type Todos []Todo

func (todos *Todos) addTodo(title string){
	todo := Todo{
		Title:  title, 
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int)error{
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int)error{
	t := *todos
	err := t.validateIndex(index)

	if err != nil{
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggleTodo(index int) error{
	t := *todos
	err := t.validateIndex(index)
	if err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if isCompleted {
		completionTime := time.Now()
		t[index].UpdatedAt = &completionTime
	}
	t[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) editTodo(index int, title string)error{
	t := *todos
	err := t.validateIndex(index)

	if err != nil{
		return err
	}
	
	t[index].Title = title

	return nil

}

func (todos *Todos) print(){
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#","Title","Completed","CreatedAt","UpdatedAt")
	for index, t := range *todos {
		completed := "❌"
		updatedAt := ""

		if t.Completed{
			completed = "✅"
			if t.UpdatedAt != nil {
				updatedAt = t.UpdatedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index),t.Title,completed,t.CreatedAt.Format(time.RFC1123),updatedAt)
	}
	table.Render()
}