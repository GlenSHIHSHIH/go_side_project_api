#!/bin/bash
LIST=`crontab -l`

SOURCE="/opt/golang-shopee/compiler/golang_file shopee-data-to-db"
# # SOURCE="/opt/golang-shopee/run_golang_files.sh >> /var/log/daily-backup.log 2>&1"

if echo "$LIST" | grep -q "$SOURCE"; then
   echo "The back job had been added.";
else
   crontab -l | { cat; echo "* * * * * $SOURCE"; } | crontab -

fi

# * * * * * echo "hello" >> /var/log/daily-backup.log 2>&1

# cd /go/golang-shopee/docker_build_go
# nohup ./compiler/golang_file shopee-api &
