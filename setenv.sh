#!/bin/sh

export PRINTERMANAGER_SQLDSN="host=$POSTGRES_HOST port=$POSTGRES_PORT user=$POSTGRES_USER password=$POSTGRES_PASS dbname=$POSTGRES_DATABASE sslmode=disable"

exec $1
