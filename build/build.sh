#!/bin/bash

version=$1
bin_name=$2

if [ ! $version ]
then
  version=test-1.0
fi

if [ ! $bin_name ]
then
  bin_name=student_server
fi

echo "start to build[$bin_name], version[$version]"

cd ../cmd/server && \
# shellcheck disable=SC2027
go build -ldflags "-X 'main._version="$version"' -X 'main._goVersion=$(go version)' -X 'main._gitHash=$(git show -s --format=%H)' -X 'main._buildTime=$(date +'%Y-%m-%d %T')'" \
         -o $bin_name

echo "success"
cd ../../
echo "$(pwd)"
create_pitrix_student_server="build/supervisor/pitrix/lib/student_server"
create_pitrix_conf="build/supervisor/pitrix/config"
mkdir ${create_pitrix_student_server} -p
mkdir ${create_pitrix_conf} -p
cp ./main.go ${create_pitrix_student_server}
cp config/stu_apiserver.yaml create_pitrix_conf
cd build
chmod +755 supervisor/DEBIAN -R
dpkg -b supervisor student_server.deb