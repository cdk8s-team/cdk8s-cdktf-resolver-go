// @cdk8s/cdktf-resolver
package cdk8scdktfresolver

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdk8s/cdktf-resolver.CdktfResolver",
		reflect.TypeOf((*CdktfResolver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
		},
		func() interface{} {
			j := jsiiProxy_CdktfResolver{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sIResolver)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdk8s/cdktf-resolver.CdktfResolverProps",
		reflect.TypeOf((*CdktfResolverProps)(nil)).Elem(),
	)
}
