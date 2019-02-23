// Package files implements functions
// that write api object fields with initialized tab writer.
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

// printFiles prints the files array.
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

// printFileDetails prints all the fields of a particular file.
// It prints all the fields in default fashion, except maps that prints as a json.
func printFileDetails(fDetails files.FileDetails) error {

	w := util.NewTabWriter()
	defer w.Flush()

	_, err := fmt.Fprintln(w, fileDetailsHeader)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(fDetails)

	for i := 0; i < v.NumField(); i++ {

		switch v.Type().Field(i).Type.Kind() {
		case reflect.Map:
			err := prettyPrintMetadata(w, v.Field(i).Interface())
			if err != nil {
				return err
			}
		default:
			_, err := fmt.Fprintf(w, "%s\t%v\t\n", v.Type().Field(i).Name, v.Field(i).Interface())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// prettyPrintMetadata prints metadata map as a json.
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
