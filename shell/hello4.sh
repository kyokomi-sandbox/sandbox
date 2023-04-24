#!/bin/bash
echo "You are using $(basename $0)"
test -z "$1" || echo "Hello $*"
exit 0

