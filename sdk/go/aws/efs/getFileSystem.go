// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides information about an Elastic File System (EFS).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/efs_file_system.html.markdown.
func LookupFileSystem(ctx *pulumi.Context, args *GetFileSystemArgs, opts ...pulumi.InvokeOption) (*GetFileSystemResult, error) {
	var rv GetFileSystemResult
	err := ctx.Invoke("aws:efs/getFileSystem:getFileSystem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFileSystem.
type GetFileSystemArgs struct {
	// Restricts the list to the file system with this creation token.
	CreationToken *string `pulumi:"creationToken"`
	// The ID that identifies the file system (e.g. fs-ccfc0d65).
	FileSystemId *string `pulumi:"fileSystemId"`
	Tags *map[string]string `pulumi:"tags"`
}

// A collection of values returned by getFileSystem.
type GetFileSystemResult struct {
	// Amazon Resource Name of the file system.
	Arn string `pulumi:"arn"`
	CreationToken string `pulumi:"creationToken"`
	// The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	DnsName string `pulumi:"dnsName"`
	// Whether EFS is encrypted.
	Encrypted bool `pulumi:"encrypted"`
	FileSystemId string `pulumi:"fileSystemId"`
	// The ARN for the KMS encryption key.
	KmsKeyId string `pulumi:"kmsKeyId"`
	// The PerformanceMode of the file system.
	PerformanceMode string `pulumi:"performanceMode"`
	// The list of tags assigned to the file system.
	Tags map[string]string `pulumi:"tags"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
