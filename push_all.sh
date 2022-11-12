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

cd ../regulations_writable_service/
function_push "writable"
cd ../regulations_read_only_service/
function_push "read-only"
cd ../regulations_supreme_service/
function_push "supreme"
cd ../regulations_mobile_service/
function_push "mobile"
cd ../regulations_http_server/
function_push "http"
cd ../regulations_parser/
function_push "parser"