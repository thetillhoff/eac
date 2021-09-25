#!/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
curl --silent "https://api.github.com/repos/golang/go/git/refs/tags" |
	grep 'ref": "refs/tags/go' | # f.e. ['    "ref": "refs/tags/go1.2.3",']
  tail -n 1 |
	cut -d'"' -f4 | # f.e. 'v1.2.3'
	cut -c$(printf "\nrefs/tags/go" | wc -m)- # f.e. '1.2.3'
