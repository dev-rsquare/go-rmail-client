package rmail

func (client *Client) CreateTip(tipId int){
	query := `
		mutation sendCreateTipMail($id: Int!){
			send_mail{
			tips{
				add_tip(tip_id: $id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id": tipId,
	}
	client.sendMail(query, vars)
}
