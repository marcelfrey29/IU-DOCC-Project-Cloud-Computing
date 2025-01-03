#!/bin/zsh

echo "Updating CloudFormation Stack..."

aws cloudformation update-stack --stack-name TravelGuides --template-body file://template.yaml --capabilities CAPABILITY_IAM

echo "Updating CloudFormation Stack."
