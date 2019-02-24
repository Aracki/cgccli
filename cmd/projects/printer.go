package projects

import (
	"fmt"
	"github.com/aracki/cgccli/api/projects"
	"github.com/aracki/cgccli/cmd/util"
)

// printProjects prints the projects array.
func printProjects(projects []projects.Project) error {

	w := util.NewTabWriter()
	defer w.Flush()

	_, err := fmt.Fprintln(w, header())
	if err != nil {
		return err
	}
	for _, p := range projects {
		_, err := fmt.Fprintf(w, "%s\t%s\t\n", p.Id, p.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func header() string {
	return fmt.Sprintf("%s\t%s\t", "ID", "NAME")
}
