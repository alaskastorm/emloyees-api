package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sort"
)

var (
	EmployeeCollection   *mongo.Collection
	DepartmentCollection *mongo.Collection
)

// ConnectToMongo ...
func ConnectToMongo() {

	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb://ala:secret@127.0.0.1:27027/rest_api")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	collections, err := client.Database("rest_api").ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	sort.Strings(collections)

	employeesCollectionIndex := sort.SearchStrings(collections, "employees")
	if collections[employeesCollectionIndex] != "employees" {
		if err := client.Database("rest_api").CreateCollection(context.TODO(), "employees"); err != nil {
			log.Fatal(err)
		}
		EmployeeCollection = client.Database("rest_api").Collection("employees")
	}else {
		EmployeeCollection = client.Database("rest_api").Collection("employees")
	}

	departmentsCollectionIndex := sort.SearchStrings(collections, "departments")
	if collections[departmentsCollectionIndex] != "departments" {
		if err := client.Database("rest_api").CreateCollection(context.TODO(), "departments"); err != nil {
			log.Fatal(err)
		}
		DepartmentCollection = client.Database("rest_api").Collection("departments")
	}else {
		DepartmentCollection = client.Database("rest_api").Collection("departments")
	}
}
