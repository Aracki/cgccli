package files

import (
	"encoding/json"
	"fmt"
	"github.com/aracki/cgccli/api/files"
	"github.com/aracki/cgccli/cmd/util"
	"io"
	"reflect"
)

var (
	filesHeader = fmt.Sprintf("%s\t%s\t%s\t%s\t",
		"ID", "NAME", "PARENT", "TYPE")
	fileDetailsHeader = fmt.Sprintf("%s\t%s\t",
		"KEY", "VALUE")
)

func printFiles(files []files.File) error {

	w := util.NewTabWriter()
	defer w.Flush()

	_, err := fmt.Fprintln(w, filesHeader)
	if err != nil {
		return err
	}
	for _, f := range files {
		_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", f.Id, f.Name, f.Parent, f.Type)
		if err != nil {
			return err
		}
	}
	return nil
}

func printFileDetails(fDetails files.FileDetails) error {

	w := util.NewTabWriter()
	defer w.Flush()

	_, err := fmt.Fprintln(w, fileDetailsHeader)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(fDetails)

	for i := 0; i < v.NumField(); i++ {

		// all cases instead of Metadata map field
		if v.Type().Field(i).Type.Kind() != reflect.Map {
			_, err := fmt.Fprintf(w, "%s\t%v\t\n", v.Type().Field(i).Name, v.Field(i).Interface())
			if err != nil {
				return err
			}
		} else {
			err := prettyPrintMetadata(w, v.Field(i).Interface())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// prettyPrintMetadata will print Metadata map[string]string on the bottom as a json.
func prettyPrintMetadata(w io.Writer, v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "\nMETADATA\n %s", string(b))
	if err != nil {
		return err
	}
	return nil
}
