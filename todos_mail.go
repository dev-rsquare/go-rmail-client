package rmail

func (client *Client) CreateTodos(todoId int) {
	query := `
		mutation sendTodoMail($id: Int!) {
			send_mail {
				covenants {
					create_covenant(covenant_id: $id) {
						shardId
					}
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id": todoId,
	}

	client.sendMail(query, vars)
}
