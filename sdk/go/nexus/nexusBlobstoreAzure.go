// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NexusBlobstoreAzure struct {
	pulumi.CustomResourceState

	// Count of blobs
	BlobCount pulumi.IntOutput `pulumi:"blobCount"`
	// The Azure specific configuration details for the Azure object that'll contain the blob store
	BucketConfiguration NexusBlobstoreAzureBucketConfigurationOutput `pulumi:"bucketConfiguration"`
	// Blobstore name
	Name pulumi.StringOutput `pulumi:"name"`
	// Soft quota of the blobstore
	SoftQuota NexusBlobstoreAzureSoftQuotaPtrOutput `pulumi:"softQuota"`
	// The total size of the blobstore in Bytes
	TotalSizeInBytes pulumi.IntOutput `pulumi:"totalSizeInBytes"`
}

// NewNexusBlobstoreAzure registers a new resource with the given unique name, arguments, and options.
func NewNexusBlobstoreAzure(ctx *pulumi.Context,
	name string, args *NexusBlobstoreAzureArgs, opts ...pulumi.ResourceOption) (*NexusBlobstoreAzure, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'BucketConfiguration'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NexusBlobstoreAzure
	err := ctx.RegisterResource("nexus:index/nexusBlobstoreAzure:NexusBlobstoreAzure", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNexusBlobstoreAzure gets an existing NexusBlobstoreAzure resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNexusBlobstoreAzure(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NexusBlobstoreAzureState, opts ...pulumi.ResourceOption) (*NexusBlobstoreAzure, error) {
	var resource NexusBlobstoreAzure
	err := ctx.ReadResource("nexus:index/nexusBlobstoreAzure:NexusBlobstoreAzure", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NexusBlobstoreAzure resources.
type nexusBlobstoreAzureState struct {
	// Count of blobs
	BlobCount *int `pulumi:"blobCount"`
	// The Azure specific configuration details for the Azure object that'll contain the blob store
	BucketConfiguration *NexusBlobstoreAzureBucketConfiguration `pulumi:"bucketConfiguration"`
	// Blobstore name
	Name *string `pulumi:"name"`
	// Soft quota of the blobstore
	SoftQuota *NexusBlobstoreAzureSoftQuota `pulumi:"softQuota"`
	// The total size of the blobstore in Bytes
	TotalSizeInBytes *int `pulumi:"totalSizeInBytes"`
}

type NexusBlobstoreAzureState struct {
	// Count of blobs
	BlobCount pulumi.IntPtrInput
	// The Azure specific configuration details for the Azure object that'll contain the blob store
	BucketConfiguration NexusBlobstoreAzureBucketConfigurationPtrInput
	// Blobstore name
	Name pulumi.StringPtrInput
	// Soft quota of the blobstore
	SoftQuota NexusBlobstoreAzureSoftQuotaPtrInput
	// The total size of the blobstore in Bytes
	TotalSizeInBytes pulumi.IntPtrInput
}

func (NexusBlobstoreAzureState) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusBlobstoreAzureState)(nil)).Elem()
}

type nexusBlobstoreAzureArgs struct {
	// The Azure specific configuration details for the Azure object that'll contain the blob store
	BucketConfiguration NexusBlobstoreAzureBucketConfiguration `pulumi:"bucketConfiguration"`
	// Blobstore name
	Name *string `pulumi:"name"`
	// Soft quota of the blobstore
	SoftQuota *NexusBlobstoreAzureSoftQuota `pulumi:"softQuota"`
}

// The set of arguments for constructing a NexusBlobstoreAzure resource.
type NexusBlobstoreAzureArgs struct {
	// The Azure specific configuration details for the Azure object that'll contain the blob store
	BucketConfiguration NexusBlobstoreAzureBucketConfigurationInput
	// Blobstore name
	Name pulumi.StringPtrInput
	// Soft quota of the blobstore
	SoftQuota NexusBlobstoreAzureSoftQuotaPtrInput
}

func (NexusBlobstoreAzureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusBlobstoreAzureArgs)(nil)).Elem()
}

type NexusBlobstoreAzureInput interface {
	pulumi.Input

	ToNexusBlobstoreAzureOutput() NexusBlobstoreAzureOutput
	ToNexusBlobstoreAzureOutputWithContext(ctx context.Context) NexusBlobstoreAzureOutput
}

func (*NexusBlobstoreAzure) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusBlobstoreAzure)(nil)).Elem()
}

func (i *NexusBlobstoreAzure) ToNexusBlobstoreAzureOutput() NexusBlobstoreAzureOutput {
	return i.ToNexusBlobstoreAzureOutputWithContext(context.Background())
}

func (i *NexusBlobstoreAzure) ToNexusBlobstoreAzureOutputWithContext(ctx context.Context) NexusBlobstoreAzureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusBlobstoreAzureOutput)
}

// NexusBlobstoreAzureArrayInput is an input type that accepts NexusBlobstoreAzureArray and NexusBlobstoreAzureArrayOutput values.
// You can construct a concrete instance of `NexusBlobstoreAzureArrayInput` via:
//
//          NexusBlobstoreAzureArray{ NexusBlobstoreAzureArgs{...} }
type NexusBlobstoreAzureArrayInput interface {
	pulumi.Input

	ToNexusBlobstoreAzureArrayOutput() NexusBlobstoreAzureArrayOutput
	ToNexusBlobstoreAzureArrayOutputWithContext(context.Context) NexusBlobstoreAzureArrayOutput
}

type NexusBlobstoreAzureArray []NexusBlobstoreAzureInput

func (NexusBlobstoreAzureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusBlobstoreAzure)(nil)).Elem()
}

