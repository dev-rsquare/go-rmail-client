package rmail

import "testing"

func TestCreateTodoMail(t *testing.T) {
	client := InitClient("development")

	client.CreateTodos(4)
}
