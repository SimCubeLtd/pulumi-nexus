// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NexusRoutingRule struct {
	pulumi.CustomResourceState

	// The description of the routing rule
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Matchers is a list of regular expressions used to identify request paths that are allowed or blocked (depending on above
	// mode)
	Matchers pulumi.StringArrayOutput `pulumi:"matchers"`
	// The mode describe how to hande with mathing requests. Possible values: `BLOCK` or `ALLOW` Default: `BLOCK`
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name of the routing rule
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewNexusRoutingRule registers a new resource with the given unique name, arguments, and options.
func NewNexusRoutingRule(ctx *pulumi.Context,
	name string, args *NexusRoutingRuleArgs, opts ...pulumi.ResourceOption) (*NexusRoutingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Matchers == nil {
		return nil, errors.New("invalid value for required argument 'Matchers'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NexusRoutingRule
	err := ctx.RegisterResource("nexus:index/nexusRoutingRule:NexusRoutingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNexusRoutingRule gets an existing NexusRoutingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNexusRoutingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NexusRoutingRuleState, opts ...pulumi.ResourceOption) (*NexusRoutingRule, error) {
	var resource NexusRoutingRule
	err := ctx.ReadResource("nexus:index/nexusRoutingRule:NexusRoutingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NexusRoutingRule resources.
type nexusRoutingRuleState struct {
	// The description of the routing rule
	Description *string `pulumi:"description"`
	// Matchers is a list of regular expressions used to identify request paths that are allowed or blocked (depending on above
	// mode)
	Matchers []string `pulumi:"matchers"`
	// The mode describe how to hande with mathing requests. Possible values: `BLOCK` or `ALLOW` Default: `BLOCK`
	Mode *string `pulumi:"mode"`
	// The name of the routing rule
	Name *string `pulumi:"name"`
}

type NexusRoutingRuleState struct {
	// The description of the routing rule
	Description pulumi.StringPtrInput
	// Matchers is a list of regular expressions used to identify request paths that are allowed or blocked (depending on above
	// mode)
	Matchers pulumi.StringArrayInput
	// The mode describe how to hande with mathing requests. Possible values: `BLOCK` or `ALLOW` Default: `BLOCK`
	Mode pulumi.StringPtrInput
	// The name of the routing rule
	Name pulumi.StringPtrInput
}

func (NexusRoutingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusRoutingRuleState)(nil)).Elem()
}

type nexusRoutingRuleArgs struct {
	// The description of the routing rule
	Description *string `pulumi:"description"`
	// Matchers is a list of regular expressions used to identify request paths that are allowed or blocked (depending on above
	// mode)
	Matchers []string `pulumi:"matchers"`
	// The mode describe how to hande with mathing requests. Possible values: `BLOCK` or `ALLOW` Default: `BLOCK`
	Mode *string `pulumi:"mode"`
	// The name of the routing rule
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a NexusRoutingRule resource.
type NexusRoutingRuleArgs struct {
	// The description of the routing rule
	Description pulumi.StringPtrInput
	// Matchers is a list of regular expressions used to identify request paths that are allowed or blocked (depending on above
	// mode)
	Matchers pulumi.StringArrayInput
	// The mode describe how to hande with mathing requests. Possible values: `BLOCK` or `ALLOW` Default: `BLOCK`
	Mode pulumi.StringPtrInput
	// The name of the routing rule
	Name pulumi.StringPtrInput
}

func (NexusRoutingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nexusRoutingRuleArgs)(nil)).Elem()
}

type NexusRoutingRuleInput interface {
	pulumi.Input

	ToNexusRoutingRuleOutput() NexusRoutingRuleOutput
	ToNexusRoutingRuleOutputWithContext(ctx context.Context) NexusRoutingRuleOutput
}

func (*NexusRoutingRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusRoutingRule)(nil)).Elem()
}

func (i *NexusRoutingRule) ToNexusRoutingRuleOutput() NexusRoutingRuleOutput {
	return i.ToNexusRoutingRuleOutputWithContext(context.Background())
}

func (i *NexusRoutingRule) ToNexusRoutingRuleOutputWithContext(ctx context.Context) NexusRoutingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRoutingRuleOutput)
}

// NexusRoutingRuleArrayInput is an input type that accepts NexusRoutingRuleArray and NexusRoutingRuleArrayOutput values.
// You can construct a concrete instance of `NexusRoutingRuleArrayInput` via:
//
//          NexusRoutingRuleArray{ NexusRoutingRuleArgs{...} }
type NexusRoutingRuleArrayInput interface {
	pulumi.Input

	ToNexusRoutingRuleArrayOutput() NexusRoutingRuleArrayOutput
	ToNexusRoutingRuleArrayOutputWithContext(context.Context) NexusRoutingRuleArrayOutput
}

