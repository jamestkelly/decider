package model

import (
	"context"
	"fmt"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

// Item represents a decision item that users can vote on
type Item struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Options []string  `json:"options"`
	Created time.Time `json:"created"`
}

// Vote represents a vote cast by a user on a particular item
type Vote struct {
	ID     string    `json:"id"`
	UserID string    `json:"userId"`
	ItemID string    `json:"itemId"`
	Option string    `json:"option"`
	Casted time.Time `json:"casted"`
}

// User represents a user of the application
type User struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Password   string   `json:"password"`
	AccessCode string   `json:"accessCode"`
	Groups     []string `json:"groups"`
}

// FirebaseDB provides a wrapper for the Firebase Realtime Database
type FirebaseDB struct {
	Client *db.Client
}

// NewFirebaseDB creates a new FirebaseDB instance using the provided Firebase app
func NewFirebaseDB(app *firebase.App) (*FirebaseDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := app.Database(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create realtime database client: %w", err)
	}

	return &FirebaseDB{Client: client}, nil
}

// GetItem returns the Item with the given ID
func (db *FirebaseDB) GetItem(ctx context.Context, itemID string) (*Item, error) {
	var item Item
	err := db.Client.NewRef(fmt.Sprintf("items/%s", itemID)).Get(ctx, &item)
	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	item.ID = itemID
	return &item, nil
}

// GetItems returns all of the Items in the database
func (db *FirebaseDB) GetItems(ctx context.Context) ([]*Item, error) {
	var items []*Item
	err := db.Client.NewRef("items").OrderByKey().Get(ctx, &items)
	if err != nil {
		return nil, fmt.Errorf("failed to get items: %w", err)
	}

	for _, item := range items {
		item.ID = item.ID[len("items/"):]
	}

	return items, nil
}

// CreateItem creates a new Item in the database
func (db *FirebaseDB) CreateItem(ctx context.Context, item *Item) error {
	ref := db.Client.NewRef("items").Push()
	item.ID = ref.Key
	item.Created = time.Now()

	if err := ref.Set(ctx, item); err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	}

	return nil
}

// GetVotesForUser returns all the Votes cast by the user with the given ID
func (db *FirebaseDB) GetVotesForUser(ctx context.Context, userID string) ([]*Vote, error) {
	var votes []*Vote
	err := db.Client.NewRef("votes").OrderByChild("userID").EqualTo(userID).Get(ctx, &votes)
	if err != nil {
		return nil, fmt.Errorf("failed to get votes: %w", err)
	}

	for _, vote := range votes {
		vote.ID = vote.ID[len("votes/"):]
