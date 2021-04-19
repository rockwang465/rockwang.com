#!/usr/bin/env bash

# rsync releases files
# rsync -azP /data/packages/sensenebula/releases/ 10.151.5.150:/data/backup/releases
# a = rlptgoD, r: recursive, l: soft-link;  directory and hide file can not rsync
rsync -ptgoDzP /data/packages/sensenebula/releases/* 10.151.5.150:/data/backup/releases