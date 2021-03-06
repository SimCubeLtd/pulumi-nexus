// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNexusUser(ctx *pulumi.Context, args *LookupNexusUserArgs, opts ...pulumi.InvokeOption) (*LookupNexusUserResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNexusUserResult
	err := ctx.Invoke("nexus:index/getNexusUser:GetNexusUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetNexusUser.
type LookupNexusUserArgs struct {
	Userid string `pulumi:"userid"`
}

// A collection of values returned by GetNexusUser.
type LookupNexusUserResult struct {
	Email     string   `pulumi:"email"`
	Firstname string   `pulumi:"firstname"`
	Id        string   `pulumi:"id"`
	Lastname  string   `pulumi:"lastname"`
	Roles     []string `pulumi:"roles"`
	Status    string   `pulumi:"status"`
	Userid    string   `pulumi:"userid"`
}

func LookupNexusUserOutput(ctx *pulumi.Context, args LookupNexusUserOutputArgs, opts ...pulumi.InvokeOption) LookupNexusUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNexusUserResult, error) {
			args := v.(LookupNexusUserArgs)
			r, err := LookupNexusUser(ctx, &args, opts...)
			var s LookupNexusUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNexusUserResultOutput)
}

// A collection of arguments for invoking GetNexusUser.
type LookupNexusUserOutputArgs struct {
	Userid pulumi.StringInput `pulumi:"userid"`
}

func (LookupNexusUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusUserArgs)(nil)).Elem()
}

// A collection of values returned by GetNexusUser.
type LookupNexusUserResultOutput struct{ *pulumi.OutputState }

func (LookupNexusUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusUserResult)(nil)).Elem()
}

func (o LookupNexusUserResultOutput) ToLookupNexusUserResultOutput() LookupNexusUserResultOutput {
	return o
}

func (o LookupNexusUserResultOutput) ToLookupNexusUserResultOutputWithContext(ctx context.Context) LookupNexusUserResultOutput {
	return o
}

func (o LookupNexusUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusUserResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o LookupNexusUserResultOutput) Firstname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusUserResult) string { return v.Firstname }).(pulumi.StringOutput)
}

func (o LookupNexusUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNexusUserResultOutput) Lastname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusUserResult) string { return v.Lastname }).(pulumi.StringOutput)
}

func (o LookupNexusUserResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNexusUserResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o LookupNexusUserResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusUserResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupNexusUserResultOutput) Userid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusUserResult) string { return v.Userid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNexusUserResultOutput{})
}
