#!/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
gh --version | head -n1 | # 
  cut -d' ' -f 3 || #
	printf "\n" # if retrieval failed
