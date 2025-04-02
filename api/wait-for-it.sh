#!/bin/bash
# wait-for-it.sh

set -e

TIMEOUT=30
HOST_PORT=$1
shift

HOST=$(echo $HOST_PORT | cut -d ':' -f 1)
PORT=$(echo $HOST_PORT | cut -d ':' -f 2)

if [ -z "$HOST" ] || [ -z "$PORT" ]; then
  echo "Error: You must specify host:port"
  exit 1
fi

echo "Waiting for $HOST:$PORT..."
for i in $(seq $TIMEOUT); do
  if nc -z -w 1 $HOST $PORT; then
    echo "$HOST:$PORT is available!"
    exec "$@"
    exit 0
  fi
  sleep 1
done

echo "Timeout exceeded while waiting for $HOST:$PORT"
exit 1