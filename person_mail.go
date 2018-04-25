package rmail

// AddPerson : send mail when added person
func (client *Client) AddPerson(personID int) {
	query := `
		mutation sendCreatePersonMail($person_id: Int!) {
			send_mail {
				people {
					create_person(person_id: $person_id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"person_id": personID,
	}
	client.sendMail(query, vars)
}
