package appCtx

import "go.mongodb.org/mongo-driver/mongo"

type AppContext interface {
	GetDB() *mongo.Database
}

type appCtx struct {
	db *mongo.Database
}

func NewAppContext(db *mongo.Database) *appCtx {
	return &appCtx{db: db}
}

func (ctx *appCtx) GetDB() *mongo.Database {
	return ctx.db
}
