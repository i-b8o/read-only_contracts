#!/bin/bash
function_get() {
    echo $2
    go get -u github.com/i-b8o/read-only_contracts@$1
}
read -p "Enter a tag: " tag
cd ../read-only_writer_service
function_get $tag "writer"
cd ../read-only_reader_service
function_get $tag "reader"
cd ../read-only_master_service
function_get $tag "master"
# cd ../read-only_searcher_service
# function_get $tag "searcher"
cd ../read-only_mobile_service
function_get $tag "mobile"
cd ../read-only_http_server
function_get $tag "http"
cd ../read-only_parser
function_get $tag "parser"
