package files

import (
	"fmt"
	"github.com/aracki/cgccli/api/files"
	"github.com/aracki/cgccli/cmd/util"
)

func printFiles(files []files.File) error {

	w := util.NewTabWriter()
	defer w.Flush()

	_, err := fmt.Fprintln(w, header())
	if err != nil {
		return err
	}
	for _, f := range files {
		_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", f.Id, f.Name, f.Parent, f.Project, f.Type)
		if err != nil {
			return err
		}
	}
	return nil
}

func header() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t", "ID", "NAME", "PARENT", "PROJECT", "TYPE")
}
