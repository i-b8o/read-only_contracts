#!/bin/bash
function_get() {
    echo $2
    go get -u github.com/i-b8o/regulations_contracts@$1
}
read -p "Enter a tag: " tag
cd ../regulations_writable_service/app
function_get $tag "writable"
cd ../../regulations_read_only_service/app
function_get $tag "read only"
cd ../../regulations_master_service/app
function_get $tag "master"
cd ../../regulations_mobile_service/app
function_get $tag "mobile"
cd ../../regulations_http_server/app
function_get $tag "http"
cd ../../regulations_parser/app
function_get $tag "parser"
