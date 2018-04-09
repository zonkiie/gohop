#!/bin/sh

echo "clean"
go clean
cd plugins
sh ./cleanup.sh
cd ..
echo "done"
