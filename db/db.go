package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var db *dynamo.DB

//Init Setup the databse stuff
func Init(awsRegion string, DatabaseEndpoint string) {
	// initiate the Database Object
	db = dynamo.New(session.New(), &aws.Config{Region: aws.String(awsRegion), Endpoint: aws.String(DatabaseEndpoint)})
}

// GetDB does stuff
func GetDB() *dynamo.DB {
	return db
}
