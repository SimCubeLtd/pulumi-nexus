// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NexusBlobstoreGroup struct {
	pulumi.CustomResourceState

	// Available space in Bytes
	AvailableSpaceInBytes pulumi.IntOutput `pulumi:"availableSpaceInBytes"`
	// Count of blobs
	BlobCount pulumi.IntOutput `pulumi:"blobCount"`
	// The policy how to fill the members. Possible values: `roundRobin` or `writeToFirst`
	FillPolicy pulumi.StringOutput `pulumi:"fillPolicy"`
	// List of the names of blob stores that are members of this group
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Blobstore name
	Name pulumi.StringOutput `pulumi:"name"`
	// Soft quota of the blobstore
	SoftQuota NexusBlobstoreGroupSoftQuotaPtrOutput `pulumi:"softQuota"`
	// The total size of the blobstore in Bytes
	TotalSizeInBytes pulumi.IntOutput `pulumi:"totalSizeInBytes"`
}

// NewNexusBlobstoreGroup registers a new resource with the given unique name, arguments, and options.
func NewNexusBlobstoreGroup(ctx *pulumi.Context,
	name string, args *NexusBlobstoreGroupArgs, opts ...pulumi.ResourceOption) (*NexusBlobstoreGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FillPolicy == nil {
		return nil, errors.New("invalid value for required argument 'FillPolicy'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NexusBlobstoreGroup
	err := ctx.RegisterResource("nexus:index/nexusBlobstoreGroup:NexusBlobstoreGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNexusBlobstoreGroup gets an existing NexusBlobstoreGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNexusBlobstoreGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NexusBlobstoreGroupState, opts ...pulumi.ResourceOption) (*NexusBlobstoreGroup, error) {
	var resource NexusBlobstoreGroup
	err := ctx.ReadResource("nexus:index/nexusBlobstoreGroup:NexusBlobstoreGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NexusBlobstoreGroup resources.
type nexusBlobstoreGroupState struct {
	// Available space in Bytes
	AvailableSpaceInBytes *int `pulumi:"availableSpaceInBytes"`
	// Count of blobs
	BlobCount *int `pulumi:"blobCount"`
	// The policy how to fill the members. Possible values: `roundRobin` or `writeToFirst`
	FillPolicy *string `pulumi:"fillPolicy"`
	// List of the names of blob stores that are members of this group
	Members []string `pulumi:"members"`
	// Blobstore name
	Name *string `pulumi:"name"`
	// Soft quota of the blobstore
	SoftQuota *NexusBlobstoreGroupSoftQuota `pulumi:"softQuota"`
	// The total size of the blobstore in Bytes
	TotalSizeInBytes *int `pulumi:"totalSizeInBytes"`
}

type NexusBlobstoreGroupState struct {
	// Available space in Bytes
	AvailableSpaceInBytes pulumi.IntPtrInput
	// Count of blobs
	BlobCount pulumi.IntPtrInput
	// The policy how to fill the members. Possible values: `roundRobin` or `writeToFirst`
	FillPolicy pulumi.StringPtrInput
	// List of the names of blob stores that are members of this group
	Members pulumi.StringArrayInput
	// Blobstore name
	Name pulumi.StringPtrInput
	// Soft quota of the blobstore
	SoftQuota NexusBlobstoreGroupSoftQuotaPtrInput
	// The total size of the blobstore in Bytes
	TotalSizeInBytes pulumi.IntPtrInput
}

func (NexusBlobstoreGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusBlobstoreGroupState)(nil)).Elem()
}

type nexusBlobstoreGroupArgs struct {
	// The policy how to fill the members. Possible values: `roundRobin` or `writeToFirst`
	FillPolicy string `pulumi:"fillPolicy"`
	// List of the names of blob stores that are members of this group
	Members []string `pulumi:"members"`
	// Blobstore name
	Name *string `pulumi:"name"`
	// Soft quota of the blobstore
	SoftQuota *NexusBlobstoreGroupSoftQuota `pulumi:"softQuota"`
}

// The set of arguments for constructing a NexusBlobstoreGroup resource.
type NexusBlobstoreGroupArgs struct {
	// The policy how to fill the members. Possible values: `roundRobin` or `writeToFirst`
	FillPolicy pulumi.StringInput
	// List of the names of blob stores that are members of this group
	Members pulumi.StringArrayInput
	// Blobstore name
	Name pulumi.StringPtrInput
	// Soft quota of the blobstore
	SoftQuota NexusBlobstoreGroupSoftQuotaPtrInput
}

func (NexusBlobstoreGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusBlobstoreGroupArgs)(nil)).Elem()
}

type NexusBlobstoreGroupInput interface {
	pulumi.Input

	ToNexusBlobstoreGroupOutput() NexusBlobstoreGroupOutput
	ToNexusBlobstoreGroupOutputWithContext(ctx context.Context) NexusBlobstoreGroupOutput
}

func (*NexusBlobstoreGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusBlobstoreGroup)(nil)).Elem()
}

