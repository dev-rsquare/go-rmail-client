package rmail

// AddPerson : send mail when added person
func (client *Client) GeneralSend(subject string, body string, recipients []string) {
	query := `
		mutation generalSend($subject: String!, $body: String!, $recipients: [String!]!) {
  			send_mail {
    			send(subject: $subject, body: $body, recipients: $recipients)
  			}
		}
	`
	vars := map[string]interface{}{
		"subject":    subject,
		"body":       body,
		"recipients": recipients,
	}
	client.sendMail(query, vars)
}
