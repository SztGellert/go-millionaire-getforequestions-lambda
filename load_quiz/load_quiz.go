package load_quiz

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func ConnectMongo() {

	// test db with private network access
	credential := options.Credential{
		Username: "admin",
		Password: "admin",
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://172.31.21.185:27017").SetAuth(credential))
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err)
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println(err.Error())
		log.Println(err)
		panic(err)
	}
}

func LoadQuestion() ([]Question, error) {

	credential := options.Credential{
		Username: "admin",
		Password: "admin",
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	mongourl := "mongodb://172.31.21.185:27017"
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongourl).SetAuth(credential))
	if err != nil {
		return nil, err
	}

	foreQuestionsCollection := client.Database("quiz").Collection("fore_questions")
	var foreQuestions []Question

	sampleStage := bson.D{{"$sample", bson.D{{"size", 1}}}}

	pipeline := mongo.Pipeline{}
	pipeline = append(pipeline, sampleStage)

	opts := options.Aggregate().SetMaxTime(2 * time.Second)
	cursor, err := foreQuestionsCollection.Aggregate(ctx, pipeline, opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &foreQuestions)
	if err != nil {
		return nil, err
	}
	return foreQuestions, nil
}
