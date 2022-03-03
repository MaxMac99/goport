package events

import (
	"context"
	"errors"
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
	ObjectId    primitive.ObjectID `bson:"_id" json:"-"`
	DeviceToken string             `bson:"token"`
	EventType   string             `bson:"event"`
	Context     *string            `bson:"context" json:",omitempty"`
	Action      *string            `bson:"action" json:",omitempty"`
	Id          *string            `bson:"id" json:",omitempty"`
	Excluded    bool               `bson:"excluded"`
	Count       int                `bson:"count" json:"-"`
}

func RegisterNotificationListener(token string, eventType string, contextName *string, action *string, id *string) error {
	deletedListeners, err := removeExcludingListeners(token, eventType, contextName, action, id, false)
	if err != nil {
		return err
	}
	existingListener, err := getNextListener(token, eventType, contextName, action, id)
	if err != nil {
		return err
	}
	if existingListener != nil && !existingListener.Excluded {
		if deletedListeners != 0 {
			return nil
		}
		return &ListenerExistsError{
			Listener: *existingListener,
		}
	}
	listener := NotificationListener{
		ObjectId:    primitive.NewObjectID(),
		DeviceToken: token,
		EventType:   eventType,
		Context:     contextName,
		Action:      action,
		Id:          id,
		Excluded:    false,
		Count:       calculateCount(eventType, contextName, action, id),
	}
	_, err = collection.InsertOne(ctx, listener)
	return err
}

func UnregisterNotificationListener(token string, eventType string, contextName *string, action *string, id *string) error {
	deletedListeners, err := removeExcludingListeners(token, eventType, contextName, action, id, true)
	if err != nil {
		return err
	}
	existingListener, err := getNextListener(token, eventType, contextName, action, id)
	if err != nil {
		return err
	}
	if existingListener != nil && existingListener.Excluded {
		if deletedListeners != 0 {
			return nil
		}
		return &ListenerExistsError{
			Listener: *existingListener,
		}
	}
	if existingListener == nil {
		return &NoListenersToExcludeError{}
	}
	listener := NotificationListener{
		ObjectId:    primitive.NewObjectID(),
		DeviceToken: token,
		EventType:   eventType,
		Context:     contextName,
		Action:      action,
		Id:          id,
		Excluded:    true,
		Count:       calculateCount(eventType, contextName, action, id),
	}
	_, err = collection.InsertOne(ctx, listener)
	return err
}

func calculateCount(eventType string, contextName *string, action *string, id *string) int {
	count := 0
	if eventType != string(AllEventType) {
		count += 1
	}
	if contextName != nil {
		count += 1
	}
	if action != nil {
		count += 1
	}
	if id != nil {
		count += 1
	}
	return count
}

func removeExcludingListeners(token string, eventType string, contextName *string, action *string, id *string, exclude bool) (int64, error) {
	filter := bson.D{
		primitive.E{Key: "token", Value: token},
		primitive.E{Key: "$nor", Value: bson.A{
			primitive.E{Key: "event", Value: eventType},
			primitive.E{Key: "context", Value: contextName},
			primitive.E{Key: "action", Value: action},
			primitive.E{Key: "id", Value: id},
			primitive.E{Key: "excluded", Value: exclude},
		}},
	}
	if eventType != string(AllEventType) {
		filter = append(filter, primitive.E{Key: "event", Value: eventType})
	}
	if contextName != nil {
		filter = append(filter, primitive.E{Key: "context", Value: contextName})
	}
	if action != nil {
		filter = append(filter, primitive.E{Key: "action", Value: action})
	}
	if id != nil {
		filter = append(filter, primitive.E{Key: "id", Value: id})
	}

	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func getNextListener(token string, eventType string, contextName *string, action *string, id *string) (*NotificationListener, error) {
	filter := bson.D{
		primitive.E{Key: "token", Value: token},
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "event", Value: string(AllEventType)}},
				bson.D{primitive.E{Key: "event", Value: eventType}},
			}},
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "context", Value: contextName}},
				bson.D{primitive.E{Key: "context", Value: nil}},
			}},
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "action", Value: action}},
				bson.D{primitive.E{Key: "action", Value: nil}},
			}},
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "id", Value: id}},
				bson.D{primitive.E{Key: "id", Value: nil}},
			}},
	}
	sort := bson.D{
		primitive.E{Key: "count", Value: -1},
	}
	var listener NotificationListener
	if err := collection.FindOne(ctx, filter, options.FindOne().SetSort(sort)).Decode(&listener); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &listener, nil
}

