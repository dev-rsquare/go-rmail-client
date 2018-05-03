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
func (client *Client) TakePhotosTask(taskIds []int) {
	query := `
		mutation sendTakePhotosMail($taskIds: [Int!]!){
			send_mail{
				tasks {
					take_photos(task_ids: $taskIds)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"taskIds": taskIds,
	}
	client.sendMail(query, vars)
}

// CheckProductsTask : send mail when created task for check products
func (client *Client) CheckProductsTask(taskIds []int) {
	query := `
		mutation sendCheckProductsMail($taskIds: [Int!]!){
			send_mail{
				tasks {
					check_products(task_ids: $taskIds)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"taskIds": taskIds,
	}
	client.sendMail(query, vars)
}

// CheckProductsTask : send mail when created task for check products
func (client *Client) ChangeTaskProgressForWorks(taskId int) {
	query := `
		mutation sendChangeSalesNotifyTaskProgressMail($taskId: Int!){
			send_mail{
				tasks {
					change_sales_notify_task_progress(task_id: $taskId)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"taskId": taskId,
	}
	client.sendMail(query, vars)
}
