package projects

import (
	"flag"
	"github.com/aracki/cgccli/api/projects"
	"github.com/spf13/viper"
	"testing"
)

var token string

func init() {
	flag.StringVar(&token, "token", "", "token for authorization")
}

var projectTest = "Belgrade Test Project"

// TestProjectsList tests the API
func TestProjectsList(t *testing.T) {

	viper.Set("token", token)

	projects, err := projects.GetProjects()
	if err != nil {
		t.Error(err)
	}

	if len(projects) < 1 {
		t.Error("no projects")
	} else {
		var ok bool
		for _, p := range projects {
			if p.Name == projectTest {
				ok = true
			}
		}
		if !ok {
			t.Error("there is no ", projectTest)
		}
	}
}
