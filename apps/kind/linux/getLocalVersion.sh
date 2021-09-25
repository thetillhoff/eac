#!/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
kind version | # f.e. 'kind v1.2.3 go1.16.4 linux/amd64'
  cut -d' ' -f2 | # f.e. 'v1.2.3'
  cut -c2- # f.e. '1.2.3'
