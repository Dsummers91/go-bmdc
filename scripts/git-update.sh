#!/bin/bash

cd ~/go-bmdc
git add .
git commit --no-gpg-sign -m "daily update"
git push origin master
