#!/bin/sh

# The Build script

sh cleanup.sh
echo "get"
go get
echo "build"
go build
echo "building plugins"
cd plugins
sh build.sh
cd ..
echo "done"
