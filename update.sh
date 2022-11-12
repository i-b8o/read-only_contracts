#!/bin/bash
echo 'tag:'
read tag
cd ../regulations_writable_service/app
go get -u github.com/i-b8o/regulations_contracts@tag
cd ../regulations_read_only_service/app
go get -u github.com/i-b8o/regulations_contracts@tag
cd ../regulations_supreme_service/app
go get -u github.com/i-b8o/regulations_contracts@tag
cd ../regulations_mobile_service/app
go get -u github.com/i-b8o/regulations_contracts@tag
cd ../regulations_parser/app
go get -u github.com/i-b8o/regulations_contracts@tag