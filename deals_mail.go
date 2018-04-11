package rmail

func (client *Client) CreateDeal(dealId int) {
	query := `
		mutation sendCreateDealMail($id: Int!) {
			send_mail {
				deals {
					create_deal(deal_id: $id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id": dealId,
	}
	client.sendMail(query, vars)
}
