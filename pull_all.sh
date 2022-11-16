#!/bin/bash
git pull origin main
cd ../read-only_writer_service/
git pull origin main
cd ../read-only_master_service/
git pull origin main
cd ../read-only_reader_service/
git pull origin main
# cd ../read-only_searcher_service/
# git pull origin main
cd ../read-only_mobile_service/
git pull origin main
cd ../read-only_http_server/
git pull origin main
cd ../read-only_parser/
git pull origin main