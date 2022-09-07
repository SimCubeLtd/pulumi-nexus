// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the nexus package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// Password of user to connect to API. Reading environment variable NEXUS_PASSWORD. Default:`admin123`
	Password pulumi.StringOutput `pulumi:"password"`
	// URL of Nexus to reach API. Reading environment variable NEXUS_URL. Default:`http://127.0.0.1:8080`
	Url pulumi.StringOutput `pulumi:"url"`
	// Username used to connect to API. Reading environment variable NEXUS_USERNAME. Default:`admin`
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:nexus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// Boolean to specify wether insecure SSL connections are allowed or not. Reading environment variable
	// NEXUS_INSECURE_SKIP_VERIFY. Default:`true`
	Insecure *bool `pulumi:"insecure"`
	// Password of user to connect to API. Reading environment variable NEXUS_PASSWORD. Default:`admin123`
	Password string `pulumi:"password"`
	// URL of Nexus to reach API. Reading environment variable NEXUS_URL. Default:`http://127.0.0.1:8080`
	Url string `pulumi:"url"`
	// Username used to connect to API. Reading environment variable NEXUS_USERNAME. Default:`admin`
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// Boolean to specify wether insecure SSL connections are allowed or not. Reading environment variable
	// NEXUS_INSECURE_SKIP_VERIFY. Default:`true`
	Insecure pulumi.BoolPtrInput
	// Password of user to connect to API. Reading environment variable NEXUS_PASSWORD. Default:`admin123`
	Password pulumi.StringInput
	// URL of Nexus to reach API. Reading environment variable NEXUS_URL. Default:`http://127.0.0.1:8080`
	Url pulumi.StringInput
	// Username used to connect to API. Reading environment variable NEXUS_USERNAME. Default:`admin`
	Username pulumi.StringInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// Password of user to connect to API. Reading environment variable NEXUS_PASSWORD. Default:`admin123`
func (o ProviderOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// URL of Nexus to reach API. Reading environment variable NEXUS_URL. Default:`http://127.0.0.1:8080`
func (o ProviderOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Username used to connect to API. Reading environment variable NEXUS_USERNAME. Default:`admin`
func (o ProviderOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}