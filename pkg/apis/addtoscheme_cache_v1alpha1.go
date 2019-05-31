package apis

import (
	"github.com/ssup2/example-k8s-operator-memcached/pkg/apis/cache/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
