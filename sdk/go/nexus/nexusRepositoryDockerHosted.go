// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NexusRepositoryDockerHosted struct {
	pulumi.CustomResourceState

	// Cleanup policies
	Cleanups NexusRepositoryDockerHostedCleanupArrayOutput `pulumi:"cleanups"`
	// Component configuration for the hosted repository
	Component NexusRepositoryDockerHostedComponentOutput `pulumi:"component"`
	// docker contains the configuration of the docker repository
	Docker NexusRepositoryDockerHostedDockerOutput `pulumi:"docker"`
	// A unique identifier for this repository
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether this repository accepts incoming requests
	Online pulumi.BoolPtrOutput `pulumi:"online"`
	// The storage configuration of the repository
	Storage NexusRepositoryDockerHostedStorageOutput `pulumi:"storage"`
}

// NewNexusRepositoryDockerHosted registers a new resource with the given unique name, arguments, and options.
func NewNexusRepositoryDockerHosted(ctx *pulumi.Context,
	name string, args *NexusRepositoryDockerHostedArgs, opts ...pulumi.ResourceOption) (*NexusRepositoryDockerHosted, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Docker == nil {
		return nil, errors.New("invalid value for required argument 'Docker'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NexusRepositoryDockerHosted
	err := ctx.RegisterResource("nexus:index/nexusRepositoryDockerHosted:NexusRepositoryDockerHosted", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNexusRepositoryDockerHosted gets an existing NexusRepositoryDockerHosted resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNexusRepositoryDockerHosted(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NexusRepositoryDockerHostedState, opts ...pulumi.ResourceOption) (*NexusRepositoryDockerHosted, error) {
	var resource NexusRepositoryDockerHosted
	err := ctx.ReadResource("nexus:index/nexusRepositoryDockerHosted:NexusRepositoryDockerHosted", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NexusRepositoryDockerHosted resources.
type nexusRepositoryDockerHostedState struct {
	// Cleanup policies
	Cleanups []NexusRepositoryDockerHostedCleanup `pulumi:"cleanups"`
	// Component configuration for the hosted repository
	Component *NexusRepositoryDockerHostedComponent `pulumi:"component"`
	// docker contains the configuration of the docker repository
	Docker *NexusRepositoryDockerHostedDocker `pulumi:"docker"`
	// A unique identifier for this repository
	Name *string `pulumi:"name"`
	// Whether this repository accepts incoming requests
	Online *bool `pulumi:"online"`
	// The storage configuration of the repository
	Storage *NexusRepositoryDockerHostedStorage `pulumi:"storage"`
}

type NexusRepositoryDockerHostedState struct {
	// Cleanup policies
	Cleanups NexusRepositoryDockerHostedCleanupArrayInput
	// Component configuration for the hosted repository
	Component NexusRepositoryDockerHostedComponentPtrInput
	// docker contains the configuration of the docker repository
	Docker NexusRepositoryDockerHostedDockerPtrInput
	// A unique identifier for this repository
	Name pulumi.StringPtrInput
	// Whether this repository accepts incoming requests
	Online pulumi.BoolPtrInput
	// The storage configuration of the repository
	Storage NexusRepositoryDockerHostedStoragePtrInput
}

func (NexusRepositoryDockerHostedState) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusRepositoryDockerHostedState)(nil)).Elem()
}

type nexusRepositoryDockerHostedArgs struct {
	// Cleanup policies
	Cleanups []NexusRepositoryDockerHostedCleanup `pulumi:"cleanups"`
	// Component configuration for the hosted repository
	Component *NexusRepositoryDockerHostedComponent `pulumi:"component"`
	// docker contains the configuration of the docker repository
	Docker NexusRepositoryDockerHostedDocker `pulumi:"docker"`
	// A unique identifier for this repository
	Name *string `pulumi:"name"`
	// Whether this repository accepts incoming requests
	Online *bool `pulumi:"online"`
	// The storage configuration of the repository
	Storage NexusRepositoryDockerHostedStorage `pulumi:"storage"`
}

// The set of arguments for constructing a NexusRepositoryDockerHosted resource.
type NexusRepositoryDockerHostedArgs struct {
	// Cleanup policies
	Cleanups NexusRepositoryDockerHostedCleanupArrayInput
	// Component configuration for the hosted repository
	Component NexusRepositoryDockerHostedComponentPtrInput
	// docker contains the configuration of the docker repository
	Docker NexusRepositoryDockerHostedDockerInput
	// A unique identifier for this repository
	Name pulumi.StringPtrInput
	// Whether this repository accepts incoming requests
	Online pulumi.BoolPtrInput
	// The storage configuration of the repository
	Storage NexusRepositoryDockerHostedStorageInput
}

func (NexusRepositoryDockerHostedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusRepositoryDockerHostedArgs)(nil)).Elem()
}

type NexusRepositoryDockerHostedInput interface {
	pulumi.Input

	ToNexusRepositoryDockerHostedOutput() NexusRepositoryDockerHostedOutput
	ToNexusRepositoryDockerHostedOutputWithContext(ctx context.Context) NexusRepositoryDockerHostedOutput
}

func (*NexusRepositoryDockerHosted) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusRepositoryDockerHosted)(nil)).Elem()
}

