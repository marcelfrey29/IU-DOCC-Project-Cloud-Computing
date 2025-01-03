echo "Updating Lambda Function Code..."

# Build and upload to S3
(cd ../backend && GOOS=linux GOARCH=arm64 go build -o bootstrap)
(cd ../backend && zip lambda-handler.zip bootstrap)
cp ../backend/lambda-handler.zip code/lambda-handler.zip
aws s3 sync ./code s3://$(aws cloudformation describe-stacks --stack-name TravelGuides --query 'Stacks[0].Outputs[?OutputKey==`DeploymentBucketName`].OutputValue' --output text)

echo "Upload of Lambda Function Code to S3 complete."

# Update Lambda with new Code from S3
aws lambda update-function-code \
    --function-name $(aws cloudformation describe-stacks --stack-name TravelGuides --query 'Stacks[0].Outputs[?OutputKey==`BackendLambdaName`].OutputValue' --output text) \
    --s3-bucket $(aws cloudformation describe-stacks --stack-name TravelGuides --query 'Stacks[0].Outputs[?OutputKey==`DeploymentBucketName`].OutputValue' --output text) \
    --s3-key "lambda-handler.zip" \
    --no-cli-pager

echo "Updated Lambda Function Code."
