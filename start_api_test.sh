#/bin/bash

export ENV="test"

LOG_FILE="log-"$JOB_ENV"-"`date '+%Y%m%d'`".txt"

go run app.go #2>&1 | tee 