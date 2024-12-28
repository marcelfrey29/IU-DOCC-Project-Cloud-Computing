package main

import (
	"context"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

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

// Get all Travel Guides from the Database.
func GetTravelGuidesFromDDB() ([]TravelGuideItem, error) {
	log.Info("Scanning for all Travel Guides.")

	var attributeValues map[string]types.AttributeValue = make(map[string]types.AttributeValue)
	attributeValues[":TG"] = &types.AttributeValueMemberS{
		Value: "TG",
	}

	response, err := ddbClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    aws.String("hashId = :TG"),
		ExpressionAttributeValues: attributeValues,
	})
	if err != nil {
		log.Error("Error while scanning for Travel Guides in DynamoDB.", err.Error())
		return nil, err
	}

	var travelGuides []TravelGuideItem
	for _, travelGuide := range response.Items {
		tgi := new(TravelGuideItem)
		attributevalue.UnmarshalMap(travelGuide, tgi)
		travelGuides = append(travelGuides, *tgi)
	}

	return travelGuides, nil
}

// Get all Travel Guides from the Database.
func GetTravelGuideFromDDB(id string) (TravelGuideItem, error) {
	log.Info("Get Travel Guide by ID.", id)

	var key map[string]types.AttributeValue = make(map[string]types.AttributeValue)
	key["hashId"] = &types.AttributeValueMemberS{
		Value: "TG",
	}
	key["rangeId"] = &types.AttributeValueMemberS{
		Value: id,
	}

	response, err := ddbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})
	if err != nil {
		log.Error("Error while getting Travel Guides from DynamoDB.", err.Error())
		return TravelGuideItem{}, err
	}

	if response.Item == nil {
		log.Warn("No Item in DynamoDB.")
		return TravelGuideItem{}, &NotFoundError{message: "The item doesn't exist."}
	}

	tgi := new(TravelGuideItem)
	attributevalue.UnmarshalMap(response.Item, tgi)

	return *tgi, nil
}

// Create a new Travel Guide in the Database
func CreateTravelGuideInDDB(travelGuide *TravelGuideItem) (*TravelGuideItem, error) {
	log.Info("Creating new Travel Guide in DynamoDB.", travelGuide.HashId, travelGuide.RangeId, travelGuide.TravelGuide)

	itemToStore, _ := attributevalue.MarshalMap(travelGuide)

	_, err := ddbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                itemToStore,
		ConditionExpression: aws.String("attribute_not_exists(hashId) and attribute_not_exists(rangeId)"),
	})
	if err != nil {
		log.Error("Error while creating Travel Guide in DyanmoDB.", err.Error())
		return nil, err
	}

	log.Info("Created Travel Guide in DyanmoDB.", travelGuide.HashId, travelGuide.RangeId)
	return travelGuide, nil
}

// Update a Travel Guide in the Database
func UpdateTravelGuideInDDB(travelGuide *TravelGuideItem) (*TravelGuideItem, error) {
	log.Info("Updating Travel Guide in DynamoDB.", travelGuide.HashId, travelGuide.RangeId, travelGuide.TravelGuide)

	itemToStore, _ := attributevalue.MarshalMap(travelGuide)

	_, err := ddbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                itemToStore,
		ConditionExpression: aws.String("attribute_exists(hashId) AND attribute_exists(rangeId)"),
	})
	if err != nil {
		log.Error("Error while updating Travel Guide in DyanmoDB.", err.Error())
		return nil, err
	}

	log.Info("Updated Travel Guide in DyanmoDB.", travelGuide.HashId, travelGuide.RangeId)
	return travelGuide, nil
}

// Delete a Travel Guides from the Database.
func DeleteGuideFromDDB(id string) error {
	log.Info("Delete Travel Guide by ID.", id)

	var key map[string]types.AttributeValue = make(map[string]types.AttributeValue)
	key["hashId"] = &types.AttributeValueMemberS{
		Value: "TG",
	}
	key["rangeId"] = &types.AttributeValueMemberS{
		Value: id,
	}

	_, err := ddbClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})
	if err != nil {
		log.Error("Error while deleting Travel Guide from DynamoDB.", err.Error())
		return err
	}

	return nil
}
