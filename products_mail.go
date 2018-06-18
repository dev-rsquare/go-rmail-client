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

// ManageProductUpdate : send mail when updated product
func (client *Client) ManageProductUpdate(oldProduct interface{}) {
	query := `
		mutation sendUpdateProductMail($oldProduct: Product!) {
			send_mail {
				products {
					update_product(old_product: $oldProduct)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"oldProduct": oldProduct,
	}
	client.sendMail(query, vars)
}

// AddProductToFolder : send mail when product put to folder
func (client *Client) AddProductToFolder(folderItemId int) {
	query := `
		mutation sendAddProductToFolderMail($folderItemId: Int!) {
			send_mail {
				folder_items {
					add_product_to_folder(folder_item_id: $folderItemId)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"folderItemId": folderItemId,
	}
	client.sendMail(query, vars)
}
