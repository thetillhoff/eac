#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
jetbrains-toolbox --version | # f.e. 'jetbrains-toolbox version 1.2.3'
	cut -d' ' -f3 || # f.e. '1.2.3'
	printf "\n" # if retrieval failed