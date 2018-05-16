package rmail

import "fmt"

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
