#/bin/sh

wget -q "https://get.helm.sh/helm-v$1-linux-amd64.tar.gz" -O "$2/helm-v$1-linux-amd64.tar.gz"
mkdir $2/extracted
tar -xzf "$2/helm-v$1-linux-amd64.tar.gz" -C "$2/extracted" --strip-components 1
sudo install -D -g root -o root -m 755 "$2/extracted/helm" "/usr/local/bin/helm"
# install SRC DEST: copies SRC to DEST, changes DEST permissions, owners in one command
# install -D: create all leading components of DEST except the last
# install -g root: change group ownership of DEST
# install -o root: change user ownership of DEST
# install -m XXX: set permissions
