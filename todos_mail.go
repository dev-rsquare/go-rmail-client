package rmail

func (client *Client) CreateTodos(todoIds []int){
	query := `
		mutation sendTodoMail($ids: [Int!]!){
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

func (client *Client) UpdateTodos(todoId int, oldTodo map[string]interface{}, userIds []int) {
	query := `
		mutation sendTodoMail($id: Int!, $old_Todo: Todo!, $added_user_id: [Int!]){
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
		"id": todoId,
		"old_Todo": oldTodo,
		"added_user_id": userIds,
	}

	client.sendMail(query, vars)
}