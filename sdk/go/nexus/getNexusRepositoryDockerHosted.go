// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNexusRepositoryDockerHosted(ctx *pulumi.Context, args *LookupNexusRepositoryDockerHostedArgs, opts ...pulumi.InvokeOption) (*LookupNexusRepositoryDockerHostedResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNexusRepositoryDockerHostedResult
	err := ctx.Invoke("nexus:index/getNexusRepositoryDockerHosted:GetNexusRepositoryDockerHosted", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetNexusRepositoryDockerHosted.
type LookupNexusRepositoryDockerHostedArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by GetNexusRepositoryDockerHosted.
type LookupNexusRepositoryDockerHostedResult struct {
	Cleanups   []GetNexusRepositoryDockerHostedCleanup   `pulumi:"cleanups"`
	Components []GetNexusRepositoryDockerHostedComponent `pulumi:"components"`
	Dockers    []GetNexusRepositoryDockerHostedDocker    `pulumi:"dockers"`
	Id         string                                    `pulumi:"id"`
	Name       string                                    `pulumi:"name"`
	Online     bool                                      `pulumi:"online"`
	Storages   []GetNexusRepositoryDockerHostedStorage   `pulumi:"storages"`
}

func LookupNexusRepositoryDockerHostedOutput(ctx *pulumi.Context, args LookupNexusRepositoryDockerHostedOutputArgs, opts ...pulumi.InvokeOption) LookupNexusRepositoryDockerHostedResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNexusRepositoryDockerHostedResult, error) {
			args := v.(LookupNexusRepositoryDockerHostedArgs)
			r, err := LookupNexusRepositoryDockerHosted(ctx, &args, opts...)
			var s LookupNexusRepositoryDockerHostedResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNexusRepositoryDockerHostedResultOutput)
}

// A collection of arguments for invoking GetNexusRepositoryDockerHosted.
type LookupNexusRepositoryDockerHostedOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupNexusRepositoryDockerHostedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusRepositoryDockerHostedArgs)(nil)).Elem()
}

// A collection of values returned by GetNexusRepositoryDockerHosted.
type LookupNexusRepositoryDockerHostedResultOutput struct{ *pulumi.OutputState }

func (LookupNexusRepositoryDockerHostedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusRepositoryDockerHostedResult)(nil)).Elem()
}

func (o LookupNexusRepositoryDockerHostedResultOutput) ToLookupNexusRepositoryDockerHostedResultOutput() LookupNexusRepositoryDockerHostedResultOutput {
	return o
}

func (o LookupNexusRepositoryDockerHostedResultOutput) ToLookupNexusRepositoryDockerHostedResultOutputWithContext(ctx context.Context) LookupNexusRepositoryDockerHostedResultOutput {
	return o
}

func (o LookupNexusRepositoryDockerHostedResultOutput) Cleanups() GetNexusRepositoryDockerHostedCleanupArrayOutput {
	return o.ApplyT(func(v LookupNexusRepositoryDockerHostedResult) []GetNexusRepositoryDockerHostedCleanup {
		return v.Cleanups
	}).(GetNexusRepositoryDockerHostedCleanupArrayOutput)
}

func (o LookupNexusRepositoryDockerHostedResultOutput) Components() GetNexusRepositoryDockerHostedComponentArrayOutput {
	return o.ApplyT(func(v LookupNexusRepositoryDockerHostedResult) []GetNexusRepositoryDockerHostedComponent {
		return v.Components
	}).(GetNexusRepositoryDockerHostedComponentArrayOutput)
}

func (o LookupNexusRepositoryDockerHostedResultOutput) Dockers() GetNexusRepositoryDockerHostedDockerArrayOutput {
	return o.ApplyT(func(v LookupNexusRepositoryDockerHostedResult) []GetNexusRepositoryDockerHostedDocker {
		return v.Dockers
	}).(GetNexusRepositoryDockerHostedDockerArrayOutput)
}

func (o LookupNexusRepositoryDockerHostedResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusRepositoryDockerHostedResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNexusRepositoryDockerHostedResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusRepositoryDockerHostedResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNexusRepositoryDockerHostedResultOutput) Online() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNexusRepositoryDockerHostedResult) bool { return v.Online }).(pulumi.BoolOutput)
}

func (o LookupNexusRepositoryDockerHostedResultOutput) Storages() GetNexusRepositoryDockerHostedStorageArrayOutput {
	return o.ApplyT(func(v LookupNexusRepositoryDockerHostedResult) []GetNexusRepositoryDockerHostedStorage {
		return v.Storages
	}).(GetNexusRepositoryDockerHostedStorageArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNexusRepositoryDockerHostedResultOutput{})
}
