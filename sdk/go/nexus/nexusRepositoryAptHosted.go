// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NexusRepositoryAptHosted struct {
	pulumi.CustomResourceState

	// Cleanup policies
	Cleanups NexusRepositoryAptHostedCleanupArrayOutput `pulumi:"cleanups"`
	// Component configuration for the hosted repository
	Component NexusRepositoryAptHostedComponentOutput `pulumi:"component"`
	// Distribution to fetch
	Distribution pulumi.StringOutput `pulumi:"distribution"`
	// A unique identifier for this repository
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether this repository accepts incoming requests
	Online pulumi.BoolPtrOutput `pulumi:"online"`
	// Signing contains signing data of hosted repositores of format Apt
	Signing NexusRepositoryAptHostedSigningOutput `pulumi:"signing"`
	// The storage configuration of the repository
	Storage NexusRepositoryAptHostedStorageOutput `pulumi:"storage"`
}

// NewNexusRepositoryAptHosted registers a new resource with the given unique name, arguments, and options.
func NewNexusRepositoryAptHosted(ctx *pulumi.Context,
	name string, args *NexusRepositoryAptHostedArgs, opts ...pulumi.ResourceOption) (*NexusRepositoryAptHosted, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Distribution == nil {
		return nil, errors.New("invalid value for required argument 'Distribution'")
	}
	if args.Signing == nil {
		return nil, errors.New("invalid value for required argument 'Signing'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NexusRepositoryAptHosted
	err := ctx.RegisterResource("nexus:index/nexusRepositoryAptHosted:NexusRepositoryAptHosted", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNexusRepositoryAptHosted gets an existing NexusRepositoryAptHosted resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNexusRepositoryAptHosted(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NexusRepositoryAptHostedState, opts ...pulumi.ResourceOption) (*NexusRepositoryAptHosted, error) {
	var resource NexusRepositoryAptHosted
	err := ctx.ReadResource("nexus:index/nexusRepositoryAptHosted:NexusRepositoryAptHosted", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NexusRepositoryAptHosted resources.
type nexusRepositoryAptHostedState struct {
	// Cleanup policies
	Cleanups []NexusRepositoryAptHostedCleanup `pulumi:"cleanups"`
	// Component configuration for the hosted repository
	Component *NexusRepositoryAptHostedComponent `pulumi:"component"`
	// Distribution to fetch
	Distribution *string `pulumi:"distribution"`
	// A unique identifier for this repository
	Name *string `pulumi:"name"`
	// Whether this repository accepts incoming requests
	Online *bool `pulumi:"online"`
	// Signing contains signing data of hosted repositores of format Apt
	Signing *NexusRepositoryAptHostedSigning `pulumi:"signing"`
	// The storage configuration of the repository
	Storage *NexusRepositoryAptHostedStorage `pulumi:"storage"`
}

type NexusRepositoryAptHostedState struct {
	// Cleanup policies
	Cleanups NexusRepositoryAptHostedCleanupArrayInput
	// Component configuration for the hosted repository
	Component NexusRepositoryAptHostedComponentPtrInput
	// Distribution to fetch
	Distribution pulumi.StringPtrInput
	// A unique identifier for this repository
	Name pulumi.StringPtrInput
	// Whether this repository accepts incoming requests
	Online pulumi.BoolPtrInput
	// Signing contains signing data of hosted repositores of format Apt
	Signing NexusRepositoryAptHostedSigningPtrInput
	// The storage configuration of the repository
	Storage NexusRepositoryAptHostedStoragePtrInput
}

func (NexusRepositoryAptHostedState) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusRepositoryAptHostedState)(nil)).Elem()
}

type nexusRepositoryAptHostedArgs struct {
	// Cleanup policies
	Cleanups []NexusRepositoryAptHostedCleanup `pulumi:"cleanups"`
	// Component configuration for the hosted repository
	Component *NexusRepositoryAptHostedComponent `pulumi:"component"`
	// Distribution to fetch
	Distribution string `pulumi:"distribution"`
	// A unique identifier for this repository
	Name *string `pulumi:"name"`
	// Whether this repository accepts incoming requests
	Online *bool `pulumi:"online"`
	// Signing contains signing data of hosted repositores of format Apt
	Signing NexusRepositoryAptHostedSigning `pulumi:"signing"`
	// The storage configuration of the repository
	Storage NexusRepositoryAptHostedStorage `pulumi:"storage"`
}

// The set of arguments for constructing a NexusRepositoryAptHosted resource.
type NexusRepositoryAptHostedArgs struct {
	// Cleanup policies
	Cleanups NexusRepositoryAptHostedCleanupArrayInput
	// Component configuration for the hosted repository
	Component NexusRepositoryAptHostedComponentPtrInput
	// Distribution to fetch
	Distribution pulumi.StringInput
	// A unique identifier for this repository
	Name pulumi.StringPtrInput
	// Whether this repository accepts incoming requests
	Online pulumi.BoolPtrInput
	// Signing contains signing data of hosted repositores of format Apt
	Signing NexusRepositoryAptHostedSigningInput
	// The storage configuration of the repository
	Storage NexusRepositoryAptHostedStorageInput
}

func (NexusRepositoryAptHostedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusRepositoryAptHostedArgs)(nil)).Elem()
}

type NexusRepositoryAptHostedInput interface {
	pulumi.Input

	ToNexusRepositoryAptHostedOutput() NexusRepositoryAptHostedOutput
	ToNexusRepositoryAptHostedOutputWithContext(ctx context.Context) NexusRepositoryAptHostedOutput
}

func (*NexusRepositoryAptHosted) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusRepositoryAptHosted)(nil)).Elem()
}

