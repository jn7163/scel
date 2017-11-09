// Copyright © 2017 wener <wenermail@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wenerme/scel/util"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:     "info [in...]",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"i"},
	Short:   "Show info about scel dict",
	Long:    `Show name, type, example, desc, counts.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		for _, v := range args {
			data, err = scelutil.Read(v)
			if err != nil {
				logrus.WithError(err).Fatal("failed to open file")
			}
			fmt.Printf(`file: %v
name: %v
type: %v
desc: %v
e.g.: %v
`,
				v,
				data.Info.Name,
				data.Info.Type,
				data.Info.Description,
				data.Info.Example,
			)
		}

	},
}

func init() {
	RootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
