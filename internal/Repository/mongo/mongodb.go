package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
	"time"

	"github.com/AthenaHelali/HTTP-Monitoring/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const UserCollection = "users"

type UserMongodb struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

func NewUserMongoDB(db *mongo.Database, logger *zap.Logger) *UserMongodb {
	return &UserMongodb{
		collection: db.Collection(UserCollection),
		logger:     logger,
	}
}
func (store *UserMongodb) RegisterUser(m model.User) (model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	m.CreatedAt = time.Now()
	if _, err := store.collection.InsertOne(ctx, m); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return model.User{}, Repository.DuplicateUserError{
				ID:      m.ID,
				Message: "user already exist",
			}
		}
		return model.User{}, fmt.Errorf("document creation on mongodb faild %v", err)
	}
	return m, nil

}
func (store *UserMongodb) GetUserByID(id string) (model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.User

	res := store.collection.FindOne(ctx, bson.M{
		"_id": id,
	})
	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, Repository.UserNotFoundError{
				ID: id,
			}
			// return user, fmt.Errorf("user %s doesn't exist. %v", id, err)
		}
		return user, fmt.Errorf("cannot read from collection %v", err)

	}
	if err := res.Decode(&user); err != nil {
		return user, fmt.Errorf("cannot decode result into user %v", err)
	}
	return user, nil

}
func (store *UserMongodb) GetAllUsers() ([]model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := store.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("cannot read from collection %v", err)
	}
	users := make([]model.User, 0)
	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Err(); err != nil {
			return nil, fmt.Errorf("cannot read current cursor from collection %v", err)
		}
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("cannot decode current cursor into user %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (store *UserMongodb) ReplaceUser(m model.User) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := store.collection.DeleteOne(ctx, bson.M{
		"_id": m.ID,
	})
	if err != nil {
		return fmt.Errorf("cannot delete user %s. %v", m.ID, err)
	}
	_, err = store.RegisterUser(m)

	return err
}
