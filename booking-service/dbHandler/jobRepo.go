package dbHandler

import (
	"booking-service/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type jobRepository struct {
	database   mongo.Database
	collection string
}

func NewJobRepo(db mongo.Database, collection string) *jobRepository {
	return &jobRepository{
		database:   db,
		collection: collection,
	}
}

func (repo *jobRepository) CreateJob(context context.Context, data *model.JobCreate) (primitive.ObjectID, error) {
	collection := repo.database.Collection(repo.collection)

	insertResult, err := collection.InsertOne(context, data)
	_id := insertResult.InsertedID.(primitive.ObjectID)

	fmt.Println("InsertedID from jobRepo", insertResult.InsertedID)

	return _id, err
}

func (repo *jobRepository) SendJob(context context.Context, data *model.JobCreate) {

}
