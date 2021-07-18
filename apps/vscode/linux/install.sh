#/bin/sh

# $1==appVersion
# $2==tmpFolder, gets created before this script is called and deleted afterwards

wget -q "https://go.microsoft.com/fwlink/?LinkID=760868" -O "$2/vscode.deb"
sudo dpkg -i "$2/vscode.deb"
