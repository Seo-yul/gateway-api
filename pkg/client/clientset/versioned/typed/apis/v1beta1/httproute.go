/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	apisv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
	applyconfigurationapisv1beta1 "sigs.k8s.io/gateway-api/applyconfiguration/apis/v1beta1"
	scheme "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/scheme"
)

// HTTPRoutesGetter has a method to return a HTTPRouteInterface.
// A group's client should implement this interface.
type HTTPRoutesGetter interface {
	HTTPRoutes(namespace string) HTTPRouteInterface
}

// HTTPRouteInterface has methods to work with HTTPRoute resources.
type HTTPRouteInterface interface {
	Create(ctx context.Context, hTTPRoute *apisv1beta1.HTTPRoute, opts v1.CreateOptions) (*apisv1beta1.HTTPRoute, error)
	Update(ctx context.Context, hTTPRoute *apisv1beta1.HTTPRoute, opts v1.UpdateOptions) (*apisv1beta1.HTTPRoute, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, hTTPRoute *apisv1beta1.HTTPRoute, opts v1.UpdateOptions) (*apisv1beta1.HTTPRoute, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*apisv1beta1.HTTPRoute, error)
	List(ctx context.Context, opts v1.ListOptions) (*apisv1beta1.HTTPRouteList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apisv1beta1.HTTPRoute, err error)
	Apply(ctx context.Context, hTTPRoute *applyconfigurationapisv1beta1.HTTPRouteApplyConfiguration, opts v1.ApplyOptions) (result *apisv1beta1.HTTPRoute, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, hTTPRoute *applyconfigurationapisv1beta1.HTTPRouteApplyConfiguration, opts v1.ApplyOptions) (result *apisv1beta1.HTTPRoute, err error)
	HTTPRouteExpansion
}

// hTTPRoutes implements HTTPRouteInterface
type hTTPRoutes struct {
	*gentype.ClientWithListAndApply[*apisv1beta1.HTTPRoute, *apisv1beta1.HTTPRouteList, *applyconfigurationapisv1beta1.HTTPRouteApplyConfiguration]
}

// newHTTPRoutes returns a HTTPRoutes
func newHTTPRoutes(c *GatewayV1beta1Client, namespace string) *hTTPRoutes {
	return &hTTPRoutes{
		gentype.NewClientWithListAndApply[*apisv1beta1.HTTPRoute, *apisv1beta1.HTTPRouteList, *applyconfigurationapisv1beta1.HTTPRouteApplyConfiguration](
			"httproutes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *apisv1beta1.HTTPRoute { return &apisv1beta1.HTTPRoute{} },
			func() *apisv1beta1.HTTPRouteList { return &apisv1beta1.HTTPRouteList{} },
		),
	}
}
