package rmail

func (client *Client) ChangeTaskEvent(taskId int, event string){
	query := `
		mutation sendTodoMail($id: Int!, $event: String!){
			send_mail{
			tasks{
				change_task_event(task_id: $id, event: $event)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id": taskId,
		"event": event,
	}
	client.sendMail(query, vars)
}
