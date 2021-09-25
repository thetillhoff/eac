#!/bin/sh

# $1==appVersion
# $2==tmpFolder, gets created before this script is called and deleted afterwards

sudo apt-get install -y \
  qemu-kvm
  libvirt-clients
  libvirt-daemon-system
  virt-manager

sudo usermod -a -G libvirt $USER
sudo usermod -a -G libvirt-qemu $USER
sudo usermod -a -G kvm $USER
