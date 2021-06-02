#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
kubectl version --client | # f.e. 'kubectl version 1.2.3'
	cut -d'"' -f 6 | # f.e. 'v1.2.3'
  cut -c2- # f.e. '1.2.3'
