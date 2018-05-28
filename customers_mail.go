package rmail

import "fmt"

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

// SleepCustomer : send mail when customer become sleep (dormant)
func (client *Client) SleepCustomer(email string) {
	query := `
		mutation sendSleepCustomerMail($email: String!) {
			send_mail {
				customers {
					sleep_customer(email: $email)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"email": email,
	}

	client.sendMail(query, vars)
}


// AdjustCustomerLevel : send mail when changed customer's level
func (client *Client) AdjustCustomerLevel(clientID int, level int) {
	queryTemplate := `
		mutation sendAdjustCustomerLevel($customer_id: Int!) {
			send_mail {
				customers {
					adjust_customer_level_to_%d(customer_id: $customer_id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"customer_id": clientID,
	}

	query := fmt.Sprintf(queryTemplate, level)

	client.sendMail(query, vars)
}

// TokenCreate : send mail when customer's token created
func (client *Client) TokenCreated(tokenID int) {
	query := `
		mutation sendTokenCreatedMail($tokenID: Int!) {
			send_mail {
				customers {
					token_created(token_id: $tokenID)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"tokenID": tokenID,
	}

	client.sendMail(query, vars)
}

// TokenRenew : send mail when customer's token renew
func (client *Client) TokenRenew(tokenID int) {
	query := `
		mutation sendTokenRenewMail($tokenID: Int!) {
			send_mail {
				customers {
					token_renew(token_id: $tokenID)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"tokenID": tokenID,
	}

	client.sendMail(query, vars)
}

// TokenExpired : send mail when customer's token expired
func (client *Client) TokenExpired(tokenID int) {
	query := `
		mutation sendTokenExpiredMail($tokenID: Int!) {
			send_mail {
				customers {
					token_expired(token_id: $tokenID)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"tokenID": tokenID,
	}

	client.sendMail(query, vars)
}
