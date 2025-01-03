#!/bin/zsh

echo "Updating CloudFormation Stack..."

aws cloudformation update-stack --stack-name TravelGuides --template-body file://template.yaml

echo "Updating CloudFormation Stack."
