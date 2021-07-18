package apps

//TODO insert shell at top of file '#<shell>' .f.e '#/bin/sh'
// -> only for linux, not on windows

var demoFiles = map[string]map[string]string{ //TODO: add proper initial scripts
	"linux": {
		"configure.sh": `#/bin/sh

echo "No configuration needed for app '%[1]v'."
`,
		"getLatestVersion.sh": `#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
curl --silent "https://api.github.com/repos/%[1]v/%[1]v/releases/latest" |
	grep 'tag_name' | # f.e. '  "tag_name": "v1.2.3",'
	cut -d'"' -f4 | # f.e. 'v1.2.3'
	cut -c2- # f.e. '1.2.3'
`,
		"getLocalVersion.sh": `#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
%[1]v --version | # f.e. '%[1]v version 1.2.3'
	cut -d' ' -f3 || # f.e. '1.2.3'
	printf "\n" # if retrieval failed
`,
		"install.sh": `#/bin/sh

# $1==appVersion
# $2==tmpFolder, gets created before this script is called and deleted afterwards

wget -q "https://github.com/%[1]v/%[1]v/releases/download/v$1/%[1]v-linux-amd64" -O "$2/%[1]v-linux-amd64"
sudo install -D -g root -o root -m 755 "$2/%[1]v-linux-amd64" "/usr/local/bin/%[1]v"
# install SRC DEST: copies SRC to DEST, changes DEST permissions, owners in one command
# install -D: create all leading components of DEST except the last
# install -g root: change group ownership of DEST
# install -o root: change user ownership of DEST
# install -m XXX: set permissions
`,
		"uninstall.sh": `#/bin/sh

sudo rm /usr/local/bin/%[1]v
`,
	},
	"darwin": {
		"configure.sh": `#/bin/sh
echo "This script is called to configure of app %[1]v."
`,
		"getLatestVersion.sh": `#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example:
curl --silent "https://api.github.com/repos/%[1]v/%[1]v/releases/latest" |
	grep 'tag_name' | # f.e. '  "tag_name": "v1.2.3",'
	cut -d'"' -f4 | # f.e. 'v1.2.3'
	cut -c2- # f.e. '1.2.3'
`,
		"getLocalVersion.sh": `#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example: '%[1]v --version' prints '%[1]v version 1.2.3', then the following line would be sufficient.
printf "$(%[1]v --version | cut -d' ' -f 3)\n"
`,
		"install.sh": `#/bin/sh

binaryName=$(openssl rand -base64 16)%[1]v
wget "https://github.com/thetillhoff/%[1]v/releases/download/v$1/%[1]v-linux-amd64" -O $binaryName
chmod +x $binaryName
sudo mv $binaryName /usr/local/bin/%[1]v
`,
		"uninstall.sh": `#/bin/sh

sudo rm /usr/local/bin/%[1]v		
`,
	},
	"windows": {
		"configure.ps1": `Write-Host "This script is called to configure app %[1]v."
`,
		"getLatestVersion.ps1": `Write-Host "This script is called to getLatestVersion of app %[1]v."
`,
		"getLocalVersion.ps1": `Write-Host "This script is called to getLocalVersion of app %[1]v."
`,
		"install.ps1": `Write-Host "This script is called to install app %[1]v."
`,
		"uninstall.ps1": `Write-Host "This script is called to uninstall app %[1]v."
`,
	},
}
