package templates

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	operatorsv1alpha1 "github.com/vincent-pli/metal3-operator/pkg/apis/baremetalhost/v1alpha1"
	"sigs.k8s.io/kustomize/v3/k8sdeps/kunstruct"
	"sigs.k8s.io/kustomize/v3/k8sdeps/transformer"
	"sigs.k8s.io/kustomize/v3/k8sdeps/validator"
	"sigs.k8s.io/kustomize/v3/pkg/fs"
	"sigs.k8s.io/kustomize/v3/pkg/loader"
	"sigs.k8s.io/kustomize/v3/pkg/plugins"
	"sigs.k8s.io/kustomize/v3/pkg/resmap"
	"sigs.k8s.io/kustomize/v3/pkg/resource"
	"sigs.k8s.io/kustomize/v3/pkg/target"
)

const TemplatesPathEnvVar = "TEMPLATES_PATH"

var components = []string{
	"baremetal-operator",
	"cluster-api-baremetal",
}

var loadTemplateRendererOnce sync.Once
var templateRenderer *TemplateRenderer

type TemplateRenderer struct {
	templatesPath string
	templates     map[string]resmap.ResMap
}

func GetTemplateRenderer() *TemplateRenderer {
	loadTemplateRendererOnce.Do(func() {
		templatesPath, found := os.LookupEnv(TemplatesPathEnvVar)
		if !found {
			log.Fatalf("TEMPLATES_PATH environment variable is required")
		}
		templateRenderer = &TemplateRenderer{
			templatesPath: templatesPath,
			templates:     map[string]resmap.ResMap{},
		}
	})
	return templateRenderer
}

func (r *TemplateRenderer) GetTemplates(multicloudHub *operatorsv1alpha1.Baremetal) ([]*resource.Resource, error) {
	var err error
	resources := []*resource.Resource{}

	kind := strings.ToLower(multicloudHub.Kind)
	for _, component := range components {
		key := fmt.Sprintf("%s-%s", kind, component)
		resMap, ok := r.templates[key]
		if !ok {
			resMap, err = r.render(path.Join(r.templatesPath, kind, component))
			if err != nil {
				return nil, err
			}
			r.templates[key] = resMap
		}
		resources = append(resources, resMap.Resources()...)

	}

	return resources, err
}

func (r *TemplateRenderer) render(kustomizationPath string) (resmap.ResMap, error) {
	ldr, err := loader.NewLoader(loader.RestrictionRootOnly, validator.NewKustValidator(), kustomizationPath, fs.MakeFsOnDisk())
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := ldr.Cleanup(); err != nil {
			log.Printf("failed to clean up loader, %v", err)
		}
	}()
	pf := transformer.NewFactoryImpl()
	rf := resmap.NewFactory(resource.NewFactory(kunstruct.NewKunstructuredFactoryImpl()), pf)
	pl := plugins.NewLoader(plugins.DefaultPluginConfig(), rf)
	kt, err := target.NewKustTarget(ldr, rf, pf, pl)
	if err != nil {
		return nil, err
	}
	return kt.MakeCustomizedResMap()
}
