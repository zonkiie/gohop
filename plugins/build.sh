#!/bin/sh

rm -fv *.so

for i in *.go; do
	echo "building " "$i"
	go build -buildmode=plugin "$i"
done

