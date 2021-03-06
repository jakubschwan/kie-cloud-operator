// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/openshift/api/network/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterNetworkLister helps list ClusterNetworks.
type ClusterNetworkLister interface {
	// List lists all ClusterNetworks in the indexer.
	List(selector labels.Selector) (ret []*v1.ClusterNetwork, err error)
	// Get retrieves the ClusterNetwork from the index for a given name.
	Get(name string) (*v1.ClusterNetwork, error)
	ClusterNetworkListerExpansion
}

// clusterNetworkLister implements the ClusterNetworkLister interface.
type clusterNetworkLister struct {
	indexer cache.Indexer
}

// NewClusterNetworkLister returns a new ClusterNetworkLister.
func NewClusterNetworkLister(indexer cache.Indexer) ClusterNetworkLister {
	return &clusterNetworkLister{indexer: indexer}
}

// List lists all ClusterNetworks in the indexer.
func (s *clusterNetworkLister) List(selector labels.Selector) (ret []*v1.ClusterNetwork, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterNetwork))
	})
	return ret, err
}

// Get retrieves the ClusterNetwork from the index for a given name.
func (s *clusterNetworkLister) Get(name string) (*v1.ClusterNetwork, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusternetwork"), name)
	}
	return obj.(*v1.ClusterNetwork), nil
}
