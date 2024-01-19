#!/bin/bash

# Check if Go is installed
if ! type go >/dev/null 2>&1; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Set GOPATH if not set
if [ -z "$GOPATH" ]; then
    export GOPATH=$HOME/go
fi

# Create necessary directories if they don't exist
mkdir -p $GOPATH/src/sai
mkdir -p $GOPATH/bin


# Build the sai project
go build -o sai main.go

# Move the sai binary to /usr/local/bin
sudo mv sai /usr/local/bin/

echo "sai has been installed successfully. Run 'sai \"question to command\"' to use the tool."
