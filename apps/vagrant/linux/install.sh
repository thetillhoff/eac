#/bin/sh

# $1==appVersion
# $2==tmpFolder, gets created before this script is called and deleted afterwards

# curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
# sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
# sudo apt-get update && sudo apt-get install vagrant

sudo apt-get install -y wget unzip # install dependencies
wget -q "https://releases.hashicorp.com/vagrant/$1/vagrant_$1_linux_amd64.zip" -O "$2/vagrant_v$1_linux_amd64.tar.gz"
mkdir $2/extracted
unzip "$2/vagrant_v$1_linux_amd64.tar.gz" -d "$2/extracted"
sudo install -D -g root -o root -m 755 "$2/extracted/vagrant" "/usr/local/bin/vagrant"
# install SRC DEST: copies SRC to DEST, changes DEST permissions, owners in one command
# install -D: create all leading components of DEST except the last
# install -g root: change group ownership of DEST
# install -o root: change user ownership of DEST
# install -m XXX: set permissions
