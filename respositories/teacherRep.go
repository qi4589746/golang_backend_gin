package respositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"golang_backend_gin/agents"
	"golang_backend_gin/models"
	"log"
)

func CreateTeacher(teacher models.Teacher) (*mongo.InsertOneResult, error) {
	log.Println("CreateTeacher error", agents.MongoDB)
	insertResult, err := agents.MongoDB.Database("goDatabase").Collection("teacher").InsertOne(context.TODO(), teacher)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Inserted post with ID:", insertResult.InsertedID)
	return insertResult, err
}
