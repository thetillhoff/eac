#!/bin/sh

# $1 is the path to add to PATH

filePath="/etc/profile.d/$(basename $1).sh"
line="export PATH=\"\$PATH;$1\""

# Check that file doesn't exist
if [ ! -e "${filePath}" ]; then
  # Create it
  printf "%s\n" "$line" > $filePath
else
  echo "Skipped addToGlobalPath, since file already exists at '$filePath'."
fi
