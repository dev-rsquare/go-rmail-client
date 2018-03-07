package rmail

func (client *Client) CreateClientTip(tipId int){
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

func (client *Client) CreateDealTip(tipId int) {
	query := `
		mutation sendTodoMail($id: Int!){
			send_mail{
			tips{
				add_tip_to_deal(tip_id: $id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id": tipId,
	}
	client.sendMail(query, vars)

}