package controller

import (
	"github.com/StephenGrider/docker-react/pkg/controller/react"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, react.Add)
}
