#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
helm version | # f.e. 'helm version 1.2.3'
	cut -d'"' -f2 |
  cut -d'v' -f2) # f.e. '1.2.3'
