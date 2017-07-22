/*
Copyright 2016 The Kubernetes Authors.

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

package serviceinstance

import (
	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"

	"github.com/feedhenry/mobile-apiserver/pkg/apis/mobile"
)

// NewREST returns a RESTStorage object that will work against broker API resources.
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) rest.Storage {
	strategy := NewStrategy(scheme)

	store := &registry.Store{
		Copier:      scheme,
		NewFunc:     func() runtime.Object { return &mobile.MobileApp{} },
		NewListFunc: func() runtime.Object { return &mobile.MobileAppList{} },
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*mobile.MobileApp).Name, nil
		},
		PredicateFunc:     MatchMobileApp,
		QualifiedResource: mobile.Resource(mobile.MobileAppsResource),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err)
	}
	return store
}
