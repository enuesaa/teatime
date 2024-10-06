package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func poc() {
	ctx := context.Background()

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("app")

	schema := bson.M{
		"bsonType": "object",
		"required": []string{"name", "value"},
		"properties": bson.M{
			"name": bson.M{
				"bsonType":    "string",
				"description": "must be a string and is required",
			},
			"value": bson.M{
				"bsonType":    "string",
				"description": "must be a string and is required",
			},
		},
	}
	validator := bson.M{
		"$jsonSchema": schema,
	}
	err = db.CreateCollection(ctx, "notes",
		options.CreateCollection().SetValidator(validator),
		options.CreateCollection().SetValidationLevel("strict"),
		options.CreateCollection().SetValidationAction("error"),
	)
	if err != nil {
		panic(err)
	}

	res, err := db.Collection("notes").InsertOne(ctx, bson.M{
		"name": "aaa",
		"value": "99",
	})
	if err != nil {
		panic(err)
	}
	id := res.InsertedID
	fmt.Printf("id: %s\n", id)

	// var result struct {
	// 	Value string
	// }
	// filter := bson.D{{Key: "name", Value: "aaa"}}
	// if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("value: %s\n", result.Value)
}