func (i *NexusBlobstoreGroup) ToNexusBlobstoreGroupOutput() NexusBlobstoreGroupOutput {
	return i.ToNexusBlobstoreGroupOutputWithContext(context.Background())
}

func (i *NexusBlobstoreGroup) ToNexusBlobstoreGroupOutputWithContext(ctx context.Context) NexusBlobstoreGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusBlobstoreGroupOutput)
}

// NexusBlobstoreGroupArrayInput is an input type that accepts NexusBlobstoreGroupArray and NexusBlobstoreGroupArrayOutput values.
// You can construct a concrete instance of `NexusBlobstoreGroupArrayInput` via:
//
//          NexusBlobstoreGroupArray{ NexusBlobstoreGroupArgs{...} }
type NexusBlobstoreGroupArrayInput interface {
	pulumi.Input

	ToNexusBlobstoreGroupArrayOutput() NexusBlobstoreGroupArrayOutput
	ToNexusBlobstoreGroupArrayOutputWithContext(context.Context) NexusBlobstoreGroupArrayOutput
}

type NexusBlobstoreGroupArray []NexusBlobstoreGroupInput

func (NexusBlobstoreGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusBlobstoreGroup)(nil)).Elem()
}

func (i NexusBlobstoreGroupArray) ToNexusBlobstoreGroupArrayOutput() NexusBlobstoreGroupArrayOutput {
	return i.ToNexusBlobstoreGroupArrayOutputWithContext(context.Background())
}

func (i NexusBlobstoreGroupArray) ToNexusBlobstoreGroupArrayOutputWithContext(ctx context.Context) NexusBlobstoreGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusBlobstoreGroupArrayOutput)
}

// NexusBlobstoreGroupMapInput is an input type that accepts NexusBlobstoreGroupMap and NexusBlobstoreGroupMapOutput values.
// You can construct a concrete instance of `NexusBlobstoreGroupMapInput` via:
//
//          NexusBlobstoreGroupMap{ "key": NexusBlobstoreGroupArgs{...} }
type NexusBlobstoreGroupMapInput interface {
	pulumi.Input

	ToNexusBlobstoreGroupMapOutput() NexusBlobstoreGroupMapOutput
	ToNexusBlobstoreGroupMapOutputWithContext(context.Context) NexusBlobstoreGroupMapOutput
}

type NexusBlobstoreGroupMap map[string]NexusBlobstoreGroupInput

func (NexusBlobstoreGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusBlobstoreGroup)(nil)).Elem()
}

func (i NexusBlobstoreGroupMap) ToNexusBlobstoreGroupMapOutput() NexusBlobstoreGroupMapOutput {
	return i.ToNexusBlobstoreGroupMapOutputWithContext(context.Background())
}

func (i NexusBlobstoreGroupMap) ToNexusBlobstoreGroupMapOutputWithContext(ctx context.Context) NexusBlobstoreGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusBlobstoreGroupMapOutput)
}

type NexusBlobstoreGroupOutput struct{ *pulumi.OutputState }

func (NexusBlobstoreGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusBlobstoreGroup)(nil)).Elem()
}

