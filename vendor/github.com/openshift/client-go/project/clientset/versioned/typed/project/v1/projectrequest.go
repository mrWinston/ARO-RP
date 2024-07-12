// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"

	projectv1 "github.com/openshift/api/project/v1"
	v1 "github.com/openshift/client-go/project/applyconfigurations/project/v1"
	scheme "github.com/openshift/client-go/project/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	rest "k8s.io/client-go/rest"
)

// ProjectRequestsGetter has a method to return a ProjectRequestInterface.
// A group's client should implement this interface.
type ProjectRequestsGetter interface {
	ProjectRequests() ProjectRequestInterface
}

// ProjectRequestInterface has methods to work with ProjectRequest resources.
type ProjectRequestInterface interface {
	Apply(ctx context.Context, projectRequest *v1.ProjectRequestApplyConfiguration, opts metav1.ApplyOptions) (result *projectv1.ProjectRequest, err error)
	Create(ctx context.Context, projectRequest *projectv1.ProjectRequest, opts metav1.CreateOptions) (*projectv1.Project, error)

	ProjectRequestExpansion
}

// projectRequests implements ProjectRequestInterface
type projectRequests struct {
	client rest.Interface
}

// newProjectRequests returns a ProjectRequests
func newProjectRequests(c *ProjectV1Client) *projectRequests {
	return &projectRequests{
		client: c.RESTClient(),
	}
}

// Apply takes the given apply declarative configuration, applies it and returns the applied projectRequest.
func (c *projectRequests) Apply(ctx context.Context, projectRequest *v1.ProjectRequestApplyConfiguration, opts metav1.ApplyOptions) (result *projectv1.ProjectRequest, err error) {
	if projectRequest == nil {
		return nil, fmt.Errorf("projectRequest provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(projectRequest)
	if err != nil {
		return nil, err
	}
	name := projectRequest.Name
	if name == nil {
		return nil, fmt.Errorf("projectRequest.Name must be provided to Apply")
	}
	result = &projectv1.ProjectRequest{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("projectrequests").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Create takes the representation of a projectRequest and creates it.  Returns the server's representation of the project, and an error, if there is any.
func (c *projectRequests) Create(ctx context.Context, projectRequest *projectv1.ProjectRequest, opts metav1.CreateOptions) (result *projectv1.Project, err error) {
	result = &projectv1.Project{}
	err = c.client.Post().
		Resource("projectrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(projectRequest).
		Do(ctx).
		Into(result)
	return
}
