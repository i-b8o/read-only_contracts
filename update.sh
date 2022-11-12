#!/bin/bash
function_get() {
    echo $2
    go get -u github.com/i-b8o/regulations_contracts@$1
}
read -p "Enter a tag: " tag
cd ../regulations_writable_service/app
function_get $tag "writable"
cd ../regulations_read_only_service/app
function_get $tag "read only"
cd ../regulations_supreme_service/app
function_get $tag "supreme"
cd ../regulations_mobile_service/app
function_get $tag "mobilt"
cd ../regulations_parser/app
function_get $tag "parser"