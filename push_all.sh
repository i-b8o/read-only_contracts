#!/bin/bash

function_push() {
    if [[ `git status --porcelain` ]]; then
        git status
        git diff
        echo $1
        echo 'commit messsage:'
        read message
        git add . && \
        git add -u && \
        git commit -m message && \
        git push origin main
    fi
}

function_push "contracts"
cd ../read-only_writer_service/
function_push "writer"
cd ../read-only_reader_service/
function_push "reader"
cd ../read-only_master_service/
function_push "master"
cd ../read-only_mobile_service/
function_push "mobile"
cd ../read-only_http_server/
function_push "http"
cd ../read-only_parser/
function_push "parser"
# cd ../read-only_searcher_server/
# function_push "searcher"