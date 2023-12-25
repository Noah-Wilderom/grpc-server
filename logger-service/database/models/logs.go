package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	Info  LogType = "info"
	Error LogType = "error"
)

type LogType string

type Log struct {
	Id        string    `bson:"_id,omitempty" json:"id"`
	Type      LogType   `bson:"type" json:"type"`
	Data      string    `bson:"data" json:"data"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type MongoLogModel struct {
	db   *mongo.Database
	coll string
}

func NewLogModel(db *mongo.Database) *MongoLogModel {
	return &MongoLogModel{
		db:   db,
		coll: "logs",
	}
}

func (m *MongoLogModel) Insert(ctx context.Context, log *Log) error {
	res, err := m.db.Collection(m.coll).InsertOne(ctx, log)
	if err != nil {
		return err
	}

	log.Id = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (m *MongoLogModel) GetAll(ctx context.Context) ([]*Log, error) {
	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	res, err := m.db.Collection(m.coll).Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)

	var logs []*Log
	for res.Next(ctx) {
		var logEntry Log
		err := res.Decode(&logEntry)
		if err != nil {
			log.Fatal(err)
		}

		logs = append(logs, &logEntry)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func (m *MongoLogModel) GetById(ctx context.Context, id string) (*Log, error) {
	return nil, errors.New("TODO")
}

func (m *MongoLogModel) Delete(ctx context.Context, chat *Log) bool {
	return false
}
