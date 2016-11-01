package counter

import (
	"./../../../model"
	"./../../mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var database = "mGoMuxRest"
var collection = "counters"

// GetNextCounter return new ID for the Resource passed as Argument
func GetNextCounter(resource string) (int, error) {
	session := mongo.GetMongoSession()
	defer session.Close()
	counterCollection := session.DB(database).C(collection)
	query := bson.M{"resource": resource}
	update := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"c": 1}},
		ReturnNew: true,
		Upsert:    true,
	}
	counter := model.Counter{}
	_, err := counterCollection.Find(query).Apply(update, &counter)
	if err != nil {
		return 0, err
	}
	return counter.C, err
}
