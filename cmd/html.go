// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	silenceMode bool
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "html",
	Short: "Builds HTML",
	Long:  `Builds the docs in HTML, ignoring warnings about reST syntax.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildHTML()
	},
}

func buildHTML() {
	if silenceMode {
		//silenceMode = true
		//c.watchConfig()
		//fmt.Println("Hiding output")
		cmdStr := "docker run -v `pwd`:/build/docs testthedocs/plone-docsbuilder html-quiet"
		exec.Command("/bin/sh", "-c", cmdStr).Output()
	} else {
		fmt.Println("Building HTML")
		fmt.Println("Hold on, this will take some seconds")
                cmdStr := "docker run -v `pwd`:/build/docs testthedocs/plone-docsbuilder html"
		out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
		//exec.Command("/bin/sh", "-c", cmdStr).Output()
		fmt.Printf("%s", out)
	}
}

func init() {
	RootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	buildCmd.Flags().BoolVarP(&silenceMode, "silence", "s", false, "hides terminal output")

	if silenceMode {
		//silenceMode = true
		//c.watchConfig()
		fmt.Println("Hiding output")
	}
}
