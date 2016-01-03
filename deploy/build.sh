#!/usr/bin/env bash

# Example for create freebsd build
# GOOS=freebsd GOARCH=amd64 deploy/build_release.sh

# export GIT_SSL_NO_VERIFY=true

set +x

function safemake {
  if [ "$?" != "0" ]; then
    echo "ERROR: BUILD FAILED"
    exit 1
  fi
}

# F_OS="freebsd"
# F_ARCH="amd64"
cd "$(dirname $0)"
BASEDIR="$(dirname $(pwd))"
CURDIR=`pwd`
# GOBASE="$(dirname $(dirname `which go`))/src/"

cd $BASEDIR

echo "0)"
echo "Create .build structure"

rm -R .build
mkdir -p .build/logs/

cp -R config .build/
cp -R templates .build/
cp -R public .build/
cp deploy/run.sh .build/run.sh
cp README.md .build/

# curl "https://www.maxmind.com/app/geoip_download?edition_id=GeoIP2-ISP&date=20150721&suffix=tar.gz&license_key=oeEqpP5QI21N" \
#   | tar -xf- -C .build/

# mv .build/GeoIP2-ISP*/*ISP*.mmdb .build/conf/GeoIP2-ISP.mmdb
# rm -R .build/GeoIP2-ISP*

echo ""
echo "1)"
echo "Print GO environment"
goproj go env

# Export vars and remove apostraphs
for l in `goproj go env`; do
  if [[ $l =~ ^[A-Za-z0-9_]+= ]]; then
    IFS='=' read -a array <<< "$l"
    v=`echo ${array[1]} | sed -e 's/^"//g' -e 's/"$//g'`
    declare "${array[0]}=$v"
  fi
done


echo ""
echo "2)"
echo "Prepare project for $GOARCH/$GOOS"

if [ "$GOHOSTOS" != "$GOOS" ] || [ "$GOHOSTARCH" != "$GOARCH" ]
then
  cd "$GOROOT/src"
  ./make.bash -v --no-clean
  safemake

  cd $CURDIR
else
  echo "Is current Arch"
fi


echo ""
echo "3)"
echo "Install dependencies"
# goproj get
safemake


echo ""
echo "4)"
echo "Build project $BASEDIR"
goproj build_deploy
safemake


echo ""
echo "5)"
echo "Print file info"
if [ ! -f "$BASEDIR/.build/www" ]; then
  echo "ERROR: FILE DONT EXISTS"
  go version
  exit 1
fi

file .build/www
safemake
