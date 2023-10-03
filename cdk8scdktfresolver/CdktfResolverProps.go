package cdk8scdktfresolver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdktfResolverProps struct {
	// The CDKTF App instance in which the outputs are deinfed in.
	App cdktf.App `field:"required" json:"app" yaml:"app"`
}

