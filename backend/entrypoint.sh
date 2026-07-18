#!/bin/sh
set -eu
cd /app/database/goose
./migrate up
cd /app
exec ./main
