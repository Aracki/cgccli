package files

import (
	"fmt"
	"github.com/aracki/cgccli/api/files"
	"github.com/aracki/cgccli/cmd/util"
	"reflect"
)

func filesHeader() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t",
		"ID", "NAME", "PARENT", "PROJECT", "TYPE")
}

func printFiles(files []files.File) error {

	w := util.NewTabWriter()
	defer w.Flush()

	_, err := fmt.Fprintln(w, filesHeader())
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

func fileDetailsHeader() string {
	return fmt.Sprintf("%s\t%s\t",
		"KEY", "VALUE")
}

func printFileDetails(fDetails files.FileDetails) error {

	w := util.NewTabWriter()
	defer w.Flush()

	_, err := fmt.Fprintln(w, fileDetailsHeader())
	if err != nil {
		return err
	}

	v := reflect.ValueOf(fDetails)

	for i := 0; i < v.NumField(); i++ {
		_, err := fmt.Fprintf(w, "%s\t%v\t\n", v.Type().Field(i).Name, v.Field(i).Interface())
		if err != nil {
			return err
		}
	}
	return nil
}
