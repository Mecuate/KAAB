package db

import (
	"context"
	"fmt"
	"kaab/src/libs/config"
	"kaab/src/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func xetup_InternalDB(apisNames []string) error {
	ctx := context.Background()
	client, err := createClient(ctx)
	if err != nil {
		return err
	}
	IntDbName := config.WEBENV.IntDbName
	cl, _ := client.ListDatabases(ctx, bson.M{})
	conn := client.Database(IntDbName)
	fmt.Println("[CONECTED] dbConn: ", cl)
	// if conn == nil {
	// 	fmt.Println("[F] empty dbConn: ", conn)
	// }

	for _, apiName := range apisNames {
		conn.Collection(apiName)
		err := conn.CreateCollection(ctx, apiName)
		if err != nil {
			fmt.Println("[ALREADY_EXIST] apis.Internal.Collections.Setup: ", apiName, IntDbName)
			continue
		}
		fmt.Println("[NEW COLLECTION CREATED] : ", apiName)
		itemName := fmt.Sprintf("%s:%s", IntDbName, apiName)
		InstData := bson.M{"_name": itemName}
		var result bson.M
		exist := conn.Collection(apiName).FindOne(ctx, InstData).Decode(&result)

		fmt.Println(":   .", result)
		if result != nil {
			fmt.Println(":   .")
			fmt.Println(":   .")
			fmt.Println(":          data coll already exist.", exist)
			config.Err("coll.data already exist")
			continue
		}

		baseData := models.CollectionBasis{
			Name:    itemName,
			Uuid:    uuid.New().String(),
			Created: fmt.Sprintf("%v", time.Now().Unix()),
		}
		res, err := conn.Collection(apiName).InsertOne(ctx, baseData)
		if err != nil {
			config.Err("Error inserting baseData [accounts]")
			return err
		}
		fmt.Println("Inserted: ", res)
	}
	return nil
}

