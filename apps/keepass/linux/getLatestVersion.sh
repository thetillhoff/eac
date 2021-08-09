#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
curl --silent "https://api.github.com/repos/keepass/keepass/releases/latest" |
	grep 'tag_name' | # f.e. '  "tag_name": "v1.2.3",'
	cut -d'"' -f4 | # f.e. 'v1.2.3'
	cut -c2- # f.e. '1.2.3'
