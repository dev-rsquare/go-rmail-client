package rmail

// CreateTip : send mail when created tip
func (client *Client) CreateTip(tipID int) {
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
		"id": tipID,
	}
	client.sendMail(query, vars)
}
