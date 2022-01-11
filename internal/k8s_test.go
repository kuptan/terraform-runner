package internal_test

import (
	"context"
	"reflect"

	"github.com/kube-champ/terraform-runner/internal"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

var _ = Describe("Kubernetes Helper", func() {
	var secret *corev1.Secret

	BeforeEach(func() {
		internal.ClientSet = fake.NewSimpleClientset()

		s := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      secretName,
				Namespace: "default",
			},
			Data: map[string][]byte{
				"key": []byte("ajwdjaw"),
			},
		}

		created, _ := internal.ClientSet.CoreV1().Secrets("default").Create(context.Background(), s, metav1.CreateOptions{})

		// Expect(err).ToNot(HaveOccurred())

		secret = created
	})

	Context("secret operations", func() {
		It("should update the data of an existing secret", func() {
			outputs := map[string][]byte{
				"new_key": []byte("this_is_a_test"),
			}

			internal.UpdateSecretWithOutputs(outputs)

			actual, err := internal.ClientSet.CoreV1().Secrets("default").Get(context.Background(), secret.Name, metav1.GetOptions{})

			Expect(err).ToNot(HaveOccurred())
			Expect(reflect.DeepEqual(actual.Data, outputs)).To(BeTrue())
		})
	})
})
