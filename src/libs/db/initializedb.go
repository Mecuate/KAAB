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

func Xetup_InternalDB(apisNames []string) error {
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
	err := Xetup_InternalDB(apisNames)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_InternalDB: %v", err))
		errReport = append(errReport, err)
	}
	err = Xetup_Instance(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_Instance: %v", err))
		errReport = append(errReport, err)
	}
	err = Xetup_Media(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_Media: %v", err))
		errReport = append(errReport, err)
	}
	err = Xetup_Files(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_Files: %v", err))
		errReport = append(errReport, err)
	}
	err = Xetup_Users(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_Users: %v", err))
		errReport = append(errReport, err)
	}
	err = Xetup_Endpoints(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_Endpoints: %v", err))
		errReport = append(errReport, err)
	}
	err = Xetup_Nodes(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_Nodes: %v", err))
		errReport = append(errReport, err)
	}
	err = Xetup_Schemas(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_Schemas: %v", err))
		errReport = append(errReport, err)
	}

	/* ----- TEMP DISABLED
	err = Xetup_Accounts(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_Accounts: %v", err))
		errReport = append(errReport, err)
	}
	err = Xetup_DataEntryEvents(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_DataEntryEvents: %v", err))
		errReport = append(errReport, err)
	}
	err = Xetup_KnownHosts(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_KnownHosts: %v", err))
		errReport = append(errReport, err)
	}
	err = Xetup_Stats(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_Stats: %v", err))
		errReport = append(errReport, err)
	}
	err = Xetup_Passwords(databaseName)
	if err != nil {
		config.Err(fmt.Sprintf("[DB_FAILED].Xetup_Passwords: %v", err))
		errReport = append(errReport, err)
	}
	*/

	config.Log("[DB_SUCCESSFUL]Inserted: Xetup_suite")
	return errReport
}

/*
	- individual functions to setup each collection
*/

func Xetup_Schemas(databaseName string) error {
	DB, err := InitMongoDB(databaseName, SCHEMAS)
	if err != nil {
		config.Err("Error building schemas")
		return err
	}
	InternalRegistryData := bson.M{"_name": SCHEMAS}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.schemas already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.schemas")
	}
	baseData := models.CollectionBasis{
		Name:    SCHEMAS,
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [schemas]")
		return err
	}
	return nil
}
func Xetup_Nodes(databaseName string) error {
	DB, err := InitMongoDB(databaseName, NODES)
	if err != nil {
		config.Err("Error building nodes")
		return err
	}
	InternalRegistryData := bson.M{"_name": NODES}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.nodes already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.nodes")
	}
	baseData := models.CollectionBasis{
		Name:    NODES,
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
func Xetup_Endpoints(databaseName string) error {
	DB, err := InitMongoDB(databaseName, ENDPOINTS)
	if err != nil {
		config.Err("Error building endpoints")
		return err
	}
	InternalRegistryData := bson.M{"_name": ENDPOINTS}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.endpoints already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.endpoints")
	}
	baseData := models.CollectionBasis{
		Name:    ENDPOINTS,
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
func Xetup_Accounts(databaseName string) error {
	DB, err := InitMongoDB(databaseName, ACCOUNTS)
	if err != nil {
		config.Err("Error building accounts")
		return err
	}
	InternalRegistryData := bson.M{"_name": ACCOUNTS}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.accounts already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.accounts")
	}
	baseData := models.CollectionBasis{
		Name:    ACCOUNTS,
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
func Xetup_Media(databaseName string) error {
	DB, err := InitMongoDB(databaseName, MEDIA)
	if err != nil {
		config.Err("Error building media")
		return err
	}
	InternalRegistryData := bson.M{"_name": MEDIA}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.media already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.media")
	}
	baseData := models.CollectionBasis{
		Name:    MEDIA,
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
func Xetup_DataEntryEvents(databaseName string) error {
	DB, err := InitMongoDB(databaseName, DATA_ENTRY_EVENTS)
	if err != nil {
		config.Err("Error building data_entry_events")
		return err
	}
	InternalRegistryData := bson.M{"_name": DATA_ENTRY_EVENTS}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.data_entry_events already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.data_entry_events")
	}
	baseData := models.CollectionBasis{
		Name:    DATA_ENTRY_EVENTS,
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
func Xetup_Files(databaseName string) error {
	DB, err := InitMongoDB(databaseName, FILES)
	if err != nil {
		config.Err("Error building files")
		return err
	}
	InternalRegistryData := bson.M{"_name": FILES}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.files already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.files")
	}
	baseData := models.CollectionBasis{
		Name:    FILES,
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
func Xetup_KnownHosts(databaseName string) error {
	DB, err := InitMongoDB(databaseName, KNOWN_HOST)
	if err != nil {
		config.Err("Error building known_host")
		return err
	}
	InternalRegistryData := bson.M{"_name": KNOWN_HOST}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.known_host already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.known_host")
	}
	baseData := models.CollectionBasis{
		Name:    KNOWN_HOST,
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
func Xetup_Passwords(databaseName string) error {
	DB, err := InitMongoDB(databaseName, PASSWORDS)
	if err != nil {
		config.Err("Error building passwords")
		return err
	}
	InternalRegistryData := bson.M{"_name": PASSWORDS}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.passwords already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.passwords")
	}
	baseData := models.CollectionBasis{
		Name:    PASSWORDS,
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
func Xetup_Stats(databaseName string) error {
	DB, err := InitMongoDB(databaseName, STATS)
	if err != nil {
		config.Err("Error building stats")
		return err
	}
	InternalRegistryData := bson.M{"_name": STATS}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.stats already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.stats")
	}
	baseData := models.CollectionBasis{
		Name:    STATS,
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
func Xetup_Users(databaseName string) error {
	DB, err := InitMongoDB(databaseName, USERS)
	if err != nil {
		config.Err("Error building users")
		return err
	}
	InternalRegistryData := bson.M{"_name": USERS}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.users already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.users")
	}
	baseData := models.CollectionBasis{
		Name:    USERS,
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
func Xetup_Instance(databaseName string) error {
	DB, err := InitMongoDB(databaseName, INSTANCE_INFO)
	if err != nil {
		config.Err("Error building Instance")
		return err
	}
	InternalRegistryData := bson.M{"_name": INSTANCE_INFO}
	exist := DB.FindOne(InternalRegistryData)
	if exist != nil {
		config.Err("coll.instanceInfo already exist")
		return fmt.Errorf("[ALREADY_EXIST]:coll.instanceInfo")
	}
	baseData := models.CollectionBasis{
		Name:    INSTANCE_INFO,
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
