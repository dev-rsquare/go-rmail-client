package rmail

func (client *Client) AddPerson(personId int) {
	query := `
		mutation sendCreatePersonMail($person_id: Int!) {
			send_mail {
				people {
					create_person(person_id: $id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"person_id": personId,
	}
	client.sendMail(query, vars)
}
