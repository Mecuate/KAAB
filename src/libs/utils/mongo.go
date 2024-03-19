package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	coll *mongo.Collection
	ctx  context.Context
}

type DBConnection struct {
	db  *DB
	err error
}

func InitMongoDB(dbname string, collname string) (DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	dbconn := make(chan DBConnection)

	go func() {
		client, err := createClient()
		coll := client.Database(dbname).Collection(collname)

		dbconn <- DBConnection{
			db:  &DB{coll, ctx},
			err: err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return DB{}, fmt.Errorf("MongoDB client connection timeout")
		case sucConn := <-dbconn:
			return *sucConn.db, sucConn.err
		}
	}

}

func createClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:7717/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}

func (db DB) InsertOne(data interface{}) {
	_, err := db.coll.InsertOne(db.ctx, data)
	if err != nil {
		panic(err)
	}
}
