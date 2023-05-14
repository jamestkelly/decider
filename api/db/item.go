package db

import (
	"decider/api/models"
	"decider/api/config"
	"context"
)

func GetItemFromDataStore(groupID, itemID string) (models.Item, error) {
	ctx := context.Background()

	ref := config.FirebaseClient.NewRef("items").Child(groupID).Child(itemID)
	var item models.Item
	if err := ref.Get(ctx, &item); err != nil {
		// Handle error (e.g., item not found)
		return models.Item{}, err
	}

	return item, nil
}