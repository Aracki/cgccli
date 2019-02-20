package version

import (
	"fmt"
	"github.com/aracki/cgc/info"
	"github.com/spf13/cobra"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of CGC CLI tool",
		Long:  `All software has versions.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(info.CGC_CLI_VERSION)
		},
	}

	return cmd
}
