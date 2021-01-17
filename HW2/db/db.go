package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Username        string
	Password        string
	DbName          string
	ClusterEndpoint string
}

const (
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

// Resource :: referene to database
type Resource struct {
	DB *mongo.Database
}

// Close :: to close database connection
func (r *Resource) Close() {
	r.Close()
}

func (c *Connection) CreateConnection() *Resource {
	connectionURI := fmt.Sprintf(connectionStringTemplate, c.Username, c.Password, c.ClusterEndpoint)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {

	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {

	}

	return &Resource{client.Database(c.DbName)}
}
