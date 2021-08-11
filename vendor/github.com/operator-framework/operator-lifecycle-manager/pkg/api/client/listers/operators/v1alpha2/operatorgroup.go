/*
Copyright 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OperatorGroupLister helps list OperatorGroups.
type OperatorGroupLister interface {
	// List lists all OperatorGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.OperatorGroup, err error)
	// OperatorGroups returns an object that can list and get OperatorGroups.
	OperatorGroups(namespace string) OperatorGroupNamespaceLister
	OperatorGroupListerExpansion
}

// operatorGroupLister implements the OperatorGroupLister interface.
type operatorGroupLister struct {
	indexer cache.Indexer
}

// NewOperatorGroupLister returns a new OperatorGroupLister.
func NewOperatorGroupLister(indexer cache.Indexer) OperatorGroupLister {
	return &operatorGroupLister{indexer: indexer}
}

// List lists all OperatorGroups in the indexer.
func (s *operatorGroupLister) List(selector labels.Selector) (ret []*v1alpha2.OperatorGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.OperatorGroup))
	})
	return ret, err
}

// OperatorGroups returns an object that can list and get OperatorGroups.
func (s *operatorGroupLister) OperatorGroups(namespace string) OperatorGroupNamespaceLister {
	return operatorGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OperatorGroupNamespaceLister helps list and get OperatorGroups.
type OperatorGroupNamespaceLister interface {
	// List lists all OperatorGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.OperatorGroup, err error)
	// Get retrieves the OperatorGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.OperatorGroup, error)
	OperatorGroupNamespaceListerExpansion
}

// operatorGroupNamespaceLister implements the OperatorGroupNamespaceLister
// interface.
type operatorGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OperatorGroups in the indexer for a given namespace.
func (s operatorGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.OperatorGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.OperatorGroup))
	})
	return ret, err
}

// Get retrieves the OperatorGroup from the indexer for a given namespace and name.
func (s operatorGroupNamespaceLister) Get(name string) (*v1alpha2.OperatorGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("operatorgroup"), name)
	}
	return obj.(*v1alpha2.OperatorGroup), nil
}