func (o NexusBlobstoreGroupOutput) ToNexusBlobstoreGroupOutput() NexusBlobstoreGroupOutput {
	return o
}

func (o NexusBlobstoreGroupOutput) ToNexusBlobstoreGroupOutputWithContext(ctx context.Context) NexusBlobstoreGroupOutput {
	return o
}

// Available space in Bytes
func (o NexusBlobstoreGroupOutput) AvailableSpaceInBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *NexusBlobstoreGroup) pulumi.IntOutput { return v.AvailableSpaceInBytes }).(pulumi.IntOutput)
}

// Count of blobs
func (o NexusBlobstoreGroupOutput) BlobCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NexusBlobstoreGroup) pulumi.IntOutput { return v.BlobCount }).(pulumi.IntOutput)
}

// The policy how to fill the members. Possible values: `roundRobin` or `writeToFirst`
func (o NexusBlobstoreGroupOutput) FillPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *NexusBlobstoreGroup) pulumi.StringOutput { return v.FillPolicy }).(pulumi.StringOutput)
}

// List of the names of blob stores that are members of this group
func (o NexusBlobstoreGroupOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NexusBlobstoreGroup) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// Blobstore name
func (o NexusBlobstoreGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NexusBlobstoreGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Soft quota of the blobstore
func (o NexusBlobstoreGroupOutput) SoftQuota() NexusBlobstoreGroupSoftQuotaPtrOutput {
	return o.ApplyT(func(v *NexusBlobstoreGroup) NexusBlobstoreGroupSoftQuotaPtrOutput { return v.SoftQuota }).(NexusBlobstoreGroupSoftQuotaPtrOutput)
}

// The total size of the blobstore in Bytes
func (o NexusBlobstoreGroupOutput) TotalSizeInBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *NexusBlobstoreGroup) pulumi.IntOutput { return v.TotalSizeInBytes }).(pulumi.IntOutput)
}

type NexusBlobstoreGroupArrayOutput struct{ *pulumi.OutputState }

func (NexusBlobstoreGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusBlobstoreGroup)(nil)).Elem()
}

func (o NexusBlobstoreGroupArrayOutput) ToNexusBlobstoreGroupArrayOutput() NexusBlobstoreGroupArrayOutput {
	return o
}

func (o NexusBlobstoreGroupArrayOutput) ToNexusBlobstoreGroupArrayOutputWithContext(ctx context.Context) NexusBlobstoreGroupArrayOutput {
	return o
}

func (o NexusBlobstoreGroupArrayOutput) Index(i pulumi.IntInput) NexusBlobstoreGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NexusBlobstoreGroup {
		return vs[0].([]*NexusBlobstoreGroup)[vs[1].(int)]
	}).(NexusBlobstoreGroupOutput)
}

type NexusBlobstoreGroupMapOutput struct{ *pulumi.OutputState }

func (NexusBlobstoreGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusBlobstoreGroup)(nil)).Elem()
}

func (o NexusBlobstoreGroupMapOutput) ToNexusBlobstoreGroupMapOutput() NexusBlobstoreGroupMapOutput {
	return o
}

func (o NexusBlobstoreGroupMapOutput) ToNexusBlobstoreGroupMapOutputWithContext(ctx context.Context) NexusBlobstoreGroupMapOutput {
	return o
}

func (o NexusBlobstoreGroupMapOutput) MapIndex(k pulumi.StringInput) NexusBlobstoreGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NexusBlobstoreGroup {
		return vs[0].(map[string]*NexusBlobstoreGroup)[vs[1].(string)]
	}).(NexusBlobstoreGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NexusBlobstoreGroupInput)(nil)).Elem(), &NexusBlobstoreGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusBlobstoreGroupArrayInput)(nil)).Elem(), NexusBlobstoreGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusBlobstoreGroupMapInput)(nil)).Elem(), NexusBlobstoreGroupMap{})
	pulumi.RegisterOutputType(NexusBlobstoreGroupOutput{})
	pulumi.RegisterOutputType(NexusBlobstoreGroupArrayOutput{})
	pulumi.RegisterOutputType(NexusBlobstoreGroupMapOutput{})
}
