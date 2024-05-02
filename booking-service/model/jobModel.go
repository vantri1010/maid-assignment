package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	Pending   = "pending"
	Assigned  = "assigned"
	Completed = "completed"
	Cancelled = "cancelled"
)

type JobCreate struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Desc      string             `bson:"desc" json:"desc"`
	Type      string             `bson:"type" json:"type"`
	Status    string             `bson:"status" json:"status"`
	Addr      string             `bson:"address" json:"address"`
	Duration  int                `bson:"duration" json:"duration"`
	Dom       string             `bson:"domestic_host" json:"domestic_host"`
	Maid      string             `bson:"maid" json:"maid"`
}
