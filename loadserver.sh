#!/bin/sh

echo $1
/usr/bin/stress-ng --cpu 1 --vm 1 --vm-bytes 1024M --hdd 1 --fork 1 --switch 1 --timeout 600 --metrics &
