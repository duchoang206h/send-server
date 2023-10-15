package config

import (
	"fmt"
)

func GetMongoURI () string {
	dbPort :=Config("DB_PORT")
	dbUser := Config("DB_USER")
	dbPass := Config("DB_PASSWORD")
	dbHost := Config("DB_HOST")
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/", dbUser, dbPass, dbHost, dbPort)
	fmt.Println("mongoURI:::", mongoURI)
	return mongoURI
}
func GetMongoDBName () string {
	dbName := Config("DB_NAME")
	return dbName
}