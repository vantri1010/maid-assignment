package transport

import (
	"booking-service/biz"
	"booking-service/dbHandler"
	"booking-service/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func CreateJob(database mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		//db := database.Client()

		var data model.JobCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := dbHandler.NewJobRepo(database, "job")
		biz := biz.NewCreateJobHandler(store)

		if err := biz.CreateAndSend(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Task created successfully",
		})

	}
}