func GetNotificationListenersWithToken(token string, eventType string, contextName *string, action *string, id *string) ([]NotificationListener, error) {
	filter := bson.D{
		primitive.E{Key: "token", Value: token},
	}
	if eventType != string(AllEventType) {
		filter = append(filter, primitive.E{Key: "event", Value: eventType})
	}
	if contextName != nil {
		filter = append(filter, primitive.E{Key: "context", Value: contextName})
	}
	if action != nil {
		filter = append(filter, primitive.E{Key: "action", Value: action})
	}
	if id != nil {
		filter = append(filter, primitive.E{Key: "id", Value: id})
	}
	sort := bson.D{
		primitive.E{Key: "count", Value: -1},
	}
	cur, err := collection.Find(ctx, filter, options.Find().SetSort(sort))
	if err != nil {
		return nil, err
	}
	listeners := make([]NotificationListener, 0)
	for cur.Next(ctx) {
		var listener NotificationListener
		if err = cur.Decode(&listener); err != nil {
			return nil, err
		}
		listeners = append(listeners, listener)
	}
	return listeners, nil
}

func GetDeviceTokensForEvent(contextName string, eventType string, action string, id string) ([]string, error) {
	matchStage := bson.D{primitive.E{Key: "$match", Value: bson.D{
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "event", Value: string(AllEventType)}},
				bson.D{primitive.E{Key: "event", Value: eventType}},
			}},
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "context", Value: contextName}},
				bson.D{primitive.E{Key: "context", Value: nil}},
			}},
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "action", Value: action}},
				bson.D{primitive.E{Key: "action", Value: nil}},
			}},
		primitive.E{Key: "$or",
			Value: bson.A{
				bson.D{primitive.E{Key: "id", Value: id}},
				bson.D{primitive.E{Key: "id", Value: nil}},
			}},
	}}}
	groupStage := bson.D{primitive.E{Key: "$group", Value: bson.D{
		primitive.E{Key: "_id", Value: "$token"},
		primitive.E{Key: "excluded", Value: bson.D{
			primitive.E{Key: "$accumulator", Value: bson.D{
				primitive.E{Key: "init", Value: "function(){return {count:0, excluded:false}}"},
				primitive.E{Key: "accumulate", Value: "function(state, count, excluded){if(count<state.count){return {count:state.count, excluded:state.excluded}}return {count:count, excluded:excluded}}"},
				primitive.E{Key: "accumulateArgs", Value: bson.A{
					"$count",
					"$excluded",
				}},
				primitive.E{Key: "merge", Value: "function(state1,state2){if(state1.count>state2.count){return {count:state1.count, excluded:state1.excluded}}return {count:state2.count, excluded:state2.excluded}}"},
				primitive.E{Key: "finalize", Value: "function(state){return state.excluded}"},
			}},
		}},
	}}}
	matchStage2 := bson.D{primitive.E{Key: "$match", Value: bson.D{
		primitive.E{Key: "excluded", Value: false},
	}}}
	groupStage2 := bson.D{primitive.E{Key: "$group", Value: bson.D{
		primitive.E{Key: "_id", Value: "$_id"},
	}}}

	cur, err := collection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, matchStage2, groupStage2})
	if err != nil {
		return nil, err
	}

	deviceTokens := make([]string, 0)
	for cur.Next(ctx) {
		var listener bson.M
		if err = cur.Decode(&listener); err != nil {
			return nil, err
		}
		deviceTokens = append(deviceTokens, listener["_id"].(string))
	}
	return deviceTokens, nil
}

func DeleteNotificationRegistrationForId(id string) error {
	shortId := id[:12]
	_, err := collection.DeleteMany(ctx, bson.D{
		primitive.E{Key: "id", Value: bson.D{
			primitive.E{Key: "$regex", Value: "^" + shortId},
		}},
	})
	return err
}
