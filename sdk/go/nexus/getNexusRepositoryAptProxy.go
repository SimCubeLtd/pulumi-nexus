// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNexusRepositoryAptProxy(ctx *pulumi.Context, args *LookupNexusRepositoryAptProxyArgs, opts ...pulumi.InvokeOption) (*LookupNexusRepositoryAptProxyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNexusRepositoryAptProxyResult
	err := ctx.Invoke("nexus:index/getNexusRepositoryAptProxy:GetNexusRepositoryAptProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetNexusRepositoryAptProxy.
type LookupNexusRepositoryAptProxyArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by GetNexusRepositoryAptProxy.
type LookupNexusRepositoryAptProxyResult struct {
	Cleanups       []GetNexusRepositoryAptProxyCleanup      `pulumi:"cleanups"`
	Distribution   string                                   `pulumi:"distribution"`
	Flat           bool                                     `pulumi:"flat"`
	HttpClients    []GetNexusRepositoryAptProxyHttpClient   `pulumi:"httpClients"`
	Id             string                                   `pulumi:"id"`
	Name           string                                   `pulumi:"name"`
	NegativeCaches []GetNexusRepositoryAptProxyNegativeCach `pulumi:"negativeCaches"`
	Online         bool                                     `pulumi:"online"`
	Proxies        []GetNexusRepositoryAptProxyProxy        `pulumi:"proxies"`
	RoutingRule    string                                   `pulumi:"routingRule"`
	Storages       []GetNexusRepositoryAptProxyStorage      `pulumi:"storages"`
}

func LookupNexusRepositoryAptProxyOutput(ctx *pulumi.Context, args LookupNexusRepositoryAptProxyOutputArgs, opts ...pulumi.InvokeOption) LookupNexusRepositoryAptProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNexusRepositoryAptProxyResult, error) {
			args := v.(LookupNexusRepositoryAptProxyArgs)
			r, err := LookupNexusRepositoryAptProxy(ctx, &args, opts...)
			var s LookupNexusRepositoryAptProxyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNexusRepositoryAptProxyResultOutput)
}

// A collection of arguments for invoking GetNexusRepositoryAptProxy.
type LookupNexusRepositoryAptProxyOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupNexusRepositoryAptProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusRepositoryAptProxyArgs)(nil)).Elem()
}

// A collection of values returned by GetNexusRepositoryAptProxy.
type LookupNexusRepositoryAptProxyResultOutput struct{ *pulumi.OutputState }

func (LookupNexusRepositoryAptProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusRepositoryAptProxyResult)(nil)).Elem()
}

func (o LookupNexusRepositoryAptProxyResultOutput) ToLookupNexusRepositoryAptProxyResultOutput() LookupNexusRepositoryAptProxyResultOutput {
	return o
}

func (o LookupNexusRepositoryAptProxyResultOutput) ToLookupNexusRepositoryAptProxyResultOutputWithContext(ctx context.Context) LookupNexusRepositoryAptProxyResultOutput {
	return o
}

func (o LookupNexusRepositoryAptProxyResultOutput) Cleanups() GetNexusRepositoryAptProxyCleanupArrayOutput {
	return o.ApplyT(func(v LookupNexusRepositoryAptProxyResult) []GetNexusRepositoryAptProxyCleanup { return v.Cleanups }).(GetNexusRepositoryAptProxyCleanupArrayOutput)
}

func (o LookupNexusRepositoryAptProxyResultOutput) Distribution() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusRepositoryAptProxyResult) string { return v.Distribution }).(pulumi.StringOutput)
}

func (o LookupNexusRepositoryAptProxyResultOutput) Flat() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNexusRepositoryAptProxyResult) bool { return v.Flat }).(pulumi.BoolOutput)
}

func (o LookupNexusRepositoryAptProxyResultOutput) HttpClients() GetNexusRepositoryAptProxyHttpClientArrayOutput {
	return o.ApplyT(func(v LookupNexusRepositoryAptProxyResult) []GetNexusRepositoryAptProxyHttpClient {
		return v.HttpClients
	}).(GetNexusRepositoryAptProxyHttpClientArrayOutput)
}

func (o LookupNexusRepositoryAptProxyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusRepositoryAptProxyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNexusRepositoryAptProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusRepositoryAptProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNexusRepositoryAptProxyResultOutput) NegativeCaches() GetNexusRepositoryAptProxyNegativeCachArrayOutput {
	return o.ApplyT(func(v LookupNexusRepositoryAptProxyResult) []GetNexusRepositoryAptProxyNegativeCach {
		return v.NegativeCaches
	}).(GetNexusRepositoryAptProxyNegativeCachArrayOutput)
}

func (o LookupNexusRepositoryAptProxyResultOutput) Online() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNexusRepositoryAptProxyResult) bool { return v.Online }).(pulumi.BoolOutput)
}

func (o LookupNexusRepositoryAptProxyResultOutput) Proxies() GetNexusRepositoryAptProxyProxyArrayOutput {
	return o.ApplyT(func(v LookupNexusRepositoryAptProxyResult) []GetNexusRepositoryAptProxyProxy { return v.Proxies }).(GetNexusRepositoryAptProxyProxyArrayOutput)
}

func (o LookupNexusRepositoryAptProxyResultOutput) RoutingRule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusRepositoryAptProxyResult) string { return v.RoutingRule }).(pulumi.StringOutput)
}

func (o LookupNexusRepositoryAptProxyResultOutput) Storages() GetNexusRepositoryAptProxyStorageArrayOutput {
	return o.ApplyT(func(v LookupNexusRepositoryAptProxyResult) []GetNexusRepositoryAptProxyStorage { return v.Storages }).(GetNexusRepositoryAptProxyStorageArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNexusRepositoryAptProxyResultOutput{})
}
