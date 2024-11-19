package repository

import (
	"context"
	"errors"

	"localloop/services/users/internal/domain/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

// NewMongoRepository creates a new instance of MongoRepository
func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		collection: db.Collection("users"),
	}
}

func (r *MongoRepository) Create(ctx context.Context, user user.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *MongoRepository) FindByEmail(ctx context.Context, email string) (user.User, error) {
	var u user.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return user.User{}, errors.New("user not found")
	}
	return u, err
}

func (r *MongoRepository) Update(ctx context.Context, email string, updates user.UpdateData) error {
	updateData := bson.M{}
	if updates.Name != "" {
		updateData["name"] = updates.Name
	}
	if updates.Hash != "" {
		updateData["password"] = updates.Hash
		updateData["salt"] = updates.Salt
	}

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"email": email},
		bson.M{"$set": updateData},
	)
	return err
}
