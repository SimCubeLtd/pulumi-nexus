// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNexusPrivileges(ctx *pulumi.Context, args *GetNexusPrivilegesArgs, opts ...pulumi.InvokeOption) (*GetNexusPrivilegesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetNexusPrivilegesResult
	err := ctx.Invoke("nexus:index/getNexusPrivileges:GetNexusPrivileges", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetNexusPrivileges.
type GetNexusPrivilegesArgs struct {
	Domain     *string `pulumi:"domain"`
	Format     *string `pulumi:"format"`
	Name       *string `pulumi:"name"`
	Repository *string `pulumi:"repository"`
	Type       *string `pulumi:"type"`
}

// A collection of values returned by GetNexusPrivileges.
type GetNexusPrivilegesResult struct {
	Domain     *string                       `pulumi:"domain"`
	Format     *string                       `pulumi:"format"`
	Id         string                        `pulumi:"id"`
	Name       *string                       `pulumi:"name"`
	Privileges []GetNexusPrivilegesPrivilege `pulumi:"privileges"`
	Repository *string                       `pulumi:"repository"`
	Type       *string                       `pulumi:"type"`
}

func GetNexusPrivilegesOutput(ctx *pulumi.Context, args GetNexusPrivilegesOutputArgs, opts ...pulumi.InvokeOption) GetNexusPrivilegesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNexusPrivilegesResult, error) {
			args := v.(GetNexusPrivilegesArgs)
			r, err := GetNexusPrivileges(ctx, &args, opts...)
			var s GetNexusPrivilegesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNexusPrivilegesResultOutput)
}

// A collection of arguments for invoking GetNexusPrivileges.
type GetNexusPrivilegesOutputArgs struct {
	Domain     pulumi.StringPtrInput `pulumi:"domain"`
	Format     pulumi.StringPtrInput `pulumi:"format"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Repository pulumi.StringPtrInput `pulumi:"repository"`
	Type       pulumi.StringPtrInput `pulumi:"type"`
}

func (GetNexusPrivilegesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNexusPrivilegesArgs)(nil)).Elem()
}

// A collection of values returned by GetNexusPrivileges.
type GetNexusPrivilegesResultOutput struct{ *pulumi.OutputState }

func (GetNexusPrivilegesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNexusPrivilegesResult)(nil)).Elem()
}

func (o GetNexusPrivilegesResultOutput) ToGetNexusPrivilegesResultOutput() GetNexusPrivilegesResultOutput {
	return o
}

func (o GetNexusPrivilegesResultOutput) ToGetNexusPrivilegesResultOutputWithContext(ctx context.Context) GetNexusPrivilegesResultOutput {
	return o
}

func (o GetNexusPrivilegesResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNexusPrivilegesResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o GetNexusPrivilegesResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNexusPrivilegesResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o GetNexusPrivilegesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNexusPrivilegesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNexusPrivilegesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNexusPrivilegesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetNexusPrivilegesResultOutput) Privileges() GetNexusPrivilegesPrivilegeArrayOutput {
	return o.ApplyT(func(v GetNexusPrivilegesResult) []GetNexusPrivilegesPrivilege { return v.Privileges }).(GetNexusPrivilegesPrivilegeArrayOutput)
}

func (o GetNexusPrivilegesResultOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNexusPrivilegesResult) *string { return v.Repository }).(pulumi.StringPtrOutput)
}

func (o GetNexusPrivilegesResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNexusPrivilegesResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNexusPrivilegesResultOutput{})
}