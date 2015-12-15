package main

import (
	"fmt"
)

var currentId int
var todos Todos

func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d todelete", id)
}

func RepoUpdateTodo(id int, t Todo) Todo {
	todoToUpdate := RepoFindTodo(id)

	if todoToUpdate.Id == 0 {
		return Todo{}
	}
	if todoToUpdate.Completed != t.Completed {
		todoToUpdate.Completed = t.Completed
	}
	if !t.Due.IsZero() && todoToUpdate.Due != t.Due {
		todoToUpdate.Due = t.Due
	}
	if t.Name != "" && todoToUpdate.Name != t.Name {
		todoToUpdate.Name = t.Name
	}

	RepoDestroyTodo(id)
	todos = append(todos, todoToUpdate)

	return todoToUpdate

}
