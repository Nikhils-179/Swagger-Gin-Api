package mongo

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
)

func ConnectMongoDB() error {
	uri := "mongodb+srv://Cluster29628:VEhEc3ljaFpj@cluster29628.62ifsp3.mongodb.net/?appName=mongosh+2.2.12"

	dbName := "myDB"
	clientOptions := options.Client().ApplyURI(uri).
	SetServerSelectionTimeout(10*time.Second).
	SetTLSConfig(&tls.Config{InsecureSkipVerify: true})

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB Atlas")
	Client = client
	Database = client.Database(dbName)
	return nil
}

func DisconnectMongoDB() error {
	if Client != nil {
		err := Client.Disconnect(context.TODO())
		if err != nil {
			return err
		}
	}

	fmt.Println("Connection to MongoDB atlas closed")
	return nil
}
