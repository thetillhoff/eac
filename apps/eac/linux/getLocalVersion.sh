#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
eac --version | # f.e. 'eac version 1.2.3'
	cut -d' ' -f 3 # f.e. '1.2.3'