func (i *NexusRepositoryDockerHosted) ToNexusRepositoryDockerHostedOutput() NexusRepositoryDockerHostedOutput {
	return i.ToNexusRepositoryDockerHostedOutputWithContext(context.Background())
}

func (i *NexusRepositoryDockerHosted) ToNexusRepositoryDockerHostedOutputWithContext(ctx context.Context) NexusRepositoryDockerHostedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRepositoryDockerHostedOutput)
}

// NexusRepositoryDockerHostedArrayInput is an input type that accepts NexusRepositoryDockerHostedArray and NexusRepositoryDockerHostedArrayOutput values.
// You can construct a concrete instance of `NexusRepositoryDockerHostedArrayInput` via:
//
//          NexusRepositoryDockerHostedArray{ NexusRepositoryDockerHostedArgs{...} }
type NexusRepositoryDockerHostedArrayInput interface {
	pulumi.Input

	ToNexusRepositoryDockerHostedArrayOutput() NexusRepositoryDockerHostedArrayOutput
	ToNexusRepositoryDockerHostedArrayOutputWithContext(context.Context) NexusRepositoryDockerHostedArrayOutput
}

type NexusRepositoryDockerHostedArray []NexusRepositoryDockerHostedInput

func (NexusRepositoryDockerHostedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusRepositoryDockerHosted)(nil)).Elem()
}

func (i NexusRepositoryDockerHostedArray) ToNexusRepositoryDockerHostedArrayOutput() NexusRepositoryDockerHostedArrayOutput {
	return i.ToNexusRepositoryDockerHostedArrayOutputWithContext(context.Background())
}

func (i NexusRepositoryDockerHostedArray) ToNexusRepositoryDockerHostedArrayOutputWithContext(ctx context.Context) NexusRepositoryDockerHostedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRepositoryDockerHostedArrayOutput)
}

// NexusRepositoryDockerHostedMapInput is an input type that accepts NexusRepositoryDockerHostedMap and NexusRepositoryDockerHostedMapOutput values.
// You can construct a concrete instance of `NexusRepositoryDockerHostedMapInput` via:
//
//          NexusRepositoryDockerHostedMap{ "key": NexusRepositoryDockerHostedArgs{...} }
type NexusRepositoryDockerHostedMapInput interface {
	pulumi.Input

	ToNexusRepositoryDockerHostedMapOutput() NexusRepositoryDockerHostedMapOutput
	ToNexusRepositoryDockerHostedMapOutputWithContext(context.Context) NexusRepositoryDockerHostedMapOutput
}

type NexusRepositoryDockerHostedMap map[string]NexusRepositoryDockerHostedInput

func (NexusRepositoryDockerHostedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusRepositoryDockerHosted)(nil)).Elem()
}

func (i NexusRepositoryDockerHostedMap) ToNexusRepositoryDockerHostedMapOutput() NexusRepositoryDockerHostedMapOutput {
	return i.ToNexusRepositoryDockerHostedMapOutputWithContext(context.Background())
}

func (i NexusRepositoryDockerHostedMap) ToNexusRepositoryDockerHostedMapOutputWithContext(ctx context.Context) NexusRepositoryDockerHostedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRepositoryDockerHostedMapOutput)
}

type NexusRepositoryDockerHostedOutput struct{ *pulumi.OutputState }

func (NexusRepositoryDockerHostedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusRepositoryDockerHosted)(nil)).Elem()
}

func (o NexusRepositoryDockerHostedOutput) ToNexusRepositoryDockerHostedOutput() NexusRepositoryDockerHostedOutput {
	return o
}

