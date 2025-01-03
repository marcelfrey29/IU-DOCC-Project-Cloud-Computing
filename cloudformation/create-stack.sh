#!/bin/zsh

echo "Creating CloudFormation Stack..."

aws cloudformation create-stack --stack-name TravelGuides --template-body file://template.yaml

echo "Created CloudFormation Stack."