func (i *NexusRepositoryAptHosted) ToNexusRepositoryAptHostedOutput() NexusRepositoryAptHostedOutput {
	return i.ToNexusRepositoryAptHostedOutputWithContext(context.Background())
}

func (i *NexusRepositoryAptHosted) ToNexusRepositoryAptHostedOutputWithContext(ctx context.Context) NexusRepositoryAptHostedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRepositoryAptHostedOutput)
}

// NexusRepositoryAptHostedArrayInput is an input type that accepts NexusRepositoryAptHostedArray and NexusRepositoryAptHostedArrayOutput values.
// You can construct a concrete instance of `NexusRepositoryAptHostedArrayInput` via:
//
//          NexusRepositoryAptHostedArray{ NexusRepositoryAptHostedArgs{...} }
type NexusRepositoryAptHostedArrayInput interface {
	pulumi.Input

	ToNexusRepositoryAptHostedArrayOutput() NexusRepositoryAptHostedArrayOutput
	ToNexusRepositoryAptHostedArrayOutputWithContext(context.Context) NexusRepositoryAptHostedArrayOutput
}

type NexusRepositoryAptHostedArray []NexusRepositoryAptHostedInput

func (NexusRepositoryAptHostedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusRepositoryAptHosted)(nil)).Elem()
}

func (i NexusRepositoryAptHostedArray) ToNexusRepositoryAptHostedArrayOutput() NexusRepositoryAptHostedArrayOutput {
	return i.ToNexusRepositoryAptHostedArrayOutputWithContext(context.Background())
}

func (i NexusRepositoryAptHostedArray) ToNexusRepositoryAptHostedArrayOutputWithContext(ctx context.Context) NexusRepositoryAptHostedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRepositoryAptHostedArrayOutput)
}

// NexusRepositoryAptHostedMapInput is an input type that accepts NexusRepositoryAptHostedMap and NexusRepositoryAptHostedMapOutput values.
// You can construct a concrete instance of `NexusRepositoryAptHostedMapInput` via:
//
//          NexusRepositoryAptHostedMap{ "key": NexusRepositoryAptHostedArgs{...} }
type NexusRepositoryAptHostedMapInput interface {
	pulumi.Input

	ToNexusRepositoryAptHostedMapOutput() NexusRepositoryAptHostedMapOutput
	ToNexusRepositoryAptHostedMapOutputWithContext(context.Context) NexusRepositoryAptHostedMapOutput
}

type NexusRepositoryAptHostedMap map[string]NexusRepositoryAptHostedInput

func (NexusRepositoryAptHostedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusRepositoryAptHosted)(nil)).Elem()
}

func (i NexusRepositoryAptHostedMap) ToNexusRepositoryAptHostedMapOutput() NexusRepositoryAptHostedMapOutput {
	return i.ToNexusRepositoryAptHostedMapOutputWithContext(context.Background())
}

func (i NexusRepositoryAptHostedMap) ToNexusRepositoryAptHostedMapOutputWithContext(ctx context.Context) NexusRepositoryAptHostedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRepositoryAptHostedMapOutput)
}

type NexusRepositoryAptHostedOutput struct{ *pulumi.OutputState }

func (NexusRepositoryAptHostedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusRepositoryAptHosted)(nil)).Elem()
}

func (o NexusRepositoryAptHostedOutput) ToNexusRepositoryAptHostedOutput() NexusRepositoryAptHostedOutput {
	return o
}

