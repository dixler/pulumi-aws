// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package qldb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to fetch information about a Quantum Ledger Database.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/qldb_ledger.html.markdown.
func LookupLedger(ctx *pulumi.Context, args *GetLedgerArgs, opts ...pulumi.InvokeOption) (*GetLedgerResult, error) {
	var rv GetLedgerResult
	err := ctx.Invoke("aws:qldb/getLedger:getLedger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLedger.
type GetLedgerArgs struct {
	// The friendly name of the ledger to match.
	Name string `pulumi:"name"`
}

// A collection of values returned by getLedger.
type GetLedgerResult struct {
	// Amazon Resource Name (ARN) of the ledger.
	Arn string `pulumi:"arn"`
	// Deletion protection on the QLDB Ledger instance. Set to `true` by default. 
	DeletionProtection bool `pulumi:"deletionProtection"`
	Name string `pulumi:"name"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
