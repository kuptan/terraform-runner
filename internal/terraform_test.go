package internal_test

import (
	"fmt"

	lib "github.com/kube-champ/terraform-runner/internal"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("terraform helper", func() {
	var tf *lib.TerraformRunner

	BeforeEach(func() {
		tfmodule := `
			variable "length" {
				type = number
				default = 16
			}
	
			resource "random_string" "random" {
				length           = var.length
				special          = true
				override_special = "/@£$"
			}
	
			output "result" {
				value = random_string.random.result
			}
		`

		mkdir(workDir)
		writeFile(fmt.Sprintf("%s/main.tf", workDir), tfmodule)
	})

	When("terraform", func() {
		Context("installation & setup", func() {
			It("should install and setup terraform binary", func() {
				runner, err := lib.Setup()
				tf = runner

				Expect(err).To(BeNil())
			})
		})

		Context("initializing a module", func() {
			It("should initialize successfully", func() {
				err := tf.Init()

				Expect(err).To(BeNil())
			})

			It("should change to workspace 'dev'", func() {
				err := tf.SelectWorkspace("dev")

				Expect(err).To(BeNil())
			})
		})

		Context("plan", func() {
			It("should run a plan command", func() {
				err := tf.Plan()

				Expect(err).To(BeNil())
			})
		})

		Context("apply", func() {
			It("should run apply command", func() {
				err := tf.Apply()

				Expect(err).To(BeNil())
			})
		})

		Context("destroy", func() {
			It("should run destroy command", func() {
				err := tf.Destroy()

				Expect(err).To(BeNil())
			})
		})
	})
})
