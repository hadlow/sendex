#!/bin/bash

output_1=$(sendex run tests/get-todo.yml --raw)
desired_output_1=$(cat tests/get-todo.output)

if [[ "$output_1" == "$desired_output_1" ]]; then
    echo "PASS: get-todo matches expected output"
else
    echo "FAIL: get-todo does not match output"
    exit 1
fi

output_2=$(sendex run tests/post-todo.yml --raw)
desired_output_2=$(cat tests/post-todo.output)

if [[ "$output_2" == "$desired_output_2" ]]; then
    echo "PASS: post-todo matches expected output"
else
    echo "FAIL: post-todo does not match output"
    exit 1
fi

output_3=$(sendex run tests/404-todo.yml --raw)
desired_output_3=$(cat tests/404-todo.output)

if [[ "$output_3" == "$desired_output_3" ]]; then
    echo "PASS: 404-todo matches expected output"
else
    echo "FAIL: 404-todo does not match output"
    exit 1
fi
