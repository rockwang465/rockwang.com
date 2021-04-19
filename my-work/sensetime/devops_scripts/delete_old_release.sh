#!/usr/bin/env bash

find /data/packages/sensenebula/releases -type f -mtime +15 | xargs rm -f