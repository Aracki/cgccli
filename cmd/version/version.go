package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of cgc cli tool",
		Long:  `All software has versions.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v0.1")
		},
	}

	return cmd
}
