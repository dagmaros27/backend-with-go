package repositories

import (
	"context"
	"net/http"
	"task_managment_api/domain"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) domain.UserRepository {
	return &userRepository{
		collection: db.Collection(domain.CollectionUser),
	}
}

func (us *userRepository) CreateUser(c context.Context, user domain.User) domain.CustomError {

	_, err := us.collection.InsertOne(c, user)
	return domain.CustomError{ErrCode: http.StatusInternalServerError, ErrMessage: err.Error()}
}

func (us *userRepository) GetUserByUsername(c context.Context, username string) (domain.User, domain.CustomError) {
	var user domain.User
	err := us.collection.FindOne(c, bson.M{"username": username}).Decode(&user)
	if err == nil {
		return domain.User{}, domain.CustomError{ErrCode: http.StatusConflict, ErrMessage: "User already exists"}
	}

	if err != mongo.ErrNoDocuments {
		return domain.User{}, domain.CustomError{ErrCode: http.StatusInternalServerError, ErrMessage: err.Error()}
	}

	return user, domain.CustomError{}
}

func (us *userRepository) GetUserByID(c context.Context, id string) (domain.User, domain.CustomError) {
	var user domain.User
	err := us.collection.FindOne(c, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return domain.User{}, domain.CustomError{ErrCode: http.StatusInternalServerError, ErrMessage: err.Error()}
	}
	return user, domain.CustomError{}
}

func (us *userRepository) UpdateUser(c context.Context, user domain.User) domain.CustomError {
	_, err := us.collection.UpdateOne(c, bson.M{"_id": user.ID}, bson.M{"$set": user})
	return domain.CustomError{ErrCode: http.StatusInternalServerError, ErrMessage: err.Error()}
}

func (us *userRepository) GetUserCount(c context.Context) (int64, domain.CustomError) {
	count, err := us.collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return 0, domain.CustomError{ErrCode: http.StatusInternalServerError, ErrMessage: err.Error()}
	}
	return count, domain.CustomError{}
}