func (o NexusRepositoryAptHostedOutput) ToNexusRepositoryAptHostedOutputWithContext(ctx context.Context) NexusRepositoryAptHostedOutput {
	return o
}

// Cleanup policies
func (o NexusRepositoryAptHostedOutput) Cleanups() NexusRepositoryAptHostedCleanupArrayOutput {
	return o.ApplyT(func(v *NexusRepositoryAptHosted) NexusRepositoryAptHostedCleanupArrayOutput { return v.Cleanups }).(NexusRepositoryAptHostedCleanupArrayOutput)
}

// Component configuration for the hosted repository
func (o NexusRepositoryAptHostedOutput) Component() NexusRepositoryAptHostedComponentOutput {
	return o.ApplyT(func(v *NexusRepositoryAptHosted) NexusRepositoryAptHostedComponentOutput { return v.Component }).(NexusRepositoryAptHostedComponentOutput)
}

// Distribution to fetch
func (o NexusRepositoryAptHostedOutput) Distribution() pulumi.StringOutput {
	return o.ApplyT(func(v *NexusRepositoryAptHosted) pulumi.StringOutput { return v.Distribution }).(pulumi.StringOutput)
}

// A unique identifier for this repository
func (o NexusRepositoryAptHostedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NexusRepositoryAptHosted) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether this repository accepts incoming requests
func (o NexusRepositoryAptHostedOutput) Online() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NexusRepositoryAptHosted) pulumi.BoolPtrOutput { return v.Online }).(pulumi.BoolPtrOutput)
}

// Signing contains signing data of hosted repositores of format Apt
func (o NexusRepositoryAptHostedOutput) Signing() NexusRepositoryAptHostedSigningOutput {
	return o.ApplyT(func(v *NexusRepositoryAptHosted) NexusRepositoryAptHostedSigningOutput { return v.Signing }).(NexusRepositoryAptHostedSigningOutput)
}

// The storage configuration of the repository
func (o NexusRepositoryAptHostedOutput) Storage() NexusRepositoryAptHostedStorageOutput {
	return o.ApplyT(func(v *NexusRepositoryAptHosted) NexusRepositoryAptHostedStorageOutput { return v.Storage }).(NexusRepositoryAptHostedStorageOutput)
}

type NexusRepositoryAptHostedArrayOutput struct{ *pulumi.OutputState }

func (NexusRepositoryAptHostedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusRepositoryAptHosted)(nil)).Elem()
}

func (o NexusRepositoryAptHostedArrayOutput) ToNexusRepositoryAptHostedArrayOutput() NexusRepositoryAptHostedArrayOutput {
	return o
}

func (o NexusRepositoryAptHostedArrayOutput) ToNexusRepositoryAptHostedArrayOutputWithContext(ctx context.Context) NexusRepositoryAptHostedArrayOutput {
	return o
}

func (o NexusRepositoryAptHostedArrayOutput) Index(i pulumi.IntInput) NexusRepositoryAptHostedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NexusRepositoryAptHosted {
		return vs[0].([]*NexusRepositoryAptHosted)[vs[1].(int)]
	}).(NexusRepositoryAptHostedOutput)
}

type NexusRepositoryAptHostedMapOutput struct{ *pulumi.OutputState }

func (NexusRepositoryAptHostedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusRepositoryAptHosted)(nil)).Elem()
}

func (o NexusRepositoryAptHostedMapOutput) ToNexusRepositoryAptHostedMapOutput() NexusRepositoryAptHostedMapOutput {
	return o
}

func (o NexusRepositoryAptHostedMapOutput) ToNexusRepositoryAptHostedMapOutputWithContext(ctx context.Context) NexusRepositoryAptHostedMapOutput {
	return o
}

func (o NexusRepositoryAptHostedMapOutput) MapIndex(k pulumi.StringInput) NexusRepositoryAptHostedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NexusRepositoryAptHosted {
		return vs[0].(map[string]*NexusRepositoryAptHosted)[vs[1].(string)]
	}).(NexusRepositoryAptHostedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRepositoryAptHostedInput)(nil)).Elem(), &NexusRepositoryAptHosted{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRepositoryAptHostedArrayInput)(nil)).Elem(), NexusRepositoryAptHostedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRepositoryAptHostedMapInput)(nil)).Elem(), NexusRepositoryAptHostedMap{})
	pulumi.RegisterOutputType(NexusRepositoryAptHostedOutput{})
	pulumi.RegisterOutputType(NexusRepositoryAptHostedArrayOutput{})
	pulumi.RegisterOutputType(NexusRepositoryAptHostedMapOutput{})
}