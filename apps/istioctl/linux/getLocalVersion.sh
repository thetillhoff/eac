#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
istioctl --version | # f.e. 'istioctl version 1.2.3'
	cut -d' ' -f3 || # f.e. '1.2.3'
	printf "\n" # if retrieval failed
