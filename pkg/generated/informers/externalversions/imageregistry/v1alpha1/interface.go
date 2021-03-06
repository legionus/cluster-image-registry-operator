// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openshift/cluster-image-registry-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ImageRegistries returns a ImageRegistryInformer.
	ImageRegistries() ImageRegistryInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ImageRegistries returns a ImageRegistryInformer.
func (v *version) ImageRegistries() ImageRegistryInformer {
	return &imageRegistryInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
