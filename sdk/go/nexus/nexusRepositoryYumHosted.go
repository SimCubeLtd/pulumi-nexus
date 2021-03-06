// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NexusRepositoryYumHosted struct {
	pulumi.CustomResourceState

	// Cleanup policies
	Cleanups NexusRepositoryYumHostedCleanupArrayOutput `pulumi:"cleanups"`
	// Component configuration for the hosted repository
	Component NexusRepositoryYumHostedComponentOutput `pulumi:"component"`
	// Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
	DeployPolicy pulumi.StringPtrOutput `pulumi:"deployPolicy"`
	// A unique identifier for this repository
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether this repository accepts incoming requests
	Online pulumi.BoolPtrOutput `pulumi:"online"`
	// Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
	RepodataDepth pulumi.IntPtrOutput `pulumi:"repodataDepth"`
	// The storage configuration of the repository
	Storage NexusRepositoryYumHostedStorageOutput `pulumi:"storage"`
}

// NewNexusRepositoryYumHosted registers a new resource with the given unique name, arguments, and options.
func NewNexusRepositoryYumHosted(ctx *pulumi.Context,
	name string, args *NexusRepositoryYumHostedArgs, opts ...pulumi.ResourceOption) (*NexusRepositoryYumHosted, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NexusRepositoryYumHosted
	err := ctx.RegisterResource("nexus:index/nexusRepositoryYumHosted:NexusRepositoryYumHosted", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNexusRepositoryYumHosted gets an existing NexusRepositoryYumHosted resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNexusRepositoryYumHosted(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NexusRepositoryYumHostedState, opts ...pulumi.ResourceOption) (*NexusRepositoryYumHosted, error) {
	var resource NexusRepositoryYumHosted
	err := ctx.ReadResource("nexus:index/nexusRepositoryYumHosted:NexusRepositoryYumHosted", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NexusRepositoryYumHosted resources.
type nexusRepositoryYumHostedState struct {
	// Cleanup policies
	Cleanups []NexusRepositoryYumHostedCleanup `pulumi:"cleanups"`
	// Component configuration for the hosted repository
	Component *NexusRepositoryYumHostedComponent `pulumi:"component"`
	// Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
	DeployPolicy *string `pulumi:"deployPolicy"`
	// A unique identifier for this repository
	Name *string `pulumi:"name"`
	// Whether this repository accepts incoming requests
	Online *bool `pulumi:"online"`
	// Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
	RepodataDepth *int `pulumi:"repodataDepth"`
	// The storage configuration of the repository
	Storage *NexusRepositoryYumHostedStorage `pulumi:"storage"`
}

type NexusRepositoryYumHostedState struct {
	// Cleanup policies
	Cleanups NexusRepositoryYumHostedCleanupArrayInput
	// Component configuration for the hosted repository
	Component NexusRepositoryYumHostedComponentPtrInput
	// Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
	DeployPolicy pulumi.StringPtrInput
	// A unique identifier for this repository
	Name pulumi.StringPtrInput
	// Whether this repository accepts incoming requests
	Online pulumi.BoolPtrInput
	// Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
	RepodataDepth pulumi.IntPtrInput
	// The storage configuration of the repository
	Storage NexusRepositoryYumHostedStoragePtrInput
}

func (NexusRepositoryYumHostedState) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusRepositoryYumHostedState)(nil)).Elem()
}

type nexusRepositoryYumHostedArgs struct {
	// Cleanup policies
	Cleanups []NexusRepositoryYumHostedCleanup `pulumi:"cleanups"`
	// Component configuration for the hosted repository
	Component *NexusRepositoryYumHostedComponent `pulumi:"component"`
	// Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
	DeployPolicy *string `pulumi:"deployPolicy"`
	// A unique identifier for this repository
	Name *string `pulumi:"name"`
	// Whether this repository accepts incoming requests
	Online *bool `pulumi:"online"`
	// Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
	RepodataDepth *int `pulumi:"repodataDepth"`
	// The storage configuration of the repository
	Storage NexusRepositoryYumHostedStorage `pulumi:"storage"`
}

// The set of arguments for constructing a NexusRepositoryYumHosted resource.
type NexusRepositoryYumHostedArgs struct {
	// Cleanup policies
	Cleanups NexusRepositoryYumHostedCleanupArrayInput
	// Component configuration for the hosted repository
	Component NexusRepositoryYumHostedComponentPtrInput
	// Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
	DeployPolicy pulumi.StringPtrInput
	// A unique identifier for this repository
	Name pulumi.StringPtrInput
	// Whether this repository accepts incoming requests
	Online pulumi.BoolPtrInput
	// Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
	RepodataDepth pulumi.IntPtrInput
	// The storage configuration of the repository
	Storage NexusRepositoryYumHostedStorageInput
}

func (NexusRepositoryYumHostedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusRepositoryYumHostedArgs)(nil)).Elem()
}

