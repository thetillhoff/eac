/*
Copyright © 2022 Till Hoffmann <till.hoffmann@enforge.de>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	eac "github.com/thetillhoff/eac/pkg/eac"
)

// listDownloadedVersionsCmd represents the versions command
var listDownloadedVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List all downloaded/cached version for each app.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		// Passing flags & viper config
		eac.Verbose = verbose || viper.GetBool("verbose")

		appListWithVersions, err := eac.ListDownloadedApps(true)
		if err != nil {
			logger.Fatal(err.Error())
		}

		for _, entry := range appListWithVersions {
			fmt.Println(entry)
		}

	},
}

func init() {
	listDownloadedCmd.AddCommand(listDownloadedVersionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
