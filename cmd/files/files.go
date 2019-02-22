// Package files provides files root command and all its subcommands.
package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aracki/cgccli/api/files"
	"github.com/spf13/cobra"
	"strings"
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

	filesUpdateCmd        = "update"
	filesUpdateFlagFile   = "file"
	filesUpdateFlagFileSh = "f"
	filesUpdateShort      = "Update file details"
	filesUpdateLong       = `This call updates the name, the full set metadata, and tags for a specified file.
Files are specified by their IDs, which you can obtain by making the API call to list files in a project.
A full list of metadata fields and their permissible values on the CGC is available on the page TCGA Metadata [` + metadataLinks + `]`
	metadataLinks = "https://docs.cancergenomicscloud.org/v1.0/docs/metadata-for-private-data"
)

// NewCmdFiles is the root command for files.
// All subcommands regarding to files are added here.
func NewCmdFiles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   filesCmd,
		Short: filesShort,
		Long:  filesLong,
	}

	cmd.AddCommand(NewCmdFilesList())
	cmd.AddCommand(NewCmdFilesStat())
	cmd.AddCommand(NewCmdFilesUpdate())

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

// NewCmdFilesUpdate will update file bases on the given arguments.
// Args can be passed in following format:
// name=<name>   tags=<tag1,tag2,...>   metadata.<key>=<value>
func NewCmdFilesUpdate() *cobra.Command {
	var fileId string

	cmd := &cobra.Command{
		Use:   filesUpdateCmd,
		Short: filesUpdateShort,
		Long:  filesUpdateLong,
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			fdMap := files.FileDetailsMap{}
			fdMetaDataMap := files.FileDetailsMetadataMap{}

			for _, arg := range args {
				if !strings.Contains(arg, "=") {
					return errors.New(fmt.Sprintf(
						"\"%s\" arg error; args need to be passed in key=value format", arg))
				} else {
					argKey := strings.Split(arg, "=")[0]
					argValue := strings.Split(arg, "=")[1]
					if argKey == "tags" {
						tagsArr := strings.Split(argValue, ",")
						fdMap[argKey] = []string(tagsArr)
					} else if strings.Contains(argKey, "metadata.") {
						metadataKey := strings.Split(argKey, ".")[1]
						fdMetaDataMap[metadataKey] = argValue
						fdMap["metadata"] = fdMetaDataMap
					} else {
						fdMap[argKey] = argValue
					}
				}
			}

			respBody, err := files.UpdateFileDetails(fileId, fdMap)
			if err != nil {
				return err
			}

			fd := files.FileDetails{}
			err = json.Unmarshal(respBody, &fd)
			if err != nil {
				return err
			}

			return printFileDetails(fd)
		},
	}
	cmd.Flags().StringVarP(&fileId, filesUpdateFlagFile, filesUpdateFlagFileSh, "", filesUpdateShort)
	cmd.MarkFlagRequired(filesStatFlagFile)
	return cmd
}
