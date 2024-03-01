package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ctx := context.Background()
	client, err := createClient(ctx)
	if err != nil {
		return DB{}, err
	}
	coll, err := doesCollectionExist(client, dbname, collname)
	if err != nil {
		config.Err(fmt.Sprintf("Error connecting to collection: %v", err))
		return DB{}, err
	}
	return DB{coll, ctx}, err
}

func doesCollectionExist(client *mongo.Client, dbname string, collname string) (*mongo.Collection, error) {
	coll := client.Database(dbname).Collection(collname)
	if coll == nil {
		client.Disconnect(context.Background())
		return nil, fmt.Errorf("collection %s does not exist", collname)
	}
	return coll, nil
}

func createClient(ctx context.Context) (*mongo.Client, error) {
	tM := reflect.TypeOf(bson.M{})
	reg := bson.NewRegistry()
	reg.RegisterTypeMapEntry(bsontype.EmbeddedDocument, tM)
	clientOptions := options.Client().ApplyURI(config.WEBENV.Mongodburi).SetRegistry(reg)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, err
	/*
			 tM := reflect.TypeOf(bson.M{})
		    reg := bson.NewRegistryBuilder().RegisterTypeMapEntry(bsontype.EmbeddedDocument, tM).Build()
		    clientOpts := options.Client().ApplyURI(SOMEURI).SetAuth(authVal).SetRegistry(reg)
		    client, err := mongo.Connect(ctx, clientOpts)
	*/
}

func (db DB) InsertOne(payload interface{}) error {
	_, err := db.coll.InsertOne(db.ctx, payload)
	if err != nil {
		config.Err(fmt.Sprintf("Error inserting data: %v", err))
	}
	return err
}

func (db DB) FindOne(query interface{}) primitive.M {
	var result bson.M
	err := db.coll.FindOne(
		context.TODO(),
		query,
	).Decode(&result)

	if err != nil {
		config.Err(fmt.Sprintf("Error finding one: %v", err))
	}

	return result
}

func InitMongoDBWithTimeOut(dbname string, collname string) (DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	dbconn := make(chan DBConnection)

	go func() {
		client, err := createClient(ctx)
		coll := client.Database(dbname).Collection(collname)

		dbconn <- DBConnection{
			db:  &DB{coll, ctx},
			err: err,
		}
	}()
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return DB{}, fmt.Errorf("MongoDB client connection timeout")
		case sucConn := <-dbconn:
			return *sucConn.db, sucConn.err
		}
	}
}
