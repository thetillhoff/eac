#!/bin/sh

# $1 is the path to remove from PATH

filePath="$HOME/.profile"
line="export PATH=\"\$PATH;$1\""

# Check if line exists
if grep -Fxq "$line" $filePath; then
  # Remove it
  sed -i "\%$line%d" $filePath
else
  echo "Skipped removeFromUserPath, since line doesn't exist in '$filePath'."
fi

# Make change for user available right away
source $filePath
