package events

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()
var collection *mongo.Collection

func setupDatabase() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://db:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("goport").Collection("notifications")
}

type NotificationListener struct {
	Id          primitive.ObjectID `bson:"_id"`
	DeviceToken string             `bson:"token"`
	EventType   string             `bson:"event"`
	Contexts    []string           `bson:"contexts"`
	Actions     []string           `bson:"actions"`
	Ids         []string           `bson:"ids"`
}

func RegisterNotificationListener(listener NotificationListener) error {
	listener.Id = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), listener)
	return err
}

func UnregisterNotificationListener(listener NotificationListener) error {
	listener.Id = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), listener)
	return err
}

func GetDeviceTokensForEvent(contextName string, eventType string, action string, id string) ([]string, error) {
	filter := bson.D{
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "event", Value: AllEventType}},
				bson.D{primitive.E{Key: "event", Value: eventType}},
			}},
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "contexts", Value: "all"}},
				bson.D{primitive.E{Key: "contexts", Value: contextName}},
				bson.D{primitive.E{Key: "contexts", Value: nil}},
				bson.D{primitive.E{Key: "contexts", Value: bson.D{
					primitive.E{Key: "$exists", Value: true},
					primitive.E{Key: "$size", Value: 0},
				}}},
			}},
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "actions", Value: "all"}},
				bson.D{primitive.E{Key: "actions", Value: action}},
				bson.D{primitive.E{Key: "actions", Value: nil}},
				bson.D{primitive.E{Key: "actions", Value: bson.D{
					primitive.E{Key: "$exists", Value: true},
					primitive.E{Key: "$size", Value: 0},
				}}},
			}},
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "ids", Value: "all"}},
				bson.D{primitive.E{Key: "ids", Value: id}},
				bson.D{primitive.E{Key: "ids", Value: nil}},
				bson.D{primitive.E{Key: "ids", Value: bson.D{
					primitive.E{Key: "$exists", Value: true},
					primitive.E{Key: "$size", Value: 0},
				}}},
			}},
	}
	projection := bson.D{
		primitive.E{Key: "token", Value: 1},
		primitive.E{Key: "_id", Value: 0},
	}
	cur, err := collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}

	deviceTokens := make([]string, 0)
	for cur.Next(ctx) {
		var listener NotificationListener
		if err = cur.Decode(&listener); err != nil {
			return nil, err
		}
		deviceTokens = append(deviceTokens, listener.DeviceToken)
	}
	return deviceTokens, nil
}

func DeleteNotificationRegistrationForId(id string) error {
	shortId := id[:12]
	cur, err := collection.Find(ctx, bson.D{
		primitive.E{Key: "ids", Value: bson.D{
			primitive.E{Key: "$regex", Value: "^" + shortId},
		}},
	})
	if err != nil {
		return err
	}
	for cur.Next(ctx) {
		var listener NotificationListener
		if err = cur.Decode(&listener); err != nil {
			return err
		}
		if len(listener.Ids) == 1 {
			fmt.Println("Delete full object")
			_, err := collection.DeleteOne(ctx, bson.D{
				primitive.E{Key: "_id", Value: listener.Id},
			})
			if err != nil {
				return err
			}
		} else {
			fmt.Println("Delete id")
			collection.UpdateByID(ctx, listener.Id, bson.D{
				primitive.E{Key: "$pull", Value: bson.D{
					primitive.E{Key: "$regex", Value: "^" + shortId},
				}},
			})
		}
	}
	return nil
}
