#!/bin/sh

SESSION_NAME="bmdc"

tmux has-session -t ${SESSION_NAME}

if [ $? != 0 ]
then
  tmux new-session -s ${SESSION_NAME} -n server -d
  tmux send-keys -t ${SESSION_NAME} 'gin run bmdc-server.go' C-m
  echo "Started BMDC server"
fi

