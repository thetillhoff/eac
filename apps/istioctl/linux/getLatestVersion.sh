#!/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
curl --silent "https://api.github.com/repos/istio/istio/releases/latest" |
	grep 'tag_name' | # f.e. '  "tag_name": "1.2.3",'
	cut -d'"' -f4 # f.e. '1.2.3'
