#bash

echo "Setting up environment variables for production profile..."
sh /"$APP"/bin/setenv.sh

APP_ENV="${1:-prod}"

# Print the value of APP_ENV for debugging purposes
echo "APP_ENV is set to: $APP_ENV"

# Set config folder path
CONFIG_FOLDER="/$APP/conf"
echo "CONFIG_FOLDER is set to: $CONFIG_FOLDER"

chmod +x /"$APP"/kubesync

# Run the Go application and inject APP_ENV and CONFIG_FOLDER
APP_ENV="$APP_ENV" CONFIG_FOLDER="$CONFIG_FOLDER" /"$APP"/kubesync