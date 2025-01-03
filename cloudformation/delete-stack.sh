#!/bin/zsh

echo "Deleting CloudFormation Stack..."

# Empty S3 Buckets, otherwise delete fails
aws s3 rm --recursive s3://$(aws cloudformation describe-stacks --stack-name TravelGuides --query 'Stacks[0].Outputs[?OutputKey==`WebAppBucketName`].OutputValue' --output text)/
aws s3 rm --recursive s3://$(aws cloudformation describe-stacks --stack-name TravelGuides --query 'Stacks[0].Outputs[?OutputKey==`DeploymentBucketName`].OutputValue' --output text)/

echo "Deleted all files from S3, ready to delete."

aws cloudformation delete-stack --stack-name TravelGuides

echo "Deleted CloudFormation Stack."
