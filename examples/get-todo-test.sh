#!/bin/bash

output=$(go run main.go run examples/get-todo.yml --raw)
file_content=$(cat examples/get-todo.yml.out)

if [[ "$output" == "$file_content" ]]; then
    echo "PASS: test matches expected output"
else
    echo "FAIL: test does not match output"
    exit 1
fi
