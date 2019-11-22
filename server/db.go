package main

import (
	"time"
	"github.com/guregu/dynamo"
	// "github.com/aws/aws-sdk-go/aws"
	
)

type db interface{
	Get(key string) (string, error) 
	Put(key string, value string) error
	Delete(key string) error
}

type SignUpLog struct {
	UserID int
	Time time.Time

	Count int    `dynamo:"Message"`
	Msg   string `dynamo:"Message"`
}


func NewDynamoDB() {
	// db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})
}


