package managedresource

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	AnnotationSuperStream           = "rabbitmq.com/super-stream"
	AnnotationSuperStreamRoutingKey = "rabbitmq.com/super-stream-routing-key"
)

type Builder struct {
	ObjectOwner metav1.Object
	Scheme      *runtime.Scheme
}

type ResourceBuilder interface {
	Build() (client.Object, error)
	Update(client.Object) error
	ResourceType() string
}

func (builder Builder) GenerateChildResourceName(suffix string) string {
	return builder.ObjectOwner.GetName() + suffix
}
