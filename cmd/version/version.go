// Package version provides version command.
package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CGCCliVersion is a current program version.
// TODO this should be outside of source code
const CGCCliVersion = "v0.0.2"

// NewCmdVersion is the command for version.
func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of CGC CLI tool",
		Long:  `All software has versions.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(CGCCliVersion)
		},
	}

	return cmd
}
