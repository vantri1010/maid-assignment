package biz

import (
	"booking-service/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateJobHandler interface {
	CreateJob(context context.Context, data *model.JobCreate) (primitive.ObjectID, error)
}

type JobBiz struct {
	repo CreateJobHandler
}

func NewCreateJobHandler(repo CreateJobHandler) *JobBiz {
	return &JobBiz{repo: repo}
}

func (biz *JobBiz) CreateAndSend(context context.Context, data *model.JobCreate) (primitive.ObjectID, error) {
	_id, err := biz.repo.CreateJob(context, data)

	if err != nil {
		return _id, model.ErrCannotCreateEntity("Job", err)
	}

	fmt.Println("InsertedID from job biz: ", _id)
	// TODO: send to maid

	return _id, nil
}
