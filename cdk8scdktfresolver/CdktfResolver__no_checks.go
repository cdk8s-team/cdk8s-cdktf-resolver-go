//go:build no_runtime_type_checking

package cdk8scdktfresolver

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdktfResolver) validateResolveParameters(context cdk8s.ResolutionContext) error {
	return nil
}

func validateNewCdktfResolverParameters(props *CdktfResolverProps) error {
	return nil
}

