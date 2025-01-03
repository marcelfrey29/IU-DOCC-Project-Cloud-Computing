echo "Uploading Lambda Code to S3..."

(cd ../backend && GOOS=linux GOARCH=amd64 go build -o bootstrap)
zip ./code/lambda-handler.zip ../backend/bootstrap
aws s3 sync ./code s3://$(aws cloudformation describe-stacks --stack-name TravelGuides --query 'Stacks[0].Outputs[?OutputKey==`DeploymentBucketName`].OutputValue' --output text)

echo "Upload of Lambda Function Code to S3 complete."
