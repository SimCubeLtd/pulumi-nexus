// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNexusBlobstore(ctx *pulumi.Context, args *LookupNexusBlobstoreArgs, opts ...pulumi.InvokeOption) (*LookupNexusBlobstoreResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNexusBlobstoreResult
	err := ctx.Invoke("nexus:index/getNexusBlobstore:GetNexusBlobstore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetNexusBlobstore.
type LookupNexusBlobstoreArgs struct {
	BucketConfiguration *GetNexusBlobstoreBucketConfiguration `pulumi:"bucketConfiguration"`
	Name                string                                `pulumi:"name"`
	Path                *string                               `pulumi:"path"`
	SoftQuota           *GetNexusBlobstoreSoftQuota           `pulumi:"softQuota"`
	Type                *string                               `pulumi:"type"`
}

// A collection of values returned by GetNexusBlobstore.
type LookupNexusBlobstoreResult struct {
	AvailableSpaceInBytes int                                   `pulumi:"availableSpaceInBytes"`
	BlobCount             int                                   `pulumi:"blobCount"`
	BucketConfiguration   *GetNexusBlobstoreBucketConfiguration `pulumi:"bucketConfiguration"`
	Id                    string                                `pulumi:"id"`
	Name                  string                                `pulumi:"name"`
	Path                  *string                               `pulumi:"path"`
	SoftQuota             *GetNexusBlobstoreSoftQuota           `pulumi:"softQuota"`
	TotalSizeInBytes      int                                   `pulumi:"totalSizeInBytes"`
	Type                  *string                               `pulumi:"type"`
}

func LookupNexusBlobstoreOutput(ctx *pulumi.Context, args LookupNexusBlobstoreOutputArgs, opts ...pulumi.InvokeOption) LookupNexusBlobstoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNexusBlobstoreResult, error) {
			args := v.(LookupNexusBlobstoreArgs)
			r, err := LookupNexusBlobstore(ctx, &args, opts...)
			var s LookupNexusBlobstoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNexusBlobstoreResultOutput)
}

// A collection of arguments for invoking GetNexusBlobstore.
type LookupNexusBlobstoreOutputArgs struct {
	BucketConfiguration GetNexusBlobstoreBucketConfigurationPtrInput `pulumi:"bucketConfiguration"`
	Name                pulumi.StringInput                           `pulumi:"name"`
	Path                pulumi.StringPtrInput                        `pulumi:"path"`
	SoftQuota           GetNexusBlobstoreSoftQuotaPtrInput           `pulumi:"softQuota"`
	Type                pulumi.StringPtrInput                        `pulumi:"type"`
}

func (LookupNexusBlobstoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusBlobstoreArgs)(nil)).Elem()
}

// A collection of values returned by GetNexusBlobstore.
type LookupNexusBlobstoreResultOutput struct{ *pulumi.OutputState }

func (LookupNexusBlobstoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusBlobstoreResult)(nil)).Elem()
}

func (o LookupNexusBlobstoreResultOutput) ToLookupNexusBlobstoreResultOutput() LookupNexusBlobstoreResultOutput {
	return o
}

func (o LookupNexusBlobstoreResultOutput) ToLookupNexusBlobstoreResultOutputWithContext(ctx context.Context) LookupNexusBlobstoreResultOutput {
	return o
}

func (o LookupNexusBlobstoreResultOutput) AvailableSpaceInBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreResult) int { return v.AvailableSpaceInBytes }).(pulumi.IntOutput)
}

func (o LookupNexusBlobstoreResultOutput) BlobCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreResult) int { return v.BlobCount }).(pulumi.IntOutput)
}

func (o LookupNexusBlobstoreResultOutput) BucketConfiguration() GetNexusBlobstoreBucketConfigurationPtrOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreResult) *GetNexusBlobstoreBucketConfiguration { return v.BucketConfiguration }).(GetNexusBlobstoreBucketConfigurationPtrOutput)
}

func (o LookupNexusBlobstoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNexusBlobstoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNexusBlobstoreResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o LookupNexusBlobstoreResultOutput) SoftQuota() GetNexusBlobstoreSoftQuotaPtrOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreResult) *GetNexusBlobstoreSoftQuota { return v.SoftQuota }).(GetNexusBlobstoreSoftQuotaPtrOutput)
}

func (o LookupNexusBlobstoreResultOutput) TotalSizeInBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreResult) int { return v.TotalSizeInBytes }).(pulumi.IntOutput)
}

func (o LookupNexusBlobstoreResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNexusBlobstoreResultOutput{})
}
