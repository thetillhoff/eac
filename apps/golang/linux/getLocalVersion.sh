#!/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
go version | # f.e. 'go version go1.2.3 linux/amd64'
	cut -d' ' -f3 | # f.e. 'go1.2.3'
  cut -c3- || # f.e. '1.2.3'
  printf "\n" # if retrieval failed
