#bash

echo "Setting up environment variables for production profile..."
sh /"$APP"/bin/setenv.sh

APP_ENV="${1:-prod}"

# Print the value of APP_ENV for debugging purposes
echo "APP_ENV is set to: $APP_ENV"

chmod +x /"$APP"/kubesync

# Run the Go application and inject APP_ENV
APP_ENV="$APP_ENV" /"$APP"/kubesync