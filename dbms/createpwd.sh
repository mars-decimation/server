#!/bin/sh

umask 377
cat /dev/urandom | tr -cd "[:alnum:]" | head -c 32 > root.txt
