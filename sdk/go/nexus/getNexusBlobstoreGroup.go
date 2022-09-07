// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNexusBlobstoreGroup(ctx *pulumi.Context, args *LookupNexusBlobstoreGroupArgs, opts ...pulumi.InvokeOption) (*LookupNexusBlobstoreGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNexusBlobstoreGroupResult
	err := ctx.Invoke("nexus:index/getNexusBlobstoreGroup:GetNexusBlobstoreGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetNexusBlobstoreGroup.
type LookupNexusBlobstoreGroupArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by GetNexusBlobstoreGroup.
type LookupNexusBlobstoreGroupResult struct {
	AvailableSpaceInBytes int                               `pulumi:"availableSpaceInBytes"`
	BlobCount             int                               `pulumi:"blobCount"`
	FillPolicy            string                            `pulumi:"fillPolicy"`
	Id                    string                            `pulumi:"id"`
	Members               []string                          `pulumi:"members"`
	Name                  string                            `pulumi:"name"`
	SoftQuotas            []GetNexusBlobstoreGroupSoftQuota `pulumi:"softQuotas"`
	TotalSizeInBytes      int                               `pulumi:"totalSizeInBytes"`
}

func LookupNexusBlobstoreGroupOutput(ctx *pulumi.Context, args LookupNexusBlobstoreGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNexusBlobstoreGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNexusBlobstoreGroupResult, error) {
			args := v.(LookupNexusBlobstoreGroupArgs)
			r, err := LookupNexusBlobstoreGroup(ctx, &args, opts...)
			var s LookupNexusBlobstoreGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNexusBlobstoreGroupResultOutput)
}

// A collection of arguments for invoking GetNexusBlobstoreGroup.
type LookupNexusBlobstoreGroupOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupNexusBlobstoreGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusBlobstoreGroupArgs)(nil)).Elem()
}

// A collection of values returned by GetNexusBlobstoreGroup.
type LookupNexusBlobstoreGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNexusBlobstoreGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusBlobstoreGroupResult)(nil)).Elem()
}

func (o LookupNexusBlobstoreGroupResultOutput) ToLookupNexusBlobstoreGroupResultOutput() LookupNexusBlobstoreGroupResultOutput {
	return o
}

func (o LookupNexusBlobstoreGroupResultOutput) ToLookupNexusBlobstoreGroupResultOutputWithContext(ctx context.Context) LookupNexusBlobstoreGroupResultOutput {
	return o
}

func (o LookupNexusBlobstoreGroupResultOutput) AvailableSpaceInBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreGroupResult) int { return v.AvailableSpaceInBytes }).(pulumi.IntOutput)
}

func (o LookupNexusBlobstoreGroupResultOutput) BlobCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreGroupResult) int { return v.BlobCount }).(pulumi.IntOutput)
}

func (o LookupNexusBlobstoreGroupResultOutput) FillPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreGroupResult) string { return v.FillPolicy }).(pulumi.StringOutput)
}

func (o LookupNexusBlobstoreGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNexusBlobstoreGroupResultOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreGroupResult) []string { return v.Members }).(pulumi.StringArrayOutput)
}

func (o LookupNexusBlobstoreGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNexusBlobstoreGroupResultOutput) SoftQuotas() GetNexusBlobstoreGroupSoftQuotaArrayOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreGroupResult) []GetNexusBlobstoreGroupSoftQuota { return v.SoftQuotas }).(GetNexusBlobstoreGroupSoftQuotaArrayOutput)
}

func (o LookupNexusBlobstoreGroupResultOutput) TotalSizeInBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNexusBlobstoreGroupResult) int { return v.TotalSizeInBytes }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNexusBlobstoreGroupResultOutput{})
}