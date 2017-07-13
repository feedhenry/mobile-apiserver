/*
Copyright 2017 The Kubernetes Authors.

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

package mobile

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kapi "k8s.io/client-go/pkg/api"
)

// MobileAppList is a list of MobileApp objects.
type MobileAppList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []MobileApp
}

// +genclient=true

// MobileApp represents a service instance provision request,
// possibly fullfilled.
type MobileApp struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   MobileAppSpec
	Status MobileAppStatus
}

// MobileAppSpec defines the requested MobileApp
type MobileAppSpec struct {
	Credential string
}

// MobileAppStatus defines the current state of the MobileApp
type MobileAppStatus struct {
	Conditions []MobileAppCondition
}

// MobileAppCondition contains condition information for a
// MobileApp.
type MobileAppCondition struct {
	// Type of the condition, currently Ready or InstantiateFailure.
	Type MobileAppConditionType
	// Status of the condition, one of True, False or Unknown.
	Status kapi.ConditionStatus
	// LastTransitionTime is the last time a condition status transitioned from
	// one state to another.
	LastTransitionTime metav1.Time
	// Reason is a brief machine readable explanation for the condition's last
	// transition.
	Reason string
	// Message is a human readable description of the details of the last
	// transition, complementing reason.
	Message string
}

// MobileAppConditionType is the type of condition pertaining to a
// MobileApp.
type MobileAppConditionType string

const (
	// MobileAppReady indicates the service instance is Ready for use
	// (provision was successful)
	MobileAppReady MobileAppConditionType = "Ready"

	// MobileAppInstantiateFailed indicates the provision request failed.
	MobileAppFailed MobileAppConditionType = "Failure"

	// TypePackage is the name of the package that defines the resource types
	// used by this broker.
	TypePackage = "github.com/feedhenry/mobile-apiserver/pkg/apis/mobile"

	// GroupName is the name of the api group used for resources created/managed
	// by this broker.
	GroupName = "sdkbroker.broker.k8s.io"

	// MobileAppsResource is the name of the resource used to represent
	// provision requests(possibly fulfilled) for service instances
	MobileAppsResource = "MobileApps"

	// MobileAppResource is the name of the resource used to represent
	// provision requests(possibly fulfilled) for service instances
	MobileAppResource = "MobileApp"

	// BrokerAPIPrefix is the route prefix for the open service broker api
	// endpoints (e.g. https://yourhost.com/broker/sdkbroker.broker.io/v2/catalog)
	BrokerAPIPrefix = "/mobile/mobile.srv.io"

	// Namespace is the namespace the broker will be deployed in and
	// under which it will create any resources
	Namespace = "mobile"
)
