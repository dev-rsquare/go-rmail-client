package rmail

// ChangeTaskEvent : send mail when changed task's event
func (client *Client) ChangeTaskEvent(taskID int, event string) {
	query := `
		mutation sendChangeTaskEventMail($id: Int!, $event: String!){
			send_mail{
			tasks{
				change_task_event(task_id: $id, event: $event)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id":    taskID,
		"event": event,
	}
	client.sendMail(query, vars)
}
