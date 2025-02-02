#!/bin/bash

set -a
TM_DB_URL="postgres://postgres:password@localhost:5432/tm_test"
set +a

go test $@
