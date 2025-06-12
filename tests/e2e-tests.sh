#!/bin/bash

# Standard GET
output_1=$(./bin/sendex run tests/get-todo.yml --raw)
desired_output_1=$(cat tests/get-todo.output)

if [ "$output_1" = "$desired_output_1" ]; then
    echo "PASS: get-todo matches expected output"
else
    echo "FAIL: get-todo does not match output"
    exit 1
fi

# Standard POST
output_2=$(./bin/sendex run tests/post-todo.yml --raw)
desired_output_2=$(cat tests/post-todo.output)

if [ "$output_2" = "$desired_output_2" ]; then
    echo "PASS: post-todo matches expected output"
else
    echo "FAIL: post-todo does not match output"
    exit 1
fi

# GET but only headers
output_3=$(./bin/sendex run tests/get-todo.yml -e --raw)
desired_output_3=$(cat tests/get-todo-headers.output)

if [ "$output_3" = "$desired_output_3" ]; then
    echo "PASS: get-todo-headers matches expected output"
else
    echo "FAIL: get-todo-headers does not match output"
    exit 1
fi

# GET but only body
output_4=$(./bin/sendex run tests/get-todo.yml -b --raw)
desired_output_4=$(cat tests/get-todo-body.output)

if [ "$output_4" = "$desired_output_4" ]; then
    echo "PASS: get-todo-body matches expected output"
else
    echo "FAIL: get-todo-body does not match output"
    exit 1
fi

# GET but only status code
output_5=$(./bin/sendex run tests/get-todo.yml -s --raw)
desired_output_5=$(cat tests/get-todo-status.output)

if [ "$output_5" = "$desired_output_5" ]; then
    echo "PASS: get-todo-status matches expected output"
else
    echo "FAIL: get-todo-status does not match output"
    exit 1
fi

# GET with formatting
output_6=$(./bin/sendex run tests/get-todo.yml)
desired_output_6=$(cat tests/get-todo-formatted.output)

if [ "$output_6" = "$desired_output_6" ]; then
    echo "PASS: get-todo-formatted matches expected output"
else
    echo "FAIL: get-todo-formatted does not match output"
    exit 1
fi

# GET with env vars
output_7=$(./bin/sendex run tests/get-todo-env.yml --raw)
desired_output_7=$(cat tests/get-todo-2.output)

if [ "$output_7" = "$desired_output_7" ]; then
    echo "PASS: get-todo-env matches expected output"
else
    echo "FAIL: get-todo-env does not match output"
    exit 1
fi
