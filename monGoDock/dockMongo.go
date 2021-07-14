package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"gopkg.in/mgo.v2/bson"
)

const (
	// Name of the database.
	DBName = "sample_training"
	//URI    = "mongodb+srv://m001-student:m001-mongodb-basics@sandbox.9t09v.mongodb.net/sample_training?retryWrites=true&w=majority"
	//URI  = "mongodb://deadpoet:Achilles@localhost/july7db"

	//The below is the URI to connect Mongo server running in another container
	//The name "mongo-server" is the one declared in docker-compose file
	URI  = "mongodb://mongo-server:27017"
	COLL = "training_collection"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

type Note struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title     string             `bson:"title" json:"title"`
	Body      string             `bson:"body" json:"body"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

var ctx context.Context
var clientOptions *options.ClientOptions

var client *mongo.Client

var tmpl *template.Template

func init() {
	tmpl, _ = template.ParseFiles("index.gohtml")
}

func main() {
	http.HandleFunc("/connectmongo/", connectHandler)
	//http.HandleFunc("/writemongo/", writeHandler)
	http.HandleFunc("/", indexHandler)
	log.Fatalln(http.ListenAndServe(":8000", nil))

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	//ctx = context.Background()
	clientOptions = options.Client().ApplyURI(URI)

	// Connect to MongoDB
	var err error

	//you can also use the below method to connect
	/*ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()*/

	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("The error is below")
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("The ping error")
		log.Fatal(err)
	}

	fmt.Fprintln(w, "Database Successfully Connected")
	dbs, _ := client.ListDatabases(context.TODO(), bson.M{})
	fmt.Println(dbs)
	fmt.Fprintln(w, dbs)

	colls, _ := client.Database("users").ListCollectionNames(context.TODO(), bson.M{})
	fmt.Fprintln(w, colls)
}

/*func writeHandler(w http.ResponseWriter, r *http.Request) {
	note := Note{}
	connectHandler(w, r)
	db := client.Database(DBName)
	collections := db.Collection(COLL)
	fmt.Println(db.Name())
	fmt.Println(collections.Name())

	// An ID for MongoDB.
	note.ID = primitive.NewObjectID()
	note.Title = "Fifth Note"
	note.Body = "Some Fifth spam text"
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()
	result, err := collections.InsertOne(ctx, note)

	if err != nil {
		fmt.Println("The error")
		fmt.Println(err)
		return
	}

	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)
	title := collections.FindOne(ctx, bson.M{"title": "Fourth Note"})
	if err := title.Err(); err != nil {
		fmt.Println(err)
		return
	}
	n := Note{}
	err = title.Decode(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n.Body)
	fmt.Println(n.CreatedAt.Clock())
	fmt.Fprintln(w, n.Body)
	fmt.Fprintln(w, n.Title)
} */
