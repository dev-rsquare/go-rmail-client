package rmail

// CreateTodos : send mail when created todos
func (client *Client) CreateTodos(todoIds []int) {
	query := `
		mutation sendCreateTodosMail($ids: [Int!]!){
			send_mail{
			todos{
				add_todos(todo_ids: $ids)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"ids": todoIds,
	}

	client.sendMail(query, vars)

}

// UpdateTodos : send mail when updated todo.
// oldTodo is the data before it is changed.
// userIds is added user's array of id.
func (client *Client) UpdateTodos(todoID int, oldTodo map[string]interface{}, userIds []int) {
	query := `
		mutation sendUpdateTodosMail($id: Int!, $old_Todo: Todo!, $added_user_id: [Int!]){
			send_mail{
				todos{
					update_todo(todo_id: $id
					old_todo: $old_Todo
					added_user_id: $added_user_id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id":            todoID,
		"old_Todo":      oldTodo,
		"added_user_id": userIds,
	}

	client.sendMail(query, vars)
}
