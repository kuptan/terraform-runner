package internal

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers function", func() {
	BeforeEach(func() {
		dirs := []string{
			"/tmp/tfvars/common-config",
			"/tmp/tfvars/data-config",
		}

		for _, d := range dirs {
			file := fmt.Sprintf("%s/data.tfvars", d)

			os.MkdirAll(d, os.ModePerm)
			os.WriteFile(file, []byte("length = 10"), 0644)
		}
	})

	Context("helper functions", func() {
		It("array should contain a string", func() {
			arr := []string{"abc", "deg"}

			Expect(arrayContains(arr, "abc")).To(BeTrue())
			Expect(arrayContains(arr, "123")).To(BeFalse())
		})

		It("should be a valid tfvar extension", func() {
			Expect(varFileExtensionAllowed("common.tfvars")).To(BeTrue())
			Expect(varFileExtensionAllowed("common.tf")).To(BeTrue())
			Expect(varFileExtensionAllowed("common.json")).To(BeTrue())
			Expect(varFileExtensionAllowed("common.terraform")).To(BeFalse())
		})
	})

	Context("listing files", func() {
		It("should check if file exist", func() {
			Expect(fileExists("/tmp/tfvars/common-config/data.tfvars")).To(BeTrue())
		})

		It("should get all var files in nested directories", func() {
			files, err := getTfVarFilesPaths("/tmp/tfvars")

			Expect(err).ToNot(HaveOccurred())
			Expect(files).To(HaveLen(2))

			Expect(files[0]).To(Equal("/tmp/tfvars/common-config/data.tfvars"))
			Expect(files[1]).To(Equal("/tmp/tfvars/data-config/data.tfvars"))
		})
	})
})
