package rmail

func (client *Client) CreateTips(tipId int){
	query := `
		mutation sendTodoMail($id: Int!){
			send_mail{
			tips{
				add_tip_to_client(tip_id: $id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id": tipId,
	}
	client.sendMail(query, vars)
}