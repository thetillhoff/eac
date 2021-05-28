package apps

//TODO insert shell at top of file '#<shell>' .f.e '#/bin/sh'
// -> only for linux, not on windows

var demoFiles = map[string]map[string]string{ //TODO: add proper initial scripts
	"linux": {
		"configure.sh": `#/bin/sh
echo "This script is called to configure of app %v"
`,
		"getLatestVersion.sh": `#/bin/sh
echo "This script is called to getLatestVersion of app %v"
`,
		"getLocalVersion.sh": `#/bin/sh

# This script has to print (only!) the version of the app to stdout. Trailing newlines are fine though.
# Example: '%v version' prints '%v version 1.2.3', then the following line would be sufficient.
printf "$(%v version | cut -d' ' -f 3)\n"
`,
		"install.sh": `#/bin/sh
echo "This script is called to install of app %v"
`,
		"uninstall.sh": `#/bin/sh
echo "This script is called to uninstall of app %v"
`,
	},
	"darwin": {
		"configure.sh": `#/bin/sh
echo "This script is called to configure of app %v"
`,
		"getLatestVersion.sh": `#/bin/sh
echo "This script is called to getLatestVersion of app %v"
`,
		"getLocalVersion.sh": `#/bin/sh
echo "This script is called to getLocalVersion of app %v"
`,
		"install.sh": `#/bin/sh
echo "This script is called to install of app %v"
`,
		"uninstall.sh": `#/bin/sh
echo "This script is called to uninstall of app %v"
`,
	},
	"windows": {
		"configure.ps1": `Write-Host "This script is called to configure of app %v"
`,
		"getLatestVersion.ps1": `Write-Host "This script is called to getLatestVersion of app %v"
`,
		"getLocalVersion.ps1": `Write-Host "This script is called to getLocalVersion of app %v"
`,
		"install.ps1": `Write-Host "This script is called to install of app %v"
`,
		"uninstall.ps1": `Write-Host "This script is called to uninstall of app %v"
`,
	},
}
