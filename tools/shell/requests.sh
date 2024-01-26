#!/bin/bash

# Check if at least one argument was provided
if [ "$#" -lt 1 ]; then
    echo "Usage: $0 {block|transactions|subscribe} [address]"
    exit 1
fi

# The API endpoint
API_ENDPOINT="http://127.0.0.1:8080"

# Function to get the latest block
block() {
    curl --location "$API_ENDPOINT/block" | jq
}

# Function to get transactions for a specific address
transactions() {
    if [ "$#" -ne 2 ]; then
        echo "Usage: $0 transactions [address]"
        exit 1
    fi
    local address=$2
    curl --location "$API_ENDPOINT/transactions/$address" | jq
}

# Function to subscribe to an address
subscribe() {
    if [ "$#" -ne 2 ]; then
        echo "Usage: $0 subscribe [address]"
        exit 1
    fi
    local address=$2
    curl -X PUT --location "$API_ENDPOINT/subscribe" \
    --header 'Content-Type: application/json' \
    --data '{
        "address": "'$address'"
    }'
}

# Check the first argument and call the corresponding function
case "$1" in
    block)
        block
        ;;
    transactions)
        transactions "$@"
        ;;
    subscribe)
        subscribe "$@"
        ;;
    *)
        echo "Invalid argument: $1"
        echo "Usage: $0 {block|transactions|subscribe} [address]"
        exit 1
        ;;
esac
