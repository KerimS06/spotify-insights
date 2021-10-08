package service

import (
	"fmt"

	"github.com/We-Code-at-Nights/spotify-insights/src/collection"
	"github.com/We-Code-at-Nights/spotify-insights/src/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	conn     *db.Connection
	collName string
}

func NewUserService(conn *db.Connection, collName string) *userService {
	return &userService{conn: conn, collName: collName}
}

func (s *userService) FindByName(name string) (collection.User, error) {
	user := collection.User{}
	err := s.conn.DB.Collection(s.collName).FindOne(s.conn.Ctx, collection.User{UserName: name}).Decode(&user)

	if err != nil {
		return collection.User{}, fmt.Errorf("user could not be found, given: %s, error: %s", name, err.Error())
	}

	return user, nil
}

func (s *userService) Insert(u collection.User) (*mongo.InsertOneResult, error) {
	_, err := s.FindByName(u.UserName)

	if err == nil {
		return nil, fmt.Errorf("error:%s, given: %s,", err.Error(), u.UserName)
	}

	res, err := s.conn.DB.Collection(s.collName).InsertOne(s.conn.Ctx, u)

	if err != nil {
		return nil, fmt.Errorf("user could not be inserted, given: %v, error: %s", u, err.Error())
	}

	return res, nil
}

func (s *userService) Delete(name string) error {
	res, err := s.conn.DB.Collection(s.collName).DeleteOne(s.conn.Ctx, collection.User{UserName: name})

	if err != nil {
		return fmt.Errorf("user could not be deleted, given: %s, error: %s", name, err.Error())
	} else if res.DeletedCount == 0 {
		return fmt.Errorf("no match found to delete, given: %s", name)
	}

	return nil
}

func (s *userService) Update(user collection.User) error {
	res, err := s.conn.DB.Collection(s.collName).UpdateOne(s.conn.Ctx, collection.User{UserName: user.UserName}, bson.M{
		"$set": user,
	})

	if err != nil {
		return fmt.Errorf("user could not be updated, given: %s, error: %s", user.UserName, err.Error())
	} else if res.MatchedCount == 0 {
		return fmt.Errorf("no match found to update, given: %s", user.UserName)
	}

	return nil
}
