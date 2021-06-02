#/bin/sh

wget -q "https://github.com/thetillhoff/eac/releases/download/v$1/eac-linux-amd64" -O "$2/eac-linux-amd64"
sudo install -D -g root -o root -m 755 "$2/eac-linux-amd64" "/usr/local/bin/eac"
# install SRC DEST: copies SRC to DEST, changes DEST permissions, owners in one command
# install -D: create all leading components of DEST except the last
# install -g root: change group ownership of DEST
# install -o root: change user ownership of DEST
# install -m XXX: set permissions