#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
gopass --version | # f.e. 'gopass 1.2.3+... ...'
	cut -d' ' -f2 | # f.e. '1.2.3+...'
  cut -d'+' -f1 || # f.e. '1.2.3'
	printf "\n" # if retrieval failed
