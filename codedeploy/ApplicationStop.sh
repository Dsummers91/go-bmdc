#!/bin/bash
if [[ $(docker ps | grep mydocker) ]]; then
  docker stop bmdc
fi

if [[ $(docker ps -a | grep mydocker) ]]; then
  docker rm --force bmdc
fi
