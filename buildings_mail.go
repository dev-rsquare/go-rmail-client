package rmail

// AddBuildingToFolder : send mail when building put to folder
func (client *Client) AddBuildingToFolder(folderItemId int) {
	query := `
		mutation sendAddBuildingToFolderMail($folderItemId: Int!) {
			send_mail {
				folder_items {
					add_building_to_folder(folder_item_id: $folderItemId)
				}
			}
		}
	`
	vars := map[string]interface{}{
		"folderItemId": folderItemId,
	}
	client.sendMail(query, vars)
}