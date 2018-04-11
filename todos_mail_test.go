package rmail

import "testing"

func TestCreateTodoMail(t *testing.T) {
	client := InitClient("development")
	ids := []int{4}

	client.CreateTodos(ids)
}
