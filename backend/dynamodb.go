package main

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gofiber/fiber/v2/log"
)

var tableName=os.Getenv("DYNAMODB_TABLE_TRAVEL_GUIDES")
var useDynamoDbLocal,_=strconv.ParseBool(os.Getenv("USE_DYNAMODB_LOCAL"))
var dynamoDbLocalUrl= os.Getenv("DYNAMODB_LOCAL_URL")

// Get the AWS Config to use, required to determine if DynamoDB Local should be used. 
func getAwsConfig() aws.Config {
	cfg,_:=config.LoadDefaultConfig(context.TODO())
	if useDynamoDbLocal {
		log.Info("DynamoDB Local should be used, setting local DynamoDB URL.",useDynamoDbLocal,dynamoDbLocalUrl)
		cfg.BaseEndpoint=&dynamoDbLocalUrl
	} 
	return cfg
}

var ddbClient = dynamodb.NewFromConfig(getAwsConfig())
