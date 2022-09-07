// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNexusBlobstoreS3(ctx *pulumi.Context, args *LookupNexusBlobstoreS3Args, opts ...pulumi.InvokeOption) (*LookupNexusBlobstoreS3Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNexusBlobstoreS3Result
	err := ctx.Invoke("nexus:index/getNexusBlobstoreS3:GetNexusBlobstoreS3", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetNexusBlobstoreS3.
type LookupNexusBlobstoreS3Args struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by GetNexusBlobstoreS3.
type LookupNexusBlobstoreS3Result struct {
	BlobCount            int                                      `pulumi:"blobCount"`
	BucketConfigurations []GetNexusBlobstoreS3BucketConfiguration `pulumi:"bucketConfigurations"`
	Id                   string                                   `pulumi:"id"`
	Name                 string                                   `pulumi:"name"`
	SoftQuotas           []GetNexusBlobstoreS3SoftQuota           `pulumi:"softQuotas"`
	TotalSizeInBytes     int                                      `pulumi:"totalSizeInBytes"`
}

func LookupNexusBlobstoreS3Output(ctx *pulumi.Context, args LookupNexusBlobstoreS3OutputArgs, opts ...pulumi.InvokeOption) LookupNexusBlobstoreS3ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNexusBlobstoreS3Result, error) {
			args := v.(LookupNexusBlobstoreS3Args)
			r, err := LookupNexusBlobstoreS3(ctx, &args, opts...)
			var s LookupNexusBlobstoreS3Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNexusBlobstoreS3ResultOutput)
}

// A collection of arguments for invoking GetNexusBlobstoreS3.
type LookupNexusBlobstoreS3OutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupNexusBlobstoreS3OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusBlobstoreS3Args)(nil)).Elem()
}

// A collection of values returned by GetNexusBlobstoreS3.
type LookupNexusBlobstoreS3ResultOutput struct{ *pulumi.OutputState }

func (LookupNexusBlobstoreS3ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusBlobstoreS3Result)(nil)).Elem()
}

func (o LookupNexusBlobstoreS3ResultOutput) ToLookupNexusBlobstoreS3ResultOutput() LookupNexusBlobstoreS3ResultOutput {
	return o
}

func (o LookupNexusBlobstoreS3ResultOutput) ToLookupNexusBlobstoreS3ResultOutputWithContext(ctx context.Context) LookupNexusBlobstoreS3ResultOutput {
	return o
}

func (o LookupNexusBlobstoreS3ResultOutput) BlobCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreS3Result) int { return v.BlobCount }).(pulumi.IntOutput)
}

func (o LookupNexusBlobstoreS3ResultOutput) BucketConfigurations() GetNexusBlobstoreS3BucketConfigurationArrayOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreS3Result) []GetNexusBlobstoreS3BucketConfiguration {
		return v.BucketConfigurations
	}).(GetNexusBlobstoreS3BucketConfigurationArrayOutput)
}

func (o LookupNexusBlobstoreS3ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreS3Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNexusBlobstoreS3ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreS3Result) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNexusBlobstoreS3ResultOutput) SoftQuotas() GetNexusBlobstoreS3SoftQuotaArrayOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreS3Result) []GetNexusBlobstoreS3SoftQuota { return v.SoftQuotas }).(GetNexusBlobstoreS3SoftQuotaArrayOutput)
}

func (o LookupNexusBlobstoreS3ResultOutput) TotalSizeInBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreS3Result) int { return v.TotalSizeInBytes }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNexusBlobstoreS3ResultOutput{})
}