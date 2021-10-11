package service

import (
	"fmt"

	"github.com/We-Code-at-Nights/spotify-insights/src/collection"
	"github.com/We-Code-at-Nights/spotify-insights/src/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type trackService struct {
	conn     *db.Connection
	collName string
}

func NewTrackService(conn *db.Connection, collName string) *trackService {
	return &trackService{conn: conn, collName: collName}
}

func (s *trackService) SearchFromDB(name, artistName string) (collection.Track, error) {
	track := collection.Track{}
	err := s.conn.DB.Collection(s.collName).FindOne(s.conn.Ctx, bson.M{"name": name, "artists.name": artistName}).Decode(&track)

	if err != nil {
		return collection.Track{}, fmt.Errorf("track could not be found in database, given: %s, %s, error: %s", name, artistName, err.Error())
	}

	return track, nil
}

func (s *trackService) Insert(track collection.Track) (*mongo.InsertOneResult, error) {
	if len(track.Artists) == 0 {
		return nil, fmt.Errorf("track cannot have 0 artists, given: %v", track)
	}
	res, err := s.conn.DB.Collection(s.collName).InsertOne(s.conn.Ctx, track)
	if err != nil {
		return nil, fmt.Errorf("track could not be inserted, given: %v, error: %s", track, err.Error())
	}
	return res, nil
}

func (s *trackService) Delete(track collection.Track) error {
	res, err := s.conn.DB.Collection(s.collName).DeleteOne(s.conn.Ctx, track)
	if err != nil {
		return fmt.Errorf("track could not be deleted, given: %v, error: %s", track, err.Error())
	} else if res.DeletedCount == 0 {
		return fmt.Errorf("no match found to delete, given: %v", track)
	}
	return nil
}

func (s *trackService) Update(track collection.Track) error {
	if len(track.Artists) == 0 {
		return fmt.Errorf("track cannot have 0 artists, given: %v", track)
	}
	res, err := s.conn.DB.Collection(s.collName).UpdateOne(s.conn.Ctx, bson.M{"name": track.Name, "artists.name": track.Artists[0].Name}, bson.M{
		"$set": track,
	})

	if err != nil {
		return fmt.Errorf("track could not be updated, given: %v, error: %s", track, err.Error())
	} else if res.MatchedCount == 0 {
		return fmt.Errorf("no match found to update, given: %v", track)
	}

	return nil
}
