#!/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
curl --silent "https://storage.googleapis.com/kubernetes-release/release/stable.txt" | # f.e. 'v1.2.3'
	cut -c2- # f.e. '1.2.3'
