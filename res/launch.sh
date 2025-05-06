#!/bin/bash

pkill -f hwt

if [ "$#" -gt 0 ]; then
 /mnt/SDCARD/Apps/DummyBanner/dummy "$@"
else
  progdir=$(dirname "$0")
  cd $progdir
  ./dummy
fi