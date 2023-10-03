//go:build !no_runtime_type_checking

package cdk8scdktfresolver

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func (c *jsiiProxy_CdktfResolver) validateResolveParameters(context cdk8s.ResolutionContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func validateNewCdktfResolverParameters(props *CdktfResolverProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

