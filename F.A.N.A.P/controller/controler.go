package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"F.A.N.A.P/algorithm"
	"F.A.N.A.P/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "fanapdb"
const colName = "rectangleList"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connected successfully.")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready.")
}

//Mongodb helper
func insertOne(rec models.Rectangle) {
	inserted, err := collection.InsertOne(context.Background(), rec)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("A rectangle is inserted: ", inserted.InsertedID)
}

func getAllRectangles() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var rectangles []primitive.M
	for cur.Next(context.Background()) {
		var rectangle bson.M
		err := cur.Decode(&rectangle)

		if err != nil {
			log.Fatal(err)
		}
		rectangles = append(rectangles, rectangle)
	}

	defer cur.Close(context.Background())

	return rectangles
}

func GetAllRectangles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all valid Rectangles")
	w.Header().Set("Content-Type", "application/json")
	validRectangles := getAllRectangles()
	json.NewEncoder(w).Encode(validRectangles)

}

func Insert(w http.ResponseWriter, r *http.Request) {
	presentTime := time.Now().Format("2006-02-01 15:04:05")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send valid data!")
	}

	var inputtype models.Inputmodel
	json.NewDecoder(r.Body).Decode(&inputtype)

	var mainrec models.Rectangle = inputtype.Main
	var m1 = models.Point{mainrec.X, mainrec.Y}
	var m2 = models.Point{mainrec.X + mainrec.Width, mainrec.Y + mainrec.Height}

	for _, element := range inputtype.Input {
		var e1 = models.Point{element.X, element.Y}
		var e2 = models.Point{element.X + element.Width, element.Y + element.Height}

		if algorithm.Overlap(m1, m2, e1, e2) {
			element.Time = presentTime
			insertOne(element)
		}

	}
}
