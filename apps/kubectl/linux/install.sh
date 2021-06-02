#/bin/sh

# $1==appName
# $2==tmpFolder, gets created before this script is called and deleted afterwards

wget -q "https://dl.k8s.io/release/v$1/bin/linux/amd64/kubectl" -O "$2/kubectl"
sudo install -D -g root -o root -m 755 "$2/kubectl" "/usr/local/bin/kubectl"
# install SRC DEST: copies SRC to DEST, changes DEST permissions, owners in one command
# install -D: create all leading components of DEST except the last
# install -g root: change group ownership of DEST
# install -o root: change user ownership of DEST
# install -m XXX: set permissions