type NexusRoutingRuleArray []NexusRoutingRuleInput

func (NexusRoutingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusRoutingRule)(nil)).Elem()
}

func (i NexusRoutingRuleArray) ToNexusRoutingRuleArrayOutput() NexusRoutingRuleArrayOutput {
	return i.ToNexusRoutingRuleArrayOutputWithContext(context.Background())
}

func (i NexusRoutingRuleArray) ToNexusRoutingRuleArrayOutputWithContext(ctx context.Context) NexusRoutingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRoutingRuleArrayOutput)
}

// NexusRoutingRuleMapInput is an input type that accepts NexusRoutingRuleMap and NexusRoutingRuleMapOutput values.
// You can construct a concrete instance of `NexusRoutingRuleMapInput` via:
//
//          NexusRoutingRuleMap{ "key": NexusRoutingRuleArgs{...} }
type NexusRoutingRuleMapInput interface {
	pulumi.Input

	ToNexusRoutingRuleMapOutput() NexusRoutingRuleMapOutput
	ToNexusRoutingRuleMapOutputWithContext(context.Context) NexusRoutingRuleMapOutput
}

type NexusRoutingRuleMap map[string]NexusRoutingRuleInput

func (NexusRoutingRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusRoutingRule)(nil)).Elem()
}

func (i NexusRoutingRuleMap) ToNexusRoutingRuleMapOutput() NexusRoutingRuleMapOutput {
	return i.ToNexusRoutingRuleMapOutputWithContext(context.Background())
}

func (i NexusRoutingRuleMap) ToNexusRoutingRuleMapOutputWithContext(ctx context.Context) NexusRoutingRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NexusRoutingRuleMapOutput)
}

type NexusRoutingRuleOutput struct{ *pulumi.OutputState }

func (NexusRoutingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NexusRoutingRule)(nil)).Elem()
}

func (o NexusRoutingRuleOutput) ToNexusRoutingRuleOutput() NexusRoutingRuleOutput {
	return o
}

func (o NexusRoutingRuleOutput) ToNexusRoutingRuleOutputWithContext(ctx context.Context) NexusRoutingRuleOutput {
	return o
}

// The description of the routing rule
func (o NexusRoutingRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NexusRoutingRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Matchers is a list of regular expressions used to identify request paths that are allowed or blocked (depending on above
// mode)
func (o NexusRoutingRuleOutput) Matchers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NexusRoutingRule) pulumi.StringArrayOutput { return v.Matchers }).(pulumi.StringArrayOutput)
}

// The mode describe how to hande with mathing requests. Possible values: `BLOCK` or `ALLOW` Default: `BLOCK`
func (o NexusRoutingRuleOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NexusRoutingRule) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the routing rule
func (o NexusRoutingRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NexusRoutingRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type NexusRoutingRuleArrayOutput struct{ *pulumi.OutputState }

func (NexusRoutingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NexusRoutingRule)(nil)).Elem()
}

func (o NexusRoutingRuleArrayOutput) ToNexusRoutingRuleArrayOutput() NexusRoutingRuleArrayOutput {
	return o
}

func (o NexusRoutingRuleArrayOutput) ToNexusRoutingRuleArrayOutputWithContext(ctx context.Context) NexusRoutingRuleArrayOutput {
	return o
}

func (o NexusRoutingRuleArrayOutput) Index(i pulumi.IntInput) NexusRoutingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NexusRoutingRule {
		return vs[0].([]*NexusRoutingRule)[vs[1].(int)]
	}).(NexusRoutingRuleOutput)
}

type NexusRoutingRuleMapOutput struct{ *pulumi.OutputState }

func (NexusRoutingRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NexusRoutingRule)(nil)).Elem()
}

func (o NexusRoutingRuleMapOutput) ToNexusRoutingRuleMapOutput() NexusRoutingRuleMapOutput {
	return o
}

func (o NexusRoutingRuleMapOutput) ToNexusRoutingRuleMapOutputWithContext(ctx context.Context) NexusRoutingRuleMapOutput {
	return o
}

func (o NexusRoutingRuleMapOutput) MapIndex(k pulumi.StringInput) NexusRoutingRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NexusRoutingRule {
		return vs[0].(map[string]*NexusRoutingRule)[vs[1].(string)]
	}).(NexusRoutingRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRoutingRuleInput)(nil)).Elem(), &NexusRoutingRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRoutingRuleArrayInput)(nil)).Elem(), NexusRoutingRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NexusRoutingRuleMapInput)(nil)).Elem(), NexusRoutingRuleMap{})
	pulumi.RegisterOutputType(NexusRoutingRuleOutput{})
	pulumi.RegisterOutputType(NexusRoutingRuleArrayOutput{})
	pulumi.RegisterOutputType(NexusRoutingRuleMapOutput{})
}
