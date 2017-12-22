package main

import (
	"encoding/json"

	"github.com/eawsy/aws-cloudformation-go-customres/service/cloudformation/customres"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/cloudformationevt"
)

var (
	// Handle is the Lambda's entrypoint.
	Handle customres.LambdaHandler
)

func init() {
	customres.Register("ExampleResource", new(ExampleResource))
	Handle = customres.HandleLambda
}

// Happy IDE means happy developer.
func main() {
}

// ExampleResource represents a simple, custom resource.
type ExampleResource struct {
	ExampleThing *string `json:",omitempty"`
}

// Create is invoked when the resource is created.
func (r *ExampleResource) Create(evt *cloudformationevt.Event, ctx *runtime.Context) (string, interface{}, error) {
	evt.PhysicalResourceID = customres.NewPhysicalResourceID(evt)
	return r.Update(evt, ctx)
}

var (
	defaultExampleThing = "THIS IS THE DEFAULT!"
)

// Update is invoked when the resource is updated.
func (r *ExampleResource) Update(evt *cloudformationevt.Event, ctx *runtime.Context) (string, interface{}, error) {
	if err := json.Unmarshal(evt.ResourceProperties, r); err != nil {
		return "", r, err
	}

	if r.ExampleThing == nil {
		r.ExampleThing = &defaultExampleThing
	}

	return evt.PhysicalResourceID, r, nil
}

// Delete is invoked when the resource is deleted.
func (r *ExampleResource) Delete(*cloudformationevt.Event, *runtime.Context) error {
	return nil
}
