package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	URI = "mongodb+srv://m001-student:m001-mongodb-basics@sandbox.9t09v.mongodb.net/sample_training?retryWrites=true&w=majority"
)

var clientOptions *options.ClientOptions
var client *mongo.Client
var dbConnerr error
var dbName string
var collName string

var tmpl *template.Template

func init() {
	tmpl, _ = template.ParseGlob("*.gohtml")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/connectMongo", connectHandler)
	http.HandleFunc("/showCollections", showDBColl)
	http.HandleFunc("/findOne", findOne)
	log.Fatalln(http.ListenAndServe(":1212", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func findOne(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		collName = r.FormValue("collName")
		coll := client.Database(dbName).Collection(collName)
		//fmt.Println("Collection type:", reflect.ValueOf(coll))
		//var movie fields.Movie

		//The below commented are to fetch results using Find
		/*cursor, err := coll.Find(context.TODO(), bson.M{})
		if err != nil {
			fmt.Println(err)
		}

		var movie bson.M
		defer cursor.Close(context.TODO())
		for cursor.Next(context.TODO()) {
			if err = cursor.Decode(&movie); err != nil {
				log.Fatal(err)
			}
			fmt.Println(movie)
		}*/

		var movie bson.M
		if err := coll.FindOne(context.TODO(), bson.M{}).Decode(&movie); err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, movie)
	}
}

func showDBColl(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		dbName = r.FormValue("dbname")
		colls, err := client.Database(dbName).ListCollectionNames(context.TODO(), bson.M{})
		if err != nil {
			fmt.Fprintln(w, "Error displaying Collections", err)
		}
		//fmt.Fprintln(w, "The collections present are - ", colls)
		tmpl.ExecuteTemplate(w, "collections.gohtml", colls)
	}
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	clientOptions = options.Client().ApplyURI(URI)
	clientOptions.SetConnectTimeout(time.Second * 10)
	client, dbConnerr = mongo.Connect(context.TODO(), clientOptions)
	if dbConnerr != nil {
		fmt.Println("Connection to DB Failed")
		//fmt.Fprintln(w, "Connect to DB Failed", dbConnerr)
	}

	// Check the connection
	err := client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("The ping error")
		log.Fatal(err)
	}

	//fmt.Fprintln(w, "Database Successfully Connected")
	dbs, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		fmt.Fprintln(w, "Error while listing databases", err)
	}
	tmpl.ExecuteTemplate(w, "connect.gohtml", dbs)
}
