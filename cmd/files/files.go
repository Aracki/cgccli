package files

import (
	"github.com/aracki/cgccli/api/files"
	"github.com/spf13/cobra"
)

var (
	filesCmd   = "files"
	filesShort = "Manage CGC files"
	filesLong  = `This command enable you to manage project files and their metadata.
You can return lists of file IDs, copy files between projects,
edit their metadata, and delete files, and download them.`

	filesListCmd           = "list"
	filesListFlagProject   = "project"
	filesListFlagProjectSh = "p"
	filesListShort         = "Lists the files in the specified project"
	filesListLong          = `This call returns a list of all files in a specified project with specified properties that you can access.
The project to find files from is specified as a query parameter in the call. 
Further file properties to filter by can also be specified as query parameters.`

	filesStatCmd        = "stat"
	filesStatFlagFile   = "file"
	filesStatFlagFileSh = "f"
	filesStatShort      = "Get file details"
	filesStatLong       = `This call returns details about a specified file. The call returns the file's name, its tags, and all of its metadata.
Files are specified by their IDs, which you can obtain by making the API call to list files in a project.`
)

func NewCmdFiles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   filesCmd,
		Short: filesShort,
		Long:  filesLong,
	}

	cmd.AddCommand(NewCmdFilesList())
	cmd.AddCommand(NewCmdFilesStat())

	return cmd
}

func NewCmdFilesList() *cobra.Command {

	var project string

	cmd := &cobra.Command{
		Use:   filesListCmd,
		Short: filesListShort,
		Long:  filesListLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			allFiles, err := files.GetFiles(project)
			if err != nil {
				return err
			}
			return printFiles(allFiles)
		},
	}
	cmd.Flags().StringVarP(&project, filesListFlagProject, filesListFlagProjectSh, "", filesListShort)
	cmd.MarkFlagRequired(filesListFlagProject)
	return cmd
}

func NewCmdFilesStat() *cobra.Command {

	var fileId string

	cmd := &cobra.Command{
		Use:   filesStatCmd,
		Short: filesStatShort,
		Long:  filesStatLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			fDetails, err := files.GetFileDetails(fileId)
			if err != nil {
				return err
			}
			return printFileDetails(*fDetails)
		},
	}
	cmd.Flags().StringVarP(&fileId, filesStatFlagFile, filesStatFlagFileSh, "", filesStatShort)
	cmd.MarkFlagRequired(filesStatFlagFile)
	return cmd
}
