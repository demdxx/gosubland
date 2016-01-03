#!/usr/bin/env bash
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/www" -basedir '/var/www/geniusrabbit.com' --mode=prod --fastcgi > /dev/null 2>&1 &
