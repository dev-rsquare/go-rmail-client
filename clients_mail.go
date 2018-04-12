package rmail

func (client *Client) AddClient(clientId int) {
	query := `
		mutation sendCreateClientMail($id: Int!) {
			send_mail {
				clients {
					create_client(client_id: $id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id": clientId,
	}
	client.sendMail(query, vars)
}
