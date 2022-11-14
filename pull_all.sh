#!/bin/bash
cd ../regulations_writable_service/
git pull origin main
cd ../regulations_read_only_service/
git pull origin main
cd ../regulations_master_service/
git pull origin main
cd ../regulations_mobile_service/
git pull origin main
cd ../regulations_http_server/
git pull origin main
cd ../regulations_parser/
git pull origin main