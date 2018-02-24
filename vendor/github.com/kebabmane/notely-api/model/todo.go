package model

import (
	"encoding/json"
	"errors"
)

// FetchAll is the model function which interfaces with the DB and returns a []byte of the category in json format.
func FetchAllTodos() ([]byte, error) {

	var todos []Todo

	db.Find(&todos)

	if len(todos) <= 0 {
		err := errors.New("Not found")
		{
			return []byte("todos not found"), err
		}
	}

	js, err := json.Marshal(todos)
	{
		return js, err
	}
}

// Create creates a new category item and returns the []byte json object and an error.
func CreateTodo(b []byte) ([]byte, error) {

	var todo Todo

	err := json.Unmarshal(b, &todo)

	if err != nil {
		return []byte("Something went wrong"), err
	}

	db.Save(&todo)

	return []byte("Todo successfully created"), nil
}

// FetchSingle gets a single todo based on param passed, returning []byte and error
func FetchSingleTodo(id string) ([]byte, error) {

	var todo Todo
	db.First(&todo, id)

	if todo.ID == 0 {
		err := errors.New("Not found")
		return []byte("todo not found"), err
	}

	js, err := json.Marshal(todo)
	if err != nil {
		js = []byte("Unable to convert todo to JSON format")
	}

	return js, err
}

// Update is the model function for PUT
func UpdateTodo(b []byte, id string) ([]byte, error) {

	var todo, updatedTodo Todo
	db.First(&todo, id)

	if todo.ID == 0 {
		err := errors.New("Not found")
		return []byte("todo not found"), err
	}

	err := json.Unmarshal(b, &updatedTodo)
	if err != nil {
		return []byte("Malformed input"), err
	}

	db.Model(&todo).Update("todo_name", updatedTodo.TodoText)
	db.Model(&todo).Update("todo_due_date", updatedTodo.TodoDueDate)
	db.Model(&todo).Update("todo_user_id", updatedTodo.UserID)
	db.Model(&todo).Update("todo_category", updatedTodo.Category)

	js, err := json.Marshal(&todo)
	if err != nil {
		return []byte("Unable to marshal json"), err
	}

	return js, nil
}

// Delete deletes the categoryo from the database
func DeleteTodo(id string) ([]byte, error) {

	var todo Todo
	db.First(&todo, id)

	if todo.ID == 0 {
		err := errors.New("Not found")
		return []byte("todo not found"), err
	}

	db.Delete(&todo)

	js, err := json.Marshal(&todo)
	if err != nil {
		panic("Unable to marshal todo into json")
	}

	return js, nil
}
