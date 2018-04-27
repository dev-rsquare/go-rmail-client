package rmail

// ChangeTaskEvent : send mail when changed task's event
func (client *Client) ChangeTaskEvent(taskID int, event string) {
	query := `
		mutation sendChangeTaskEventMail($id: Int!, $event: String!){
			send_mail{
				tasks {
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

// TakePhotosTask : send mail when created task for taking photos
func (client *Client) TakePhotosTask(taskID int) {
	query := `
		mutation sendTakePhotosMail($taskId: Int!){
			send_mail{
				tasks {
					take_photos(task_id: $taskId)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"taskId":    taskID,
	}
	client.sendMail(query, vars)
}
