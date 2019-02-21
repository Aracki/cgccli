package projects

import (
	"fmt"
	"github.com/aracki/cgccli/api"
	"github.com/aracki/cgccli/cmd/util"
)

func printProjects(projects []api.Project) {

	w := util.NewTabWriter()
	defer w.Flush()

	fmt.Fprintln(w, header())
	for _, p := range projects {
		fmt.Fprintf(w, "%s\t%s\t\n", p.Id, p.Name)
	}
}

func header() string {
	return fmt.Sprintf("%s\t%s\t", "ID", "NAME")
}
