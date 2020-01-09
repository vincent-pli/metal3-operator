package render

import (
	operatorsv1alpha1 "github.com/vincent-pli/metal3-operator/pkg/apis/baremetalhost/v1alpha1"
	"github.com/vincent-pli/metal3-operator/pkg/render/templates"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/kustomize/v3/pkg/resource"
)

type renderFn func(*resource.Resource) (*unstructured.Unstructured, error)

type Renderer struct {
	cr *operatorsv1alpha1.Baremetal
}

func NewRenderer(baremetal *operatorsv1alpha1.Baremetal) *Renderer {
	renderer := &Renderer{cr: baremetal}
	return renderer
}

func (r *Renderer) Render() ([]*unstructured.Unstructured, error) {
	templates, err := templates.GetTemplateRenderer().GetTemplates(r.cr)
	if err != nil {
		return nil, err
	}

	uobjs := []*unstructured.Unstructured{}
	for _, template := range templates {
		uobjs = append(uobjs, &unstructured.Unstructured{Object: template.Map()})
	}
	return uobjs, nil
}