func InitialDataBaseBuild(databaseName string, apisNames []string) error {
	err := xetup_InternalDB(apisNames)
	if err != nil {
		return err
	}
	// err = xetup_Instance(databaseName)
	// if err != nil {
	// 	return err
	// }
	// err = xetup_Accounts(databaseName)
	// if err != nil {
	// 	return err
	// }
	// err = xetup_Audios(databaseName)
	// if err != nil {
	// 	return err
	// }
	// err = xetup_DataEntryEvents(databaseName)
	// if err != nil {
	// 	return err
	// }
	// err = xetup_Files(databaseName)
	// if err != nil {
	// 	return err
	// }
	// err = xetup_KnownHosts(databaseName)
	// if err != nil {
	// 	return err
	// }
	// err = xetup_Medias(databaseName)
	// if err != nil {
	// 	return err
	// }
	// err = xetup_Passwords(databaseName)
	// if err != nil {
	// 	return err
	// }
	// err = xetup_Stats(databaseName)
	// if err != nil {
	// 	return err
	// }
	// err = xetup_Users(databaseName)
	// if err != nil {
	// 	return err
	// }
	// err = xetup_Videos(databaseName)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func xetup_Accounts(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "accounts")
	if err != nil {
		config.Err("Error building accounts")
		return err
	}
	InstData := bson.M{"_name": "accounts"}
	exist := DB.FindOne(InstData)
	if exist != nil {
		fmt.Println("accounts coll already exist.", exist)
		config.Err("coll.accounts already exist")
		return fmt.Errorf("coll.accounts already exist")
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
func xetup_Audios(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "audios")
	if err != nil {
		config.Err("Error building audios")
		return err
	}
	InstData := bson.M{"_name": "audios"}
	exist := DB.FindOne(InstData)
	if exist != nil {
		fmt.Println("audios coll already exist.", exist)
		config.Err("coll.audios already exist")
		return fmt.Errorf("coll.audios already exist")
	}
	baseData := models.CollectionBasis{
		Name:    "audios",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [audios]")
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
	InstData := bson.M{"_name": "data_entry_events"}
	exist := DB.FindOne(InstData)
	if exist != nil {
		fmt.Println("data_entry_events coll already exist.", exist)
		config.Err("coll.data_entry_events already exist")
		return fmt.Errorf("coll.data_entry_events already exist")
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
	InstData := bson.M{"_name": "files"}
	exist := DB.FindOne(InstData)
	if exist != nil {
		fmt.Println("files coll already exist.", exist)
		config.Err("coll.files already exist")
		return fmt.Errorf("coll.files already exist")
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
	InstData := bson.M{"_name": "known_host"}
	exist := DB.FindOne(InstData)
	if exist != nil {
		fmt.Println("known_host coll already exist.", exist)
		config.Err("coll.known_host already exist")
		return fmt.Errorf("coll.known_host already exist")
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
func xetup_Medias(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "medias")
	if err != nil {
		config.Err("Error building medias")
		return err
	}
	InstData := bson.M{"_name": "medias"}
	exist := DB.FindOne(InstData)
	if exist != nil {
		fmt.Println("medias coll already exist.", exist)
		config.Err("coll.medias already exist")
		return fmt.Errorf("coll.medias already exist")
	}
	baseData := models.CollectionBasis{
		Name:    "medias",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [medias]")
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
	InstData := bson.M{"_name": "passwords"}
	exist := DB.FindOne(InstData)
	if exist != nil {
		fmt.Println("passwords coll already exist.", exist)
		config.Err("coll.passwords already exist")
		return fmt.Errorf("coll.passwords already exist")
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
	InstData := bson.M{"_name": "stats"}
	exist := DB.FindOne(InstData)
	if exist != nil {
		fmt.Println("stats coll already exist.", exist)
		config.Err("coll.stats already exist")
		return fmt.Errorf("coll.stats already exist")
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
	InstData := bson.M{"_name": "users"}
	exist := DB.FindOne(InstData)
	if exist != nil {
		fmt.Println("users coll already exist.", exist)
		config.Err("coll.users already exist")
		return fmt.Errorf("coll.users already exist")
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
func xetup_Videos(databaseName string) error {
	DB, err := InitMongoDB(databaseName, "videos")
	if err != nil {
		config.Err("Error building videos")
		return err
	}
	InstData := bson.M{"_name": "videos"}
	exist := DB.FindOne(InstData)
	if exist != nil {
		fmt.Println("videos coll already exist.", exist)
		config.Err("coll.videos already exist")
		return fmt.Errorf("coll.videos already exist")
	}
	baseData := models.CollectionBasis{
		Name:    "videos",
		Uuid:    uuid.New().String(),
		Created: fmt.Sprintf("%v", time.Now().Unix()),
	}
	err = DB.InsertOne(baseData)
	if err != nil {
		config.Err("Error inserting baseData [videos]")
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
	InstData := bson.M{"_name": "instanceInfo"}
	exist := DB.FindOne(InstData)
	if exist != nil {
		fmt.Println("instanceInfo coll already exist.", exist)
		config.Err("coll.instanceInfo already exist")
		return fmt.Errorf("coll.instanceInfo already exist")
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
func SaveInstance(databaseName string, instanceName string) (string, error) {
	DB, err := InitMongoDB(databaseName, "instanceInfo")
	if err != nil {
		config.Err(fmt.Sprintf("Error building Instance [%s] Initial Data: %v", instanceName, err))
		return "", err
	}
	InstData := bson.M{"collection_name": instanceName}

	exist := DB.FindOne(InstData)
	fmt.Println("exist", exist)

	if exist != nil {
		fmt.Println("failed to create; instance already exist.", exist)
		config.Err(fmt.Sprintf("Error building Instance [%s] already exist: %v", instanceName, err))
		return "", err
	}

	instanceID := uuid.New().String()
	newInstanceData := models.InstanceCollection{
		Name:           instanceName,
		Uuid:           instanceID,
		Owner:          "",
		Members:        []string{""},
		Admin:          []string{""},
		EndpointsList:  models.EndpointsCollectionList{},
		SchemasList:    models.SchemasCollectionList{},
		TextFilesList:  models.TextFilesCollectionList{},
		MediaFilesList: models.MediaFilesCollectionList{},
	}
	err = DB.InsertOne(newInstanceData)
	if err != nil {
		config.Err(fmt.Sprintf("-Error saving Instance [%s] Initial Data: %v", instanceName, err))
		return "", err
	}
	return instanceID, err
}
