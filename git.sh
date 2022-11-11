#!/bin/bash
echo "contracts"
git add . && \
git add -u && \
git commit -m $1 && \
git push origin main
