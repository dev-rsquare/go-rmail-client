package rmail

func (client *Client) ManageProductCreate(productId int) {
	query := `
		mutation sendCreateProductMail($id: Int!) {
			send_mail {
				products {
					create_product(product_id: $id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id": productId,
	}
	client.sendMail(query, vars)
}
