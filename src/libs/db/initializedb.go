package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func xetup_InternalDB(apisNames []string) error {
	ctx := context.Background()
	client, err := createClient(ctx)
	if err != nil {
		return err
	}

	IntDbName := config.WEBENV.IntDbName
	conn := client.Database(IntDbName)
	if conn == nil {
		return fmt.Errorf(fmt.Sprintf("[DB_FAILED_CONNECTION] %s connection cannot be stablished to :", IntDbName))
	}

	for _, apiName := range apisNames {
		prevConn := conn.Collection(apiName)
		if prevConn == nil {
			config.Log(fmt.Sprintf("[DB_COLLECTION_CONNECT_FAILED] apis.Internal.Collections.Setup: %s|%s|%v", apiName, IntDbName, prevConn))
			continue
		}

		err := conn.CreateCollection(ctx, apiName)
		if err != nil {
			config.Log(fmt.Sprintf("[DB_FAILED_COLLECTION_CREATION] apis.Internal.Collections.Setup: %s|%s|%v", apiName, IntDbName, err))
			continue
		}
		config.Log(fmt.Sprintf("[DB_NEW_COLLECTION_CREATED]: %s", apiName))

		itemName := fmt.Sprintf("%s:%s", IntDbName, apiName)
		InternalRegistryData := bson.M{"_name": itemName}
		var result bson.M
		conn.Collection(apiName).FindOne(ctx, InternalRegistryData).Decode(&result)
		if result != nil {
			config.Log(fmt.Sprintf("[DB_COLL_ITEM_ALREADY_EXIST]: %v", result))
			continue
		}

		baseData := models.CollectionBasis{
			Name:    itemName,
			Uuid:    uuid.New().String(),
			Created: fmt.Sprintf("%v", time.Now().Unix()),
		}
		res, err := conn.Collection(apiName).InsertOne(ctx, baseData)
		if err != nil {
			config.Err(fmt.Sprintf("[DB] failed insertion of baseData %s", itemName))
			return err
		}
		config.Log(fmt.Sprintf("[DB_SUCCESSFUL] ++++ Inserted: %v", res))
	}
	return nil
}

func DatabaseSetup(databaseName string, apisNames []string) {
	dberr := InitialDataBaseBuild(databaseName, apisNames)
	if dberr != nil {
		var isPanic = false
		for _, itemErr := range dberr {
			if !strings.Contains(itemErr.Error(), "[ALREADY_EXIST]") {
				config.Err(fmt.Sprintf("Error building Initial Data: %v", itemErr))
				isPanic = true
			}
		}
		if isPanic {
			panic(dberr)
		}
	}
}

func InitialDataBaseBuild(databaseName string, apisNames []string) []error {
	errReport := []error{}
	err := xetup_InternalDB(apisNames)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_InternalDB: %v", err))
		errReport = append(errReport, err)
	}
	err = xetup_Instance(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_Instance: %v", err))
		errReport = append(errReport, err)
	}
	err = xetup_Accounts(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_Accounts: %v", err))
		errReport = append(errReport, err)
	}
	err = xetup_Media(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_Media: %v", err))
		errReport = append(errReport, err)
	}
	err = xetup_DataEntryEvents(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_DataEntryEvents: %v", err))
		errReport = append(errReport, err)
	}
	err = xetup_Files(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_Files: %v", err))
		errReport = append(errReport, err)
	}
	err = xetup_KnownHosts(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_KnownHosts: %v", err))
		errReport = append(errReport, err)
	}
	err = xetup_Passwords(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_Passwords: %v", err))
		errReport = append(errReport, err)
	}
	err = xetup_Stats(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_Stats: %v", err))
		errReport = append(errReport, err)
	}
	err = xetup_Users(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_Users: %v", err))
		errReport = append(errReport, err)
	}
	err = xetup_Endpoints(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_Endpoints: %v", err))
		errReport = append(errReport, err)
	}
	err = xetup_Nodes(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].xetup_Nodes: %v", err))
		errReport = append(errReport, err)
	}
	config.Log("[DB_SUCCESSFUL]++Inserted: xetup_suite")
	return errReport
}

