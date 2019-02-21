package files

import (
	"github.com/aracki/cgccli/api/files"
	"github.com/spf13/cobra"
)

func NewCmdFiles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "files",
		Short: "Cancer Genomics Cloud files",
		Long: `This command enable you to manage project files and their metadata.
You can return lists of file IDs, copy files between projects,
edit their metadata, and delete files, and download them.`,
	}

	cmd.AddCommand(NewCmdFilesList())
	return cmd
}

func NewCmdFilesList() *cobra.Command {
	var project string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all files in a project",
		Long: `This call returns a list of all files in a specified project with specified properties that you can access.
For each file, the call returns its ID and filename`,
		RunE: func(cmd *cobra.Command, args []string) error {
			allFiles, err := files.GetFiles(project)
			if err != nil {
				return err
			}
			return printFiles(allFiles)
		},
	}

	cmd.Flags().StringVarP(&project, "project", "p", "", "Retrieve the files belonging to the specified project.")
	cmd.MarkFlagRequired("project")

	return cmd
}
