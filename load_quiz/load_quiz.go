package load_quiz

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

func ConnectMongo() {

	mongoUser := os.Getenv("MONGODB_USER")
	mongoPassword := os.Getenv("MONGODB_PASSWORD")

	// test db with private network access
	credential := options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	}
	mongoURI := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI).SetAuth(credential))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
}

func LoadQuestion() ([]Question, error) {

	mongoUser := os.Getenv("MONGODB_USER")
	mongoPassword := os.Getenv("MONGODB_PASSWORD")

	// test db with private network access
	credential := options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	mongoURI := os.Getenv("MONGODB_URI")
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI).SetAuth(credential))
	if err != nil {
		return nil, err
	}

	mongoDatabase := os.Getenv("MONGODB_DATABASE")
	mongoCollection := os.Getenv("MONGODB_COLLECTION")
	foreQuestionsCollection := client.Database(mongoDatabase).Collection(mongoCollection)
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
