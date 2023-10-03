package cdk8scdktfresolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-cdktf-resolver-go/cdk8scdktfresolver/jsii"

	"github.com/cdk8s-team/cdk8s-cdktf-resolver-go/cdk8scdktfresolver/internal"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type CdktfResolver interface {
	cdk8s.IResolver
	// This function is invoked on every property during cdk8s synthesis.
	//
	// To replace a value, implementations must invoke `context.replaceValue`.
	Resolve(context cdk8s.ResolutionContext)
}

// The jsii proxy struct for CdktfResolver
type jsiiProxy_CdktfResolver struct {
	internal.Type__cdk8sIResolver
}

func NewCdktfResolver(props *CdktfResolverProps) CdktfResolver {
	_init_.Initialize()

	if err := validateNewCdktfResolverParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdktfResolver{}

	_jsii_.Create(
		"@cdk8s/cdktf-resolver.CdktfResolver",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCdktfResolver_Override(c CdktfResolver, props *CdktfResolverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdk8s/cdktf-resolver.CdktfResolver",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CdktfResolver) Resolve(context cdk8s.ResolutionContext) {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"resolve",
		[]interface{}{context},
	)
}

