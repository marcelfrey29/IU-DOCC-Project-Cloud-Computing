#!/bin/zsh

echo "Creating CloudFormation Stack..."

aws cloudformation create-stack --stack-name TravelGuides --template-body file://template.yaml --capabilities CAPABILITY_IAM

echo "Created CloudFormation Stack."
