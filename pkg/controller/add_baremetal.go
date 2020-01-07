package controller

import (
	"github.com/vincent-pli/metal3-operator/pkg/controller/baremetal"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, baremetal.Add)
}
