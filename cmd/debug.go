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

	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Builds HTML in debug mode",
	Long:  `Generates warnings for all missing references, the build stops at the first warning.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		color.Yellow("Building in debug mode")
		fmt.Println("Hold on, this will take some seconds")
		buildDebug()

	},
}

func buildDebug() {
	cmdStr := "docker run -v `pwd`/docs:/build/docs testthedocs/plone-docsbuilder testbuild"
	//out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	exec.Command("/bin/sh", "-c", cmdStr).Output()
	//fmt.Printf("%s", out)
}

func init() {
	RootCmd.AddCommand(debugCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// debugCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// debugCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
