package internal_test

import (
	"os"
	"testing"

	lib "github.com/kube-champ/terraform-runner/internal"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTerraform(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Terraform Suite")
}

func createFile(filePath string) *os.File {
	file, _ := os.Create(filePath)

	return file
}

func writeFile(filePath string, content string) {
	file := createFile(filePath)

	file.WriteString(content)
}

func mkdir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0700)
	}
}

const (
	workDir    string = "/tmp/tftest"
	secretName string = "my-secret"
)

var _ = BeforeSuite(func() {

	os.Setenv("TERRAFORM_WORKING_DIR", workDir)
	os.Setenv("TERRAFORM_VERSION", "1.0.2")
	os.Setenv("TERRAFORM_WORKSPACE", "default")
	os.Setenv("OUTPUT_SECRET_NAME", secretName)

	lib.LoadEnv()
})
