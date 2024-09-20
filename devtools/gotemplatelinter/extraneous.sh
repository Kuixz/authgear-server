#!/bin/bash

# Check if the file argument is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <json_file>"
    exit 1
fi

# Read the JSON file and extract keys
keys=$(jq 'paths | map(tostring) | join(".")' "$1")

# rg
for key in $keys; do
  if rg "$key" --type 'html' > /dev/null; then
    echo -n "       "
  else
    echo -n "UNUSED "
  fi
  echo "$key"
done