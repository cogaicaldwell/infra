package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"coginfra/types"
)

type MongoStorage struct {
	Client *mongo.Client
}

func ConnectToMongo(dsn string) *MongoStorage {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	// utils.Logger.Println("[info] Pinged your deployment. You successfully connected to MongoDB!")
	return &MongoStorage{Client: client}
}

func (s *MongoStorage) GetDocuments(path string) (types.Documents, error) {
	var result types.Documents
	filter := bson.M{"path": path}
	err := s.Client.Database(types.DATABASE_NAME).
		Collection(types.DOCUMENT_COLLECTION).
		FindOne(context.Background(), filter).
		Decode(&result)
	if err != nil {
		return types.Documents{}, err
	}
	return result, nil
}

func (s *MongoStorage) UpdateDocuments(path, doc string) error {
	_, err := s.GetDocuments(path)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	// If path doesn't exist, create a new Documents document
	if err == mongo.ErrNoDocuments {
		newDocuments := types.Documents{
			Path: path,
			Documents: []types.Document{
				{Doc: doc, Id: primitive.NewObjectIDFromTimestamp(time.Now())},
			},
		}
		_, err := s.Client.Database(types.DATABASE_NAME).
			Collection(types.DOCUMENT_COLLECTION).
			InsertOne(context.Background(), newDocuments)
		return err
	}

	// If path exists, append the new document to the existing Documents document
	_, err = s.Client.Database(types.DATABASE_NAME).
		Collection(types.DOCUMENT_COLLECTION).
		UpdateOne(context.Background(), bson.M{"path": path}, bson.M{"$push": bson.M{"documents": types.Document{Doc: doc, Id: primitive.NewObjectIDFromTimestamp(time.Now())}}})

	return err
}

func (s *MongoStorage) DeleteDocuments(key string) error {
	return nil
}

func (s *MongoStorage) ValidateApiKey(apiKey string) (bool, error) {
	var res bson.M

	err := s.Client.Database(types.DATABASE_NAME).
		Collection(types.API_COLLECTION).
		FindOne(context.Background(), bson.M{"_id": apiKey}).
		Decode(&res)
	if err != nil {
		// fmt.Println(err)
		if err == mongo.ErrNoDocuments {
			return false, nil // API key not found
		}
		return false, err // An error occurred during the search
	}

	return true, nil
}

func (s *MongoStorage) CreateAPIKey() (string, error) {
	apiKey := uuid.New()

	apiStruct := types.CreateNewApiKey(apiKey.String())

	d, e := s.Client.Database(types.DATABASE_NAME).
		Collection(types.API_COLLECTION).
		InsertOne(context.Background(), apiStruct)
	if e != nil {
		return "", e
	}
	// fmt.Println(d)
	return apiKey.String(), nil
}

func (s *MongoStorage) Disconnect() error {
	if err := s.Client.Disconnect(context.Background()); err != nil {
		return err
	}
	return nil
}