type NexusRepositoryYumHostedInput interface {
	pulumi.Input

	ToNexusRepositoryYumHostedOutput() NexusRepositoryYumHostedOutput
	ToNexusRepositoryYumHostedOutputWithContext(ctx context.Context) NexusRepositoryYumHostedOutput
}

func (*NexusRepositoryYumHosted) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusRepositoryYumHosted)(nil)).Elem()
}

func (i *NexusRepositoryYumHosted) ToNexusRepositoryYumHostedOutput() NexusRepositoryYumHostedOutput {
	return i.ToNexusRepositoryYumHostedOutputWithContext(context.Background())
}

func (i *NexusRepositoryYumHosted) ToNexusRepositoryYumHostedOutputWithContext(ctx context.Context) NexusRepositoryYumHostedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRepositoryYumHostedOutput)
}

// NexusRepositoryYumHostedArrayInput is an input type that accepts NexusRepositoryYumHostedArray and NexusRepositoryYumHostedArrayOutput values.
// You can construct a concrete instance of `NexusRepositoryYumHostedArrayInput` via:
//
//          NexusRepositoryYumHostedArray{ NexusRepositoryYumHostedArgs{...} }
type NexusRepositoryYumHostedArrayInput interface {
	pulumi.Input

	ToNexusRepositoryYumHostedArrayOutput() NexusRepositoryYumHostedArrayOutput
	ToNexusRepositoryYumHostedArrayOutputWithContext(context.Context) NexusRepositoryYumHostedArrayOutput
}

type NexusRepositoryYumHostedArray []NexusRepositoryYumHostedInput

func (NexusRepositoryYumHostedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusRepositoryYumHosted)(nil)).Elem()
}

func (i NexusRepositoryYumHostedArray) ToNexusRepositoryYumHostedArrayOutput() NexusRepositoryYumHostedArrayOutput {
	return i.ToNexusRepositoryYumHostedArrayOutputWithContext(context.Background())
}

func (i NexusRepositoryYumHostedArray) ToNexusRepositoryYumHostedArrayOutputWithContext(ctx context.Context) NexusRepositoryYumHostedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRepositoryYumHostedArrayOutput)
}

// NexusRepositoryYumHostedMapInput is an input type that accepts NexusRepositoryYumHostedMap and NexusRepositoryYumHostedMapOutput values.
// You can construct a concrete instance of `NexusRepositoryYumHostedMapInput` via:
//
//          NexusRepositoryYumHostedMap{ "key": NexusRepositoryYumHostedArgs{...} }
type NexusRepositoryYumHostedMapInput interface {
	pulumi.Input

	ToNexusRepositoryYumHostedMapOutput() NexusRepositoryYumHostedMapOutput
	ToNexusRepositoryYumHostedMapOutputWithContext(context.Context) NexusRepositoryYumHostedMapOutput
}

type NexusRepositoryYumHostedMap map[string]NexusRepositoryYumHostedInput

func (NexusRepositoryYumHostedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusRepositoryYumHosted)(nil)).Elem()
}

func (i NexusRepositoryYumHostedMap) ToNexusRepositoryYumHostedMapOutput() NexusRepositoryYumHostedMapOutput {
	return i.ToNexusRepositoryYumHostedMapOutputWithContext(context.Background())
}

func (i NexusRepositoryYumHostedMap) ToNexusRepositoryYumHostedMapOutputWithContext(ctx context.Context) NexusRepositoryYumHostedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRepositoryYumHostedMapOutput)
}

type NexusRepositoryYumHostedOutput struct{ *pulumi.OutputState }

func (NexusRepositoryYumHostedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusRepositoryYumHosted)(nil)).Elem()
}

func (o NexusRepositoryYumHostedOutput) ToNexusRepositoryYumHostedOutput() NexusRepositoryYumHostedOutput {
	return o
}

