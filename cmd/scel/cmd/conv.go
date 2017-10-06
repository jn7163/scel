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
	"github.com/spf13/cobra"
)

var excludeExt = false
var optimizeExt = false
var excludeCommonPy = false
var optimize []string
var exclude []string

// convCmd represents the conv command
var convCmd = &cobra.Command{
	Use:     "conv [in] [out]",
	Args:    cobra.MinimumNArgs(2),
	Aliases: []string{"c"},
	Short:   "Conversion between different format",
	Long: `Convert scel to pb.

Optimize:
	e - ext data
Exclude:
	e - ext data
	P - common pinyin table`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range optimize {
			switch v {
			case "e":
				optimizeExt = true
			default:
				panic("Unknown optimize option: " + v)
			}
		}
		for _, v := range exclude {
			switch v {
			case "e":
				excludeExt = true
			case "P":
				excludeCommonPy = true
			default:
				panic("Unknown exclude option: " + v)
			}
		}

		open(args[0])

		if excludeExt {
			doExcludeExt()
		} else if optimizeExt {
			doOptimizeExt()
		}

		if excludeCommonPy {
			doExcludeCommonPy()
		}

		write(args[1])
	},
}

func init() {
	RootCmd.AddCommand(convCmd)

	convCmd.Flags().StringArrayVarP(&optimize, "optimize", "o", nil, "Optimize")
	convCmd.Flags().StringArrayVarP(&exclude, "exclude", "e", nil, "Exclude")
}
