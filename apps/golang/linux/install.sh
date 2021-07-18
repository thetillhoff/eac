#/bin/sh

# $1==appVersion
# $2==tmpFolder, gets created before this script is called and deleted afterwards

# from https://golang.org/doc/install

wget -q "https://golang.org/dl/go$1.linux-amd64.tar.gz" -O "$2/go$1.linux-amd64.tar.gz"
sudo rm -rf /usr/local/go # Delete previous installation if any
sudo tar -xzf "$2/go$1.linux-amd64.tar.gz" -C /usr/local
