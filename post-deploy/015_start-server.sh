#!/usr/bin/env bash

MONEX_HOME="/apps/core"
cd $MONEX_HOME

echo "$MONEX_HOME"
echo "Starting new server..."

# Store output in a logfile and save the PID to a file so we can kill the process later
./dist/server >> /apps/core/gosvr.log 2>&1 & echo $! > /apps/core/gosvr.pid

echo "Verifying the server is running..."
if ! ps -p $(cat /apps/core/gosvr.pid) > /dev/null 2>&1; then
    echo "ERROR: Process is not running!"
    exit 1
fi
echo "OK!"
