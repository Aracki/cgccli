package files

import (
	"flag"
	"github.com/aracki/cgccli/api/files"
	"github.com/spf13/viper"
	"testing"
)

var token string

func init() {
	flag.StringVar(&token, "token", "", "token for authorization")
}

var projectTest = "aracki_ivan/belgrade-test-project"
var fileNameTest = "Homo_sapiens_asssembly38.fasta.fai"

// TestProjectsList tests the API
func TestFilesList(t *testing.T) {

	viper.Set("token", token)

	files, err := files.GetFiles(projectTest)
	if err != nil {
		t.Error(err)
	}

	if len(files) < 1 {
		t.Error("no files")
	} else {
		var ok bool
		for _, f := range files {
			if f.Name == fileNameTest {
				ok = true
			}
		}
		if !ok {
			t.Error("there is no file", fileNameTest)
		}
	}
}