func (o NexusRepositoryDockerHostedOutput) ToNexusRepositoryDockerHostedOutputWithContext(ctx context.Context) NexusRepositoryDockerHostedOutput {
	return o
}

// Cleanup policies
func (o NexusRepositoryDockerHostedOutput) Cleanups() NexusRepositoryDockerHostedCleanupArrayOutput {
	return o.ApplyT(func(v *NexusRepositoryDockerHosted) NexusRepositoryDockerHostedCleanupArrayOutput { return v.Cleanups }).(NexusRepositoryDockerHostedCleanupArrayOutput)
}

// Component configuration for the hosted repository
func (o NexusRepositoryDockerHostedOutput) Component() NexusRepositoryDockerHostedComponentOutput {
	return o.ApplyT(func(v *NexusRepositoryDockerHosted) NexusRepositoryDockerHostedComponentOutput { return v.Component }).(NexusRepositoryDockerHostedComponentOutput)
}

// docker contains the configuration of the docker repository
func (o NexusRepositoryDockerHostedOutput) Docker() NexusRepositoryDockerHostedDockerOutput {
	return o.ApplyT(func(v *NexusRepositoryDockerHosted) NexusRepositoryDockerHostedDockerOutput { return v.Docker }).(NexusRepositoryDockerHostedDockerOutput)
}

// A unique identifier for this repository
func (o NexusRepositoryDockerHostedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NexusRepositoryDockerHosted) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether this repository accepts incoming requests
func (o NexusRepositoryDockerHostedOutput) Online() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NexusRepositoryDockerHosted) pulumi.BoolPtrOutput { return v.Online }).(pulumi.BoolPtrOutput)
}

// The storage configuration of the repository
func (o NexusRepositoryDockerHostedOutput) Storage() NexusRepositoryDockerHostedStorageOutput {
	return o.ApplyT(func(v *NexusRepositoryDockerHosted) NexusRepositoryDockerHostedStorageOutput { return v.Storage }).(NexusRepositoryDockerHostedStorageOutput)
}

type NexusRepositoryDockerHostedArrayOutput struct{ *pulumi.OutputState }

func (NexusRepositoryDockerHostedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusRepositoryDockerHosted)(nil)).Elem()
}

func (o NexusRepositoryDockerHostedArrayOutput) ToNexusRepositoryDockerHostedArrayOutput() NexusRepositoryDockerHostedArrayOutput {
	return o
}

func (o NexusRepositoryDockerHostedArrayOutput) ToNexusRepositoryDockerHostedArrayOutputWithContext(ctx context.Context) NexusRepositoryDockerHostedArrayOutput {
	return o
}

func (o NexusRepositoryDockerHostedArrayOutput) Index(i pulumi.IntInput) NexusRepositoryDockerHostedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NexusRepositoryDockerHosted {
		return vs[0].([]*NexusRepositoryDockerHosted)[vs[1].(int)]
	}).(NexusRepositoryDockerHostedOutput)
}

type NexusRepositoryDockerHostedMapOutput struct{ *pulumi.OutputState }

func (NexusRepositoryDockerHostedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusRepositoryDockerHosted)(nil)).Elem()
}

func (o NexusRepositoryDockerHostedMapOutput) ToNexusRepositoryDockerHostedMapOutput() NexusRepositoryDockerHostedMapOutput {
	return o
}

func (o NexusRepositoryDockerHostedMapOutput) ToNexusRepositoryDockerHostedMapOutputWithContext(ctx context.Context) NexusRepositoryDockerHostedMapOutput {
	return o
}

func (o NexusRepositoryDockerHostedMapOutput) MapIndex(k pulumi.StringInput) NexusRepositoryDockerHostedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NexusRepositoryDockerHosted {
		return vs[0].(map[string]*NexusRepositoryDockerHosted)[vs[1].(string)]
	}).(NexusRepositoryDockerHostedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRepositoryDockerHostedInput)(nil)).Elem(), &NexusRepositoryDockerHosted{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRepositoryDockerHostedArrayInput)(nil)).Elem(), NexusRepositoryDockerHostedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRepositoryDockerHostedMapInput)(nil)).Elem(), NexusRepositoryDockerHostedMap{})
	pulumi.RegisterOutputType(NexusRepositoryDockerHostedOutput{})
	pulumi.RegisterOutputType(NexusRepositoryDockerHostedArrayOutput{})
	pulumi.RegisterOutputType(NexusRepositoryDockerHostedMapOutput{})
}
