#!/bin/sh

sudo rm /usr/local/bin/{{ index . "name" }}

sudo ~/.eac/shared/removeFromGlobalPath.sh /usr/local/bin/{{ index . "name" }}
