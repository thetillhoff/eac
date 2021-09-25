#!/bin/sh

# $1==appVersion
# $2==tmpFolder, gets created before this script is called and deleted afterwards

wget -q "https://github.com/thetillhoff/python/releases/download/v$1/python-linux-amd64" -O "$2/python-linux-amd64"
sudo install -D -g root -o root -m 755 "$2/python-linux-amd64" "/usr/local/bin/python"
# install SRC DEST: copies SRC to DEST, changes DEST permissions, owners in one command
# install -D: create all leading components of DEST except the last
# install -g root: change group ownership of DEST
# install -o root: change user ownership of DEST
# install -m XXX: set permissions
