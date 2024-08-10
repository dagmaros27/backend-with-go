package data

import (
	"context"
	"errors"
	"task_managment_api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecret string = "samketnon"

type Claims struct {
	UserId string `json:"userId"`
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.StandardClaims
}


type UserService interface {
	CreateUser(user models.User) error
	AuthenticateUser( username, password string) (string, error)
	PromoteUser(userID string) error
}


type MongoUserService struct {
	collection *mongo.Collection
}

func NewMongoUserService(db *mongo.Database) *MongoUserService {
	return &MongoUserService{
		collection: db.Collection("users"),
	}
}

func (us *MongoUserService) CreateUser(user models.User) error {
	filter := bson.M{"username": user.Username}
	var existingUser models.User
	err := us.collection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err == nil {
		return errors.New("user already exists")
	}

	if err != mongo.ErrNoDocuments {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	count, err := us.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	_, err = us.collection.InsertOne(context.TODO(), user)
	return err
}

func (us *MongoUserService) AuthenticateUser( username, password string) (string, error) {
	var user models.User
	err := us.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId:  user.ID,
		Username: username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (us *MongoUserService) PromoteUser(username string) error {
	result, err := us.collection.UpdateOne(context.TODO(), bson.M{"username": username}, bson.M{"$set": bson.M{"role": "admin"}})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
