package rmail

// AddClient : send mail when added client
func (client *Client) AddClient(clientID int) {
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
		"id": clientID,
	}
	client.sendMail(query, vars)
}