func (o NexusRepositoryYumHostedOutput) ToNexusRepositoryYumHostedOutputWithContext(ctx context.Context) NexusRepositoryYumHostedOutput {
	return o
}

// Cleanup policies
func (o NexusRepositoryYumHostedOutput) Cleanups() NexusRepositoryYumHostedCleanupArrayOutput {
	return o.ApplyT(func(v *NexusRepositoryYumHosted) NexusRepositoryYumHostedCleanupArrayOutput { return v.Cleanups }).(NexusRepositoryYumHostedCleanupArrayOutput)
}

// Component configuration for the hosted repository
func (o NexusRepositoryYumHostedOutput) Component() NexusRepositoryYumHostedComponentOutput {
	return o.ApplyT(func(v *NexusRepositoryYumHosted) NexusRepositoryYumHostedComponentOutput { return v.Component }).(NexusRepositoryYumHostedComponentOutput)
}

// Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
func (o NexusRepositoryYumHostedOutput) DeployPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NexusRepositoryYumHosted) pulumi.StringPtrOutput { return v.DeployPolicy }).(pulumi.StringPtrOutput)
}

// A unique identifier for this repository
func (o NexusRepositoryYumHostedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NexusRepositoryYumHosted) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether this repository accepts incoming requests
func (o NexusRepositoryYumHostedOutput) Online() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NexusRepositoryYumHosted) pulumi.BoolPtrOutput { return v.Online }).(pulumi.BoolPtrOutput)
}

// Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
func (o NexusRepositoryYumHostedOutput) RepodataDepth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NexusRepositoryYumHosted) pulumi.IntPtrOutput { return v.RepodataDepth }).(pulumi.IntPtrOutput)
}

// The storage configuration of the repository
func (o NexusRepositoryYumHostedOutput) Storage() NexusRepositoryYumHostedStorageOutput {
	return o.ApplyT(func(v *NexusRepositoryYumHosted) NexusRepositoryYumHostedStorageOutput { return v.Storage }).(NexusRepositoryYumHostedStorageOutput)
}

type NexusRepositoryYumHostedArrayOutput struct{ *pulumi.OutputState }

func (NexusRepositoryYumHostedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusRepositoryYumHosted)(nil)).Elem()
}

func (o NexusRepositoryYumHostedArrayOutput) ToNexusRepositoryYumHostedArrayOutput() NexusRepositoryYumHostedArrayOutput {
	return o
}

func (o NexusRepositoryYumHostedArrayOutput) ToNexusRepositoryYumHostedArrayOutputWithContext(ctx context.Context) NexusRepositoryYumHostedArrayOutput {
	return o
}

func (o NexusRepositoryYumHostedArrayOutput) Index(i pulumi.IntInput) NexusRepositoryYumHostedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NexusRepositoryYumHosted {
		return vs[0].([]*NexusRepositoryYumHosted)[vs[1].(int)]
	}).(NexusRepositoryYumHostedOutput)
}

type NexusRepositoryYumHostedMapOutput struct{ *pulumi.OutputState }

func (NexusRepositoryYumHostedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusRepositoryYumHosted)(nil)).Elem()
}

func (o NexusRepositoryYumHostedMapOutput) ToNexusRepositoryYumHostedMapOutput() NexusRepositoryYumHostedMapOutput {
	return o
}

func (o NexusRepositoryYumHostedMapOutput) ToNexusRepositoryYumHostedMapOutputWithContext(ctx context.Context) NexusRepositoryYumHostedMapOutput {
	return o
}

func (o NexusRepositoryYumHostedMapOutput) MapIndex(k pulumi.StringInput) NexusRepositoryYumHostedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NexusRepositoryYumHosted {
		return vs[0].(map[string]*NexusRepositoryYumHosted)[vs[1].(string)]
	}).(NexusRepositoryYumHostedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRepositoryYumHostedInput)(nil)).Elem(), &NexusRepositoryYumHosted{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRepositoryYumHostedArrayInput)(nil)).Elem(), NexusRepositoryYumHostedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRepositoryYumHostedMapInput)(nil)).Elem(), NexusRepositoryYumHostedMap{})
	pulumi.RegisterOutputType(NexusRepositoryYumHostedOutput{})
	pulumi.RegisterOutputType(NexusRepositoryYumHostedArrayOutput{})
	pulumi.RegisterOutputType(NexusRepositoryYumHostedMapOutput{})
}
