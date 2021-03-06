// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/cluster-image-registry-operator/pkg/apis/imageregistry/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ImageRegistryLister helps list ImageRegistries.
type ImageRegistryLister interface {
	// List lists all ImageRegistries in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ImageRegistry, err error)
	// Get retrieves the ImageRegistry from the index for a given name.
	Get(name string) (*v1alpha1.ImageRegistry, error)
	ImageRegistryListerExpansion
}

// imageRegistryLister implements the ImageRegistryLister interface.
type imageRegistryLister struct {
	indexer cache.Indexer
}

// NewImageRegistryLister returns a new ImageRegistryLister.
func NewImageRegistryLister(indexer cache.Indexer) ImageRegistryLister {
	return &imageRegistryLister{indexer: indexer}
}

// List lists all ImageRegistries in the indexer.
func (s *imageRegistryLister) List(selector labels.Selector) (ret []*v1alpha1.ImageRegistry, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ImageRegistry))
	})
	return ret, err
}

// Get retrieves the ImageRegistry from the index for a given name.
func (s *imageRegistryLister) Get(name string) (*v1alpha1.ImageRegistry, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("imageregistry"), name)
	}
	return obj.(*v1alpha1.ImageRegistry), nil
}
