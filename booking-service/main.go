package main

import (
	"maid-assignment/component/appCtx"
)

func main() {
	env := appCtx.NewEnv()
	mongoClient := appCtx.NewMongoDatabase(env)
	defer appCtx.CloseMongoDBConnection(mongoClient)

	//appcontext := appCtx.NewAppContext(mongoClient.Database("booking"))
	//
	//log.Println(appcontext)

}
