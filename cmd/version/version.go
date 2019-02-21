package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

const CgcCliVersion = "v0.1"

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of CGC CLI tool",
		Long:  `All software has versions.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(CgcCliVersion)
		},
	}

	return cmd
}
