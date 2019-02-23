// Package projects provides projects root command and all its subcommands.
package projects

import (
	"github.com/aracki/cgccli/api/projects"
	"github.com/spf13/cobra"
)

var (
	projectsCmd   = "projects"
	projectsShort = "Cancer Genomics Cloud projects"
	projectsLong  = `Projects are the core building blocks of the CGC Platform. 
Each project corresponds to a distinct scientific investigation, 
serving as a container for its data, analysis pipelines, and results. 
Projects are shared only by designated project members.`

	projectsListCmd   = "list"
	projectsListShort = "List all your projects"
	projectsListLong  = `This call returns a list of all projects you are a member of. 
Each project's project_id and URL on the CGC will be returned.`
)

// NewCmdProjects is the root command for projects.
// All subcommands regarding to projects are added here.
func NewCmdProjects() *cobra.Command {
	cmd := &cobra.Command{
		Use:   projectsCmd,
		Short: projectsShort,
		Long:  projectsLong,
	}

	cmd.AddCommand(NewCmdProjectsList())

	return cmd
}

// NewCmdProjectsList lists all the projects.
func NewCmdProjectsList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   projectsListCmd,
		Short: projectsListShort,
		Long:  projectsListLong,
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
