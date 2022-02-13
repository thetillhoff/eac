#!/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
code --version | head -n 1 || # f.e. '1.2.3'
	printf "\n" # if retrieval failed