/*
	- individual functions to setup each collection
*/

func xetup_Nodes(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "nodes")
	if err != nil {
		config.Err("Error building nodes")
		return err
	}
	InternalRegistryData := bson.M{"_name": "nodes"}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.nodes already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.nodes")
	}
	baseData := models.CollectionBasis{
		Name:    "nodes",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [nodes]")
		return err
	}
	return nil
}
func xetup_Endpoints(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "endpoints")
	if err != nil {
		config.Err("Error building endpoints")
		return err
	}
	InternalRegistryData := bson.M{"_name": "endpoints"}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.endpoints already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.endpoints")
	}
	baseData := models.CollectionBasis{
		Name:    "endpoints",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [endpoints]")
		return err
	}
	return nil
}
func xetup_Accounts(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "accounts")
	if err != nil {
		config.Err("Error building accounts")
		return err
	}
	InternalRegistryData := bson.M{"_name": "accounts"}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.accounts already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.accounts")
	}
	baseData := models.CollectionBasis{
		Name:    "accounts",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [accounts]")
		return err
	}
	return nil
}
func xetup_Media(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "media")
	if err != nil {
		config.Err("Error building media")
		return err
	}
	InternalRegistryData := bson.M{"_name": "media"}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.media already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.media")
	}
	baseData := models.CollectionBasis{
		Name:    "media",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [media]")
		return err
	}
	return nil
}
func xetup_DataEntryEvents(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "data_entry_events")
	if err != nil {
		config.Err("Error building data_entry_events")
		return err
	}
	InternalRegistryData := bson.M{"_name": "data_entry_events"}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.data_entry_events already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.data_entry_events")
	}
	baseData := models.CollectionBasis{
		Name:    "data_entry_events",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [data_entry_events]")
		return err
	}
	return nil
}
func xetup_Files(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "files")
	if err != nil {
		config.Err("Error building files")
		return err
	}
	InternalRegistryData := bson.M{"_name": "files"}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.files already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.files")
	}
	baseData := models.CollectionBasis{
		Name:    "files",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [files]")
		return err
	}
	return nil
}
func xetup_KnownHosts(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "known_host")
	if err != nil {
		config.Err("Error building known_host")
		return err
	}
	InternalRegistryData := bson.M{"_name": "known_host"}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.known_host already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.known_host")
	}
	baseData := models.CollectionBasis{
		Name:    "known_host",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [known_host]")
		return err
	}
	return nil
}
func xetup_Passwords(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "passwords")
	if err != nil {
		config.Err("Error building passwords")
		return err
	}
	InternalRegistryData := bson.M{"_name": "passwords"}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.passwords already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.passwords")
	}
	baseData := models.CollectionBasis{
		Name:    "passwords",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [passwords]")
		return err
	}
	return nil
}
func xetup_Stats(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "stats")
	if err != nil {
		config.Err("Error building stats")
		return err
	}
	InternalRegistryData := bson.M{"_name": "stats"}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.stats already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.stats")
	}
	baseData := models.CollectionBasis{
		Name:    "stats",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [stats]")
		return err
	}
	return nil
}
func xetup_Users(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "users")
	if err != nil {
		config.Err("Error building users")
		return err
	}
	InternalRegistryData := bson.M{"_name": "users"}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.users already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.users")
	}
	baseData := models.CollectionBasis{
		Name:    "users",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [users]")
		return err
	}
	return nil
}
func xetup_Instance(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "instanceInfo")
	if err != nil {
		config.Err("Error building Instance")
		return err
	}
	InternalRegistryData := bson.M{"_name": "instanceInfo"}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.instanceInfo already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.instanceInfo")
	}
	baseData := models.CollectionBasis{
		Name:    "instanceInfo",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [instanceInfo]")
		return err
	}
	return nil
}
