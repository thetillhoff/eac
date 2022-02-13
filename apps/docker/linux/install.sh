#!/bin/sh

# $1==appVersion
# $2==tmpFolder, gets created before this script is called and deleted afterwards

# Uninstall older version
sudo apt-get remove -y docker docker-engine docker.io containerd runc

# Install prerequisites
sudo apt-get install -y ca-certificates curl gnupg lsb-release

# Install apt repository-key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install application
sudo apt update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io
