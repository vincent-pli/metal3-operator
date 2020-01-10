package baremetal

import (
	"github.com/vincent-pli/metal3-operator/pkg/extension/imagereplacement"
)

func init() {
	activities = append(activities, imagereplacement.Configure)
}