func (i NexusBlobstoreAzureArray) ToNexusBlobstoreAzureArrayOutput() NexusBlobstoreAzureArrayOutput {
	return i.ToNexusBlobstoreAzureArrayOutputWithContext(context.Background())
}

func (i NexusBlobstoreAzureArray) ToNexusBlobstoreAzureArrayOutputWithContext(ctx context.Context) NexusBlobstoreAzureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusBlobstoreAzureArrayOutput)
}

// NexusBlobstoreAzureMapInput is an input type that accepts NexusBlobstoreAzureMap and NexusBlobstoreAzureMapOutput values.
// You can construct a concrete instance of `NexusBlobstoreAzureMapInput` via:
//
//          NexusBlobstoreAzureMap{ "key": NexusBlobstoreAzureArgs{...} }
type NexusBlobstoreAzureMapInput interface {
	pulumi.Input

	ToNexusBlobstoreAzureMapOutput() NexusBlobstoreAzureMapOutput
	ToNexusBlobstoreAzureMapOutputWithContext(context.Context) NexusBlobstoreAzureMapOutput
}

type NexusBlobstoreAzureMap map[string]NexusBlobstoreAzureInput

func (NexusBlobstoreAzureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusBlobstoreAzure)(nil)).Elem()
}

func (i NexusBlobstoreAzureMap) ToNexusBlobstoreAzureMapOutput() NexusBlobstoreAzureMapOutput {
	return i.ToNexusBlobstoreAzureMapOutputWithContext(context.Background())
}

func (i NexusBlobstoreAzureMap) ToNexusBlobstoreAzureMapOutputWithContext(ctx context.Context) NexusBlobstoreAzureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusBlobstoreAzureMapOutput)
}

type NexusBlobstoreAzureOutput struct{ *pulumi.OutputState }

func (NexusBlobstoreAzureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusBlobstoreAzure)(nil)).Elem()
}

func (o NexusBlobstoreAzureOutput) ToNexusBlobstoreAzureOutput() NexusBlobstoreAzureOutput {
	return o
}

func (o NexusBlobstoreAzureOutput) ToNexusBlobstoreAzureOutputWithContext(ctx context.Context) NexusBlobstoreAzureOutput {
	return o
}

// Count of blobs
func (o NexusBlobstoreAzureOutput) BlobCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NexusBlobstoreAzure) pulumi.IntOutput { return v.BlobCount }).(pulumi.IntOutput)
}

// The Azure specific configuration details for the Azure object that'll contain the blob store
func (o NexusBlobstoreAzureOutput) BucketConfiguration() NexusBlobstoreAzureBucketConfigurationOutput {
	return o.ApplyT(func(v *NexusBlobstoreAzure) NexusBlobstoreAzureBucketConfigurationOutput {
		return v.BucketConfiguration
	}).(NexusBlobstoreAzureBucketConfigurationOutput)
}

// Blobstore name
func (o NexusBlobstoreAzureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NexusBlobstoreAzure) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Soft quota of the blobstore
func (o NexusBlobstoreAzureOutput) SoftQuota() NexusBlobstoreAzureSoftQuotaPtrOutput {
	return o.ApplyT(func(v *NexusBlobstoreAzure) NexusBlobstoreAzureSoftQuotaPtrOutput { return v.SoftQuota }).(NexusBlobstoreAzureSoftQuotaPtrOutput)
}

// The total size of the blobstore in Bytes
func (o NexusBlobstoreAzureOutput) TotalSizeInBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *NexusBlobstoreAzure) pulumi.IntOutput { return v.TotalSizeInBytes }).(pulumi.IntOutput)
}

type NexusBlobstoreAzureArrayOutput struct{ *pulumi.OutputState }

func (NexusBlobstoreAzureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusBlobstoreAzure)(nil)).Elem()
}

func (o NexusBlobstoreAzureArrayOutput) ToNexusBlobstoreAzureArrayOutput() NexusBlobstoreAzureArrayOutput {
	return o
}

func (o NexusBlobstoreAzureArrayOutput) ToNexusBlobstoreAzureArrayOutputWithContext(ctx context.Context) NexusBlobstoreAzureArrayOutput {
	return o
}

func (o NexusBlobstoreAzureArrayOutput) Index(i pulumi.IntInput) NexusBlobstoreAzureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NexusBlobstoreAzure {
		return vs[0].([]*NexusBlobstoreAzure)[vs[1].(int)]
	}).(NexusBlobstoreAzureOutput)
}

type NexusBlobstoreAzureMapOutput struct{ *pulumi.OutputState }

func (NexusBlobstoreAzureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusBlobstoreAzure)(nil)).Elem()
}

func (o NexusBlobstoreAzureMapOutput) ToNexusBlobstoreAzureMapOutput() NexusBlobstoreAzureMapOutput {
	return o
}

func (o NexusBlobstoreAzureMapOutput) ToNexusBlobstoreAzureMapOutputWithContext(ctx context.Context) NexusBlobstoreAzureMapOutput {
	return o
}

func (o NexusBlobstoreAzureMapOutput) MapIndex(k pulumi.StringInput) NexusBlobstoreAzureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NexusBlobstoreAzure {
		return vs[0].(map[string]*NexusBlobstoreAzure)[vs[1].(string)]
	}).(NexusBlobstoreAzureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NexusBlobstoreAzureInput)(nil)).Elem(), &NexusBlobstoreAzure{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusBlobstoreAzureArrayInput)(nil)).Elem(), NexusBlobstoreAzureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusBlobstoreAzureMapInput)(nil)).Elem(), NexusBlobstoreAzureMap{})
	pulumi.RegisterOutputType(NexusBlobstoreAzureOutput{})
	pulumi.RegisterOutputType(NexusBlobstoreAzureArrayOutput{})
	pulumi.RegisterOutputType(NexusBlobstoreAzureMapOutput{})
}
