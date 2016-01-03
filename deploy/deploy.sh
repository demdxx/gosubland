#!/usr/bin/env bash

set +xv

DEPLOY_DIR=$(cd "$(dirname "$0")"; pwd)
PROJECT_DIR=`dirname $DEPLOY_DIR`
cd $PROJECT_DIR

APP='geniusrabbit.com'
SRCD="${PROJECT_DIR}/.build/"
DSTD="/var/www/${APP}"
SYNCUSER="root"
SYNCOPTS="--delete"
SYNCHOSTS="46.101.187.170"

# Sync project

echo "Start sync"

for SYNCHOST in ${SYNCHOSTS}; do

  rsync -vizap -c -e "ssh -o Compression=no -o StrictHostKeyChecking=no -x" ${SYNCOPTS} ${SRCD} ${SYNCUSER}@${SYNCHOST}:${DSTD}
  # ssh ${SYNCUSER}@${SYNCHOST} 'cd ${DSTD}/deploy && sh build.sh; sudo /etc/init.d/supervisor restart; sudo /etc/init.d/nginx reload;'

done
