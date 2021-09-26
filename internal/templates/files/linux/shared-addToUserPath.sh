#!/bin/sh

# $1 is the path to add to PATH

filePath="$HOME/.profile"
line="export PATH=\"\$PATH;$1\""

# Check that file not already exists
if [ ! -e "${filePath}" ]; then
	# Create it
	printf "%s\n" "$line" > $filePath
else
	# Check if line exists
	if ! grep -Fxq "$line" $filePath; then
		# Append to it
		printf "%s\n" "$line" >> $filePath
	else
		echo "Skipped addToUserPath, since line already exists in '$filePath'."
	fi
fi

# Make change for user available right away
source $filePath
