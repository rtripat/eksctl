package kubernetes

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	apiextensionsv1beta1.AddToScheme(scheme.Scheme)
}
