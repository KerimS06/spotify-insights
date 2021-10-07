package db

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Ctx    context.Context
	Client *mongo.Client
	DB     *mongo.Database
}

func NewConnection(ctx context.Context, uri, dbName string) (Connection, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return Connection{}, err
	}
	return Connection{Client: client, Ctx: ctx, DB: client.Database(dbName)}, nil
}

func (c *Connection) NewCollection(collName string) (*mongo.Collection, error) {
	err := c.DB.CreateCollection(c.Ctx, collName)
	if err != nil {
		if strings.Contains(err.Error(), "Collection already exists") {
			return c.GetCollection(collName), nil
		} else {
			return nil, err
		}
	}
	return c.GetCollection(collName), nil
}

func (c *Connection) GetCollection(collName string) *mongo.Collection {
	return c.DB.Collection(collName)
}

func (c *Connection) GetDatabaseNames(filter interface{}) ([]string, error) {
	return c.Client.ListDatabaseNames(c.Ctx, filter)
}

func (c *Connection) GetCollectionNames(filter interface{}) ([]string, error) {
	return c.DB.ListCollectionNames(c.Ctx, filter)
}
