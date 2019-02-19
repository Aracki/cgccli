package cmd

import (
	"fmt"
	"github.com/aracki/cgc/api"
	"github.com/spf13/cobra"
	"log"
)

func NewCmdProjects() *cobra.Command {
	cmd := &cobra.Command{
		Use: "projects",
		Short: "CGC projects",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("y")
			log.Println(api.GetProjects())
		},
	}

	return cmd
}
