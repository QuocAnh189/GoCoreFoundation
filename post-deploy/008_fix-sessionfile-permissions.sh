#!/usr/bin/env bash

MONEX_HOME="/apps/core-foundation"
cd $MONEX_HOME

echo "$MONEX_HOME"
echo "Fixing session file permissions..."

# Fix sessionfile permissions
chmod 775 $MONEX_HOME/data/*.dat

echo "Session file permissions fixed!"