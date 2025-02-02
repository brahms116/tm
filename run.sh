#!/bin/bash

set -a
TM_DB_URL="postgres://postgres:password@localhost:5432/tm_dev"
set +a

go run $@
