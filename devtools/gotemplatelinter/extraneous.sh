#!/bin/bash

# Check if the file argument is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <json_file>"
    exit 1
fi

# Read the JSON file and extract keys
keys=$(jq 'paths | map(tostring) | join(".")' "$1")

# Check all found
all_found=true

# rg
for key in $keys; do
  if ! rg "$key" --type 'html' > /dev/null; then
    all_found=false
    echo "UNUSED $key"  # Would be better only to print this if -v is passed. But I don't know how to implement that. doglaugh
  fi
done

# Exit with error if there is unused key
if [ "$all_found" = false ]; then
    echo "Error: Some keys are unused"
    exit 1
fi

exit 0