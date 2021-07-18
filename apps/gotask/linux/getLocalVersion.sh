#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
task --version | # f.e. 'Task version: v1.2.3 ...'
	cut -d' ' -f3 | # f.e. 'v1.2.3'
  cut -c2- || # f.e. '1.2.3'
	printf "\n" # if retrieval failed
