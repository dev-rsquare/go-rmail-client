package rmail

// CreateDeal : send mail when created deal
func (client *Client) CreateDeal(dealID int) {
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
		"id": dealID,
	}
	client.sendMail(query, vars)
}

// AddUserToDeal : send mail when added user to deal
func (client *Client) AddUserToDeal(dealID int, userID int) {
	query := `
		mutation sendAddUserToDealMail($deal_id: Int!, $user_id: Int!) {
			send_mail {
				deals {
					add_user_to_deal(deal_id: $deal_id, user_id: $user_id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"deal_id": dealID,
		"user_id": userID,
	}
	client.sendMail(query, vars)
}

// RemoveUserFromDeal : send mail when removed user from deal
func (client *Client) RemoveUserFromDeal(dealID int, userID int) {
	query := `
		mutation sendRemoveUserFromDealMail($deal_id: Int!, $user_id: Int!) {
			send_mail {
				deals {
					remove_user_from_deal(deal_id: $deal_id, user_id: $user_id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"deal_id": dealID,
		"user_id": userID,
	}
	client.sendMail(query, vars)
}

// ChangeDealProgress : send mail when changed deal's progress
func (client *Client) ChangeDealProgress(dealID int, previousProgress string) {
	query := `
		mutation sendChangeDealProgressMail($deal_id: Int!, $previous_progress: String!) {
			send_mail {
				deals {
					change_deal_progress(deal_id: $deal_id, previous_progress: $previous_progress)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"deal_id":           dealID,
		"previous_progress": previousProgress,
	}
	client.sendMail(query, vars)
}

func (client *Client) InformServicesAndManagersTemplate(dealId int, clientId int, personId int) ([]byte, error) {
	query := `
		query ($dealId: Int!, $clientId: Int!, $personId: Int!) {
			template {
				deals {
					inform_services_and_managers(deal_id: $dealId, client_id: $clientId, person_id: $personId) {
						body
					}
				}
			}
		}
	`

	vars := map[string]interface{} {
		"dealId": dealId,
		"clientId": clientId,
		"personId": personId,
	}

	return client.template(query, vars)
}