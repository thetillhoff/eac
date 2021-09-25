#!/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
#sudo apt-get install curl -y
curl --silent "https://releases.hashicorp.com/vagrant/" |
	grep '>vagrant_' |
	cut -d'_' -f2 |
	cut -d'<' -f1 | # f.e. '1.2.3\n1.2.2\n...'
  head -n 1 # f.e. '1.2.3'
