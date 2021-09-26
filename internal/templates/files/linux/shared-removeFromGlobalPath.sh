#!/bin/sh

# $1 is the path to remove from PATH

filePath="/etc/profile.d/$(basename $1).sh"
line="export PATH=\"\$PATH;$1\""

# Check that file exists
if [ -e "${filePath}" ]; then
  # Remove it
  rm $filePath
else
  echo "Skipped removeFromGlobalPath, since file doesn't exist at '$filePath'."
fi
