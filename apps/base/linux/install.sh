#!/bin/sh

# $1==appVersion
# $2==tmpFolder, gets created before this script is called and deleted afterwards

# Install basic system tools on a Debian-based OS
sudo apt-get install -y \
  apt-transport-https \
  ca-certificates \
  gnupg2 \
  software-properties-common \
  curl \
  wget

# Uninstall annoying 'call-home'-software
sudo apt-get purge -y \
  kerneloops
  whoopsie
