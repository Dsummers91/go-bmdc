#!/bin/bash

cd /home/deon/go-bmdc
git add .
git commit --no-gpg-sign -m "daily update"
git push origin master

