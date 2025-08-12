package data

import (
	"context"
	"errors"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"TaskManagerWithMongoDb/models"
)

var userCollection *mongo.Collection

func InitUserCollection(db *mongo.Database) {
	userCollection = db.Collection("users")
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func CreateUser(ctx context.Context, username, password, role string) (models.User, error) {
	var existing models.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&existing)
	if err == nil {
		return models.User{}, errors.New("username already exists")
	}
	hashed, err := HashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: hashed,
		Role:     role,
	}
	_, err = userCollection.InsertOne(ctx, user)
	return user, err
}

func GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func PromoteUser(ctx context.Context, username string) error {
	res, err := userCollection.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"role": "admin"}})
	if err != nil || res.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

func IsFirstUser(ctx context.Context) bool {
	count, _ := userCollection.CountDocuments(ctx, bson.M{})
	return count == 0
}
