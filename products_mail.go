package rmail

// ManageProductCreate : send mail when created product
func (client *Client) ManageProductCreate(productID int) {
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
		"id": productID,
	}
	client.sendMail(query, vars)
}

// ManageProductDelete : send mail when deleted product
func (client *Client) ManageProductDelete(productID int) {
	query := `
		mutation sendDeleteProductMail($id: Int!) {
			send_mail {
				products {
					delete_product(product_id: $id)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"id": productID,
	}
	client.sendMail(query, vars)
}
