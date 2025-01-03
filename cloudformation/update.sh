#!/bin/zsh

echo "Updating..."

./update-stack.sh
./deploy-web-app.sh
./deploy-backend.sh

echo "Update Complete."