package projects

import (
	"fmt"
	"github.com/aracki/cgccli/api"
	"os"
)

func printProjects(projects []api.Project) {

	w := os.Stdout

	for i := 0; i < len(projects); i++ {
		fmt.Fprintf(w, "id: %s \n", projects[i].Id)
		fmt.Fprintf(w, "name: %s \n\n", projects[i].Name)
	}
}
