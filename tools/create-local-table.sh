#!/bin/zsh

aws dynamodb create-table \
    --table-name TravelGuides \
    --attribute-definitions \
        AttributeName=hashId,AttributeType=S \
        AttributeName=rangeId,AttributeType=S \
    --key-schema \
        AttributeName=hashId,KeyType=HASH \
        AttributeName=rangeId,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD \
    --endpoint-url http://localhost:8000