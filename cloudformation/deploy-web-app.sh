echo "Deploying Web App to S3..."

aws s3 sync ../web-app/dist s3://$(aws cloudformation describe-stacks --stack-name TravelGuides --query 'Stacks[0].Outputs[?OutputKey==`WebAppBucketName`].OutputValue' --output text)

echo "Deployment of Web App complete."