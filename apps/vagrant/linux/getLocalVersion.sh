#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
vagrant --version | # f.e. 'vagrant version 1.2.3'
	cut -d' ' -f 2 # f.e. '1.2.3'
