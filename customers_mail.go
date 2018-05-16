package rmail

// AddClient : send mail when added client
func (client *Client) FollowNewCustomer(customerID int) {
	query := `
		mutation sendFollowNewCustomerMail($customerId: Int!) {
			send_mail {
				customers {
					follow_new_customer(customer_id: $customerId)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"customer_id": customerID,
	}
	client.sendMail(query, vars)
}