echo "Deploying Web App to S3..."

(cd ../web-app && npm install --frozen-lockfile && npm run build:cloud)
aws s3 sync ../web-app/dist s3://$(aws cloudformation describe-stacks --stack-name TravelGuides --query 'Stacks[0].Outputs[?OutputKey==`WebAppBucketName`].OutputValue' --output text)

echo "Deployment of Web App complete."