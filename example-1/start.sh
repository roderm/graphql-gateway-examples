#!/bin/bash

function cleanup {
    kill "$ACCOUNTS_PID"
}
trap cleanup EXIT

go build -o /tmp/srv-accounts ./accounts
go build -o /tmp/srv-gateway ./gateway

/tmp/srv-accounts &
ACCOUNTS_PID=$!

sleep 1

/tmp/srv-gateway
