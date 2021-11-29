#!/bin/bash

# run systemd
# exec /usr/sbin/init
exec /usr/sbin/init
exec "$@"
systemctl
# run the command given as arguments from CMD
# exec systemctl