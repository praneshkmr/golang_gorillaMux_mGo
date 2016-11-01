package counter

import (
	"./../../../model"
	"./../../mongo"
	"gopkg.in/mgo.v2/bson"
)

var database = "mGoMuxRest"
var collection = "counters"

// GetNextCounter return new ID for the Resource passed as Argument
func GetNextCounter(resource string) (int, error) {
	session := mongo.GetMongoSession()
	defer session.Close()
	counterCollection := session.DB(database).C(collection)
	selector := bson.M{"resource": resource}
	update := bson.M{"$inc": bson.M{"c": 1}}
	_, err := counterCollection.Upsert(selector, update)
	if err != nil {
		return 0, err
	}
	counter := model.Counter{}
	err = counterCollection.Find(bson.M{"resource": resource}).One(&counter)
	if err != nil {
		return 0, err
	}
	return counter.C, err
}
