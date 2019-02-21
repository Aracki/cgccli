package projects

import (
	"github.com/aracki/cgccli/api/projects"
	"github.com/spf13/cobra"
)

func NewCmdProjects() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects",
		Short: "Cancer Genomics Cloud projects",
		Long: `Projects are the core building blocks of the CGC Platform. 
Each project corresponds to a distinct scientific investigation, 
serving as a container for its data, analysis pipelines, and results. 
Projects are shared only by designated project members.`,
	}

	cmd.AddCommand(NewCmdProjectsList())

	return cmd
}

func NewCmdProjectsList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			allProjects, err := projects.GetProjects()
			if err != nil {
				return err
			}
			return printProjects(allProjects)
		},
	}

	return cmd
}
