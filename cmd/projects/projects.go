package projects

import (
	"fmt"
	"github.com/aracki/cgc/api"
	"github.com/spf13/cobra"
)

func NewCmdProjects() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects",
		Short: "Cancer Genomics Cloud projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			projects, err := api.GetProjects()
			if err != nil {
				return err
			}
			output(projects)
			return nil
		},
	}

	return cmd
}

func output(projects []api.Project) {

	for i := 0; i < len(projects); i++ {
		fmt.Printf("id: %s \n", projects[i].Id)
		fmt.Printf("name: %s \n\n", projects[i].Name)
	}
}
