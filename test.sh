#!/bin/bash

set -a
TM_DB_URL="postgres://postgres:password@localhost:5432/tm_test"
TM_API_KEY="API_KEY"
set +a

go test $@
