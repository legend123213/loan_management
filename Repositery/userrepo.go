package repositery

import (
	"context"

	domain "github.com/legend123213/loan_managment/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)





type UserServiceRepo struct{
	DBClient *mongo.Database

}

func NewUserServiceRepo(client *mongo.Database) *UserServiceRepo{
	return &UserServiceRepo{
		DBClient: client,
	}
}

func (repo *UserServiceRepo) CreateUser(user *domain.User) (*domain.User, error){
	collection := repo.DBClient.Collection("users")
	Newuser, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	user.ID = Newuser.InsertedID.(primitive.ObjectID).Hex()
	return user, nil
}
func (repo *UserServiceRepo) GetUser(id string) (*domain.User, error){
	collection := repo.DBClient.Collection("users")
	var user domain.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(context.Background(), bson.E{"_id",oid}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo *UserServiceRepo) GetUsers() ([]*domain.User, error){
	collection := repo.DBClient.Collection("users")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var users []*domain.User
	for cursor.Next(context.Background()){
		var user domain.User
		cursor.Decode(&user)
		users = append(users, &user)
	}
	return users, nil
}

func (repo *UserServiceRepo) UpdateUser(user *domain.User) (*domain.User, error){
	collection := repo.DBClient.Collection("users")
	_, err := collection.ReplaceOne(context.Background(), bson.E{"_id", user.ID}, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repo *UserServiceRepo) DeleteUser(id string) error{
	collection := repo.DBClient.Collection("users")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.Background(), bson.E{"_id", oid})
	if err != nil {
		return err
	}
	return nil
}
