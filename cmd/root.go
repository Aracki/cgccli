// Copyright © 2019 aracki.ivan@gmail.com
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

// Package cmd provides all CLI commands.
package cmd

import (
	"fmt"
	"os"

	"github.com/aracki/cgccli/cmd/files"
	"github.com/aracki/cgccli/cmd/projects"
	"github.com/aracki/cgccli/cmd/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cgccli",
	Short: "The Command Line Interface tool for the Cancer Genomics Cloud Public API",
	Long:  `Use the cgccli to access the Platform via the API.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(version.NewCmdVersion())
	rootCmd.AddCommand(projects.NewCmdProjects())
	rootCmd.AddCommand(files.NewCmdFiles())

	// make --token PERSISTENT & GLOBAL flag and put it in the Viper registry
	rootCmd.PersistentFlags().String("token", "", "Your authentication token that encodes your CGC credentials.")
	rootCmd.MarkPersistentFlagRequired("token")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}
