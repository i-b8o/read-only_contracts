#!/bin/bash
function_get() {
    echo $2
    go get -u github.com/i-b8o/regulations_contracts@$1
}
read -p "Enter a tag: " tag
cd ../regulations_writable_service/app
function_get $tag "writable"
pwd
cd ../../regulations_read_only_service/app
function_get $tag "read only"
pwd
cd ../../regulations_supreme_service/app
function_get $tag "supreme"
pwd
cd ../../regulations_mobile_service/app
function_get $tag "mobile"
pwd
cd ../../regulations_parser/app
function_get $tag "parser"
pwd