#!/bin/sh

sudo apt-get remove -y docker-ce docker-ce-cli containerd.io
sudo rm /etc/apt/sources.list.d/docker.list
sudo rm /usr/share/keyrings/docker-archive-keyring.gpg
