package main

import (
	"context"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	"github.com/gofiber/fiber/v2/log"
)

var tableName = os.Getenv("DYNAMODB_TABLE_TRAVEL_GUIDES")
var useDynamoDbLocal, _ = strconv.ParseBool(os.Getenv("USE_DYNAMODB_LOCAL"))
var dynamoDbLocalUrl = os.Getenv("DYNAMODB_LOCAL_URL")

// Get the AWS Config to use, required to determine if DynamoDB Local should be used.
func getAwsConfig() aws.Config {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	if useDynamoDbLocal {
		log.Info("DynamoDB Local should be used, setting local DynamoDB URL.", useDynamoDbLocal, dynamoDbLocalUrl)
		cfg.BaseEndpoint = &dynamoDbLocalUrl
	}
	return cfg
}

var ddbClient = dynamodb.NewFromConfig(getAwsConfig())

// Create a new Travel Guide in the Database
func CreateTravelGuideInDDB(travelGuide *TravelGuideItem) (*TravelGuideItem, error) {
	log.Info("Creating new Travel Guide in DynamoDB.", travelGuide.Id, travelGuide.TravelGuide)

	itemToStore, _ := attributevalue.MarshalMap(travelGuide)

	_, err := ddbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                itemToStore,
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	})
	if err != nil {
		log.Error("Error while creating Travel Guide in DyanmoDB.", err.Error())
		return nil, err
	}

	log.Info("Created Travel Guide in DyanmoDB.", travelGuide.Id)
	return travelGuide, nil
}
