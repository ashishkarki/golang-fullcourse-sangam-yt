package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repository : Repo is our data access layer where DB access happens
type Repository struct {
	coll *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repository {
	return &Repository{
		coll: db.Collection("notes"),
	}
}

func (repo *Repository) CreateNote(ctx context.Context, note Note) (Note, error) {
	// ctx above is parent context whereas childCtx below is child context with additional 5 secs
	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	_, err := repo.coll.InsertOne(childCtx, note)
	if err != nil {
		return Note{}, fmt.Errorf("failed to insert note: %w", err)
	}

	return note, nil
}
