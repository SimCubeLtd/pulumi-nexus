// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nexus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetItemSection struct {
	Fields []GetItemSectionField `pulumi:"fields"`
	Id     string                `pulumi:"id"`
	Label  string                `pulumi:"label"`
}

// GetItemSectionInput is an input type that accepts GetItemSectionArgs and GetItemSectionOutput values.
// You can construct a concrete instance of `GetItemSectionInput` via:
//
//          GetItemSectionArgs{...}
type GetItemSectionInput interface {
	pulumi.Input

	ToGetItemSectionOutput() GetItemSectionOutput
	ToGetItemSectionOutputWithContext(context.Context) GetItemSectionOutput
}

type GetItemSectionArgs struct {
	Fields GetItemSectionFieldArrayInput `pulumi:"fields"`
	Id     pulumi.StringInput            `pulumi:"id"`
	Label  pulumi.StringInput            `pulumi:"label"`
}

func (GetItemSectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetItemSection)(nil)).Elem()
}

func (i GetItemSectionArgs) ToGetItemSectionOutput() GetItemSectionOutput {
	return i.ToGetItemSectionOutputWithContext(context.Background())
}

func (i GetItemSectionArgs) ToGetItemSectionOutputWithContext(ctx context.Context) GetItemSectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetItemSectionOutput)
}

// GetItemSectionArrayInput is an input type that accepts GetItemSectionArray and GetItemSectionArrayOutput values.
// You can construct a concrete instance of `GetItemSectionArrayInput` via:
//
//          GetItemSectionArray{ GetItemSectionArgs{...} }
type GetItemSectionArrayInput interface {
	pulumi.Input

	ToGetItemSectionArrayOutput() GetItemSectionArrayOutput
	ToGetItemSectionArrayOutputWithContext(context.Context) GetItemSectionArrayOutput
}

type GetItemSectionArray []GetItemSectionInput

func (GetItemSectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetItemSection)(nil)).Elem()
}

func (i GetItemSectionArray) ToGetItemSectionArrayOutput() GetItemSectionArrayOutput {
	return i.ToGetItemSectionArrayOutputWithContext(context.Background())
}

func (i GetItemSectionArray) ToGetItemSectionArrayOutputWithContext(ctx context.Context) GetItemSectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetItemSectionArrayOutput)
}

type GetItemSectionOutput struct{ *pulumi.OutputState }

func (GetItemSectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetItemSection)(nil)).Elem()
}

func (o GetItemSectionOutput) ToGetItemSectionOutput() GetItemSectionOutput {
	return o
}

func (o GetItemSectionOutput) ToGetItemSectionOutputWithContext(ctx context.Context) GetItemSectionOutput {
	return o
}

func (o GetItemSectionOutput) Fields() GetItemSectionFieldArrayOutput {
	return o.ApplyT(func(v GetItemSection) []GetItemSectionField { return v.Fields }).(GetItemSectionFieldArrayOutput)
}

func (o GetItemSectionOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetItemSection) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetItemSectionOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v GetItemSection) string { return v.Label }).(pulumi.StringOutput)
}

type GetItemSectionArrayOutput struct{ *pulumi.OutputState }

func (GetItemSectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetItemSection)(nil)).Elem()
}

func (o GetItemSectionArrayOutput) ToGetItemSectionArrayOutput() GetItemSectionArrayOutput {
	return o
}

func (o GetItemSectionArrayOutput) ToGetItemSectionArrayOutputWithContext(ctx context.Context) GetItemSectionArrayOutput {
	return o
}

func (o GetItemSectionArrayOutput) Index(i pulumi.IntInput) GetItemSectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetItemSection {
		return vs[0].([]GetItemSection)[vs[1].(int)]
	}).(GetItemSectionOutput)
}

type GetItemSectionField struct {
	Id      string `pulumi:"id"`
	Label   string `pulumi:"label"`
	Purpose string `pulumi:"purpose"`
	// (Only applies to the database category) The type of database. One of ["db2" "filemaker" "msaccess" "mssql" "mysql" "oracle" "postgresql" "sqlite" "other"]
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

// GetItemSectionFieldInput is an input type that accepts GetItemSectionFieldArgs and GetItemSectionFieldOutput values.
// You can construct a concrete instance of `GetItemSectionFieldInput` via:
//
//          GetItemSectionFieldArgs{...}
type GetItemSectionFieldInput interface {
	pulumi.Input

	ToGetItemSectionFieldOutput() GetItemSectionFieldOutput
	ToGetItemSectionFieldOutputWithContext(context.Context) GetItemSectionFieldOutput
}

type GetItemSectionFieldArgs struct {
	Id      pulumi.StringInput `pulumi:"id"`
	Label   pulumi.StringInput `pulumi:"label"`
	Purpose pulumi.StringInput `pulumi:"purpose"`
	// (Only applies to the database category) The type of database. One of ["db2" "filemaker" "msaccess" "mssql" "mysql" "oracle" "postgresql" "sqlite" "other"]
	Type  pulumi.StringInput `pulumi:"type"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (GetItemSectionFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetItemSectionField)(nil)).Elem()
}

func (i GetItemSectionFieldArgs) ToGetItemSectionFieldOutput() GetItemSectionFieldOutput {
	return i.ToGetItemSectionFieldOutputWithContext(context.Background())
}

func (i GetItemSectionFieldArgs) ToGetItemSectionFieldOutputWithContext(ctx context.Context) GetItemSectionFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetItemSectionFieldOutput)
}

// GetItemSectionFieldArrayInput is an input type that accepts GetItemSectionFieldArray and GetItemSectionFieldArrayOutput values.
// You can construct a concrete instance of `GetItemSectionFieldArrayInput` via:
//
//          GetItemSectionFieldArray{ GetItemSectionFieldArgs{...} }
type GetItemSectionFieldArrayInput interface {
	pulumi.Input

	ToGetItemSectionFieldArrayOutput() GetItemSectionFieldArrayOutput
	ToGetItemSectionFieldArrayOutputWithContext(context.Context) GetItemSectionFieldArrayOutput
}

type GetItemSectionFieldArray []GetItemSectionFieldInput

func (GetItemSectionFieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetItemSectionField)(nil)).Elem()
}

func (i GetItemSectionFieldArray) ToGetItemSectionFieldArrayOutput() GetItemSectionFieldArrayOutput {
	return i.ToGetItemSectionFieldArrayOutputWithContext(context.Background())
}

func (i GetItemSectionFieldArray) ToGetItemSectionFieldArrayOutputWithContext(ctx context.Context) GetItemSectionFieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetItemSectionFieldArrayOutput)
}

type GetItemSectionFieldOutput struct{ *pulumi.OutputState }

func (GetItemSectionFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetItemSectionField)(nil)).Elem()
}

func (o GetItemSectionFieldOutput) ToGetItemSectionFieldOutput() GetItemSectionFieldOutput {
	return o
}

func (o GetItemSectionFieldOutput) ToGetItemSectionFieldOutputWithContext(ctx context.Context) GetItemSectionFieldOutput {
	return o
}

func (o GetItemSectionFieldOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetItemSectionField) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetItemSectionFieldOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v GetItemSectionField) string { return v.Label }).(pulumi.StringOutput)
}

func (o GetItemSectionFieldOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v GetItemSectionField) string { return v.Purpose }).(pulumi.StringOutput)
}

// (Only applies to the database category) The type of database. One of ["db2" "filemaker" "msaccess" "mssql" "mysql" "oracle" "postgresql" "sqlite" "other"]
func (o GetItemSectionFieldOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetItemSectionField) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetItemSectionFieldOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetItemSectionField) string { return v.Value }).(pulumi.StringOutput)
}

type GetItemSectionFieldArrayOutput struct{ *pulumi.OutputState }

func (GetItemSectionFieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetItemSectionField)(nil)).Elem()
}

func (o GetItemSectionFieldArrayOutput) ToGetItemSectionFieldArrayOutput() GetItemSectionFieldArrayOutput {
	return o
}

func (o GetItemSectionFieldArrayOutput) ToGetItemSectionFieldArrayOutputWithContext(ctx context.Context) GetItemSectionFieldArrayOutput {
	return o
}

func (o GetItemSectionFieldArrayOutput) Index(i pulumi.IntInput) GetItemSectionFieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetItemSectionField {
		return vs[0].([]GetItemSectionField)[vs[1].(int)]
	}).(GetItemSectionFieldOutput)
}

type ItemPasswordRecipe struct {
	// Use digits [0-9] when generating the password.
	Digits *bool `pulumi:"digits"`
	// The length of the password to be generated.
	Length *int `pulumi:"length"`
	// Use letters [a-zA-Z] when generating the password.
	Letters *bool `pulumi:"letters"`
	// Use symbols [!@.-_*] when generating the password.
	Symbols *bool `pulumi:"symbols"`
}

// ItemPasswordRecipeInput is an input type that accepts ItemPasswordRecipeArgs and ItemPasswordRecipeOutput values.
// You can construct a concrete instance of `ItemPasswordRecipeInput` via:
//
//          ItemPasswordRecipeArgs{...}
type ItemPasswordRecipeInput interface {
	pulumi.Input

	ToItemPasswordRecipeOutput() ItemPasswordRecipeOutput
	ToItemPasswordRecipeOutputWithContext(context.Context) ItemPasswordRecipeOutput
}

type ItemPasswordRecipeArgs struct {
	// Use digits [0-9] when generating the password.
	Digits pulumi.BoolPtrInput `pulumi:"digits"`
	// The length of the password to be generated.
	Length pulumi.IntPtrInput `pulumi:"length"`
	// Use letters [a-zA-Z] when generating the password.
	Letters pulumi.BoolPtrInput `pulumi:"letters"`
	// Use symbols [!@.-_*] when generating the password.
	Symbols pulumi.BoolPtrInput `pulumi:"symbols"`
}

func (ItemPasswordRecipeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemPasswordRecipe)(nil)).Elem()
}

func (i ItemPasswordRecipeArgs) ToItemPasswordRecipeOutput() ItemPasswordRecipeOutput {
	return i.ToItemPasswordRecipeOutputWithContext(context.Background())
}

func (i ItemPasswordRecipeArgs) ToItemPasswordRecipeOutputWithContext(ctx context.Context) ItemPasswordRecipeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemPasswordRecipeOutput)
}

func (i ItemPasswordRecipeArgs) ToItemPasswordRecipePtrOutput() ItemPasswordRecipePtrOutput {
	return i.ToItemPasswordRecipePtrOutputWithContext(context.Background())
}

func (i ItemPasswordRecipeArgs) ToItemPasswordRecipePtrOutputWithContext(ctx context.Context) ItemPasswordRecipePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemPasswordRecipeOutput).ToItemPasswordRecipePtrOutputWithContext(ctx)
}

// ItemPasswordRecipePtrInput is an input type that accepts ItemPasswordRecipeArgs, ItemPasswordRecipePtr and ItemPasswordRecipePtrOutput values.
// You can construct a concrete instance of `ItemPasswordRecipePtrInput` via:
//
//          ItemPasswordRecipeArgs{...}
//
//  or:
//
//          nil
type ItemPasswordRecipePtrInput interface {
	pulumi.Input

	ToItemPasswordRecipePtrOutput() ItemPasswordRecipePtrOutput
	ToItemPasswordRecipePtrOutputWithContext(context.Context) ItemPasswordRecipePtrOutput
}

type itemPasswordRecipePtrType ItemPasswordRecipeArgs

func ItemPasswordRecipePtr(v *ItemPasswordRecipeArgs) ItemPasswordRecipePtrInput {
	return (*itemPasswordRecipePtrType)(v)
}

func (*itemPasswordRecipePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ItemPasswordRecipe)(nil)).Elem()
}

func (i *itemPasswordRecipePtrType) ToItemPasswordRecipePtrOutput() ItemPasswordRecipePtrOutput {
	return i.ToItemPasswordRecipePtrOutputWithContext(context.Background())
}

func (i *itemPasswordRecipePtrType) ToItemPasswordRecipePtrOutputWithContext(ctx context.Context) ItemPasswordRecipePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemPasswordRecipePtrOutput)
}

type ItemPasswordRecipeOutput struct{ *pulumi.OutputState }

func (ItemPasswordRecipeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemPasswordRecipe)(nil)).Elem()
}

func (o ItemPasswordRecipeOutput) ToItemPasswordRecipeOutput() ItemPasswordRecipeOutput {
	return o
}

func (o ItemPasswordRecipeOutput) ToItemPasswordRecipeOutputWithContext(ctx context.Context) ItemPasswordRecipeOutput {
	return o
}

func (o ItemPasswordRecipeOutput) ToItemPasswordRecipePtrOutput() ItemPasswordRecipePtrOutput {
	return o.ToItemPasswordRecipePtrOutputWithContext(context.Background())
}

func (o ItemPasswordRecipeOutput) ToItemPasswordRecipePtrOutputWithContext(ctx context.Context) ItemPasswordRecipePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ItemPasswordRecipe) *ItemPasswordRecipe {
		return &v
	}).(ItemPasswordRecipePtrOutput)
}

// Use digits [0-9] when generating the password.
func (o ItemPasswordRecipeOutput) Digits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ItemPasswordRecipe) *bool { return v.Digits }).(pulumi.BoolPtrOutput)
}

// The length of the password to be generated.
func (o ItemPasswordRecipeOutput) Length() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ItemPasswordRecipe) *int { return v.Length }).(pulumi.IntPtrOutput)
}

// Use letters [a-zA-Z] when generating the password.
func (o ItemPasswordRecipeOutput) Letters() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ItemPasswordRecipe) *bool { return v.Letters }).(pulumi.BoolPtrOutput)
}

// Use symbols [!@.-_*] when generating the password.
func (o ItemPasswordRecipeOutput) Symbols() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ItemPasswordRecipe) *bool { return v.Symbols }).(pulumi.BoolPtrOutput)
}

type ItemPasswordRecipePtrOutput struct{ *pulumi.OutputState }

func (ItemPasswordRecipePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ItemPasswordRecipe)(nil)).Elem()
}

func (o ItemPasswordRecipePtrOutput) ToItemPasswordRecipePtrOutput() ItemPasswordRecipePtrOutput {
	return o
}

func (o ItemPasswordRecipePtrOutput) ToItemPasswordRecipePtrOutputWithContext(ctx context.Context) ItemPasswordRecipePtrOutput {
	return o
}

func (o ItemPasswordRecipePtrOutput) Elem() ItemPasswordRecipeOutput {
	return o.ApplyT(func(v *ItemPasswordRecipe) ItemPasswordRecipe {
		if v != nil {
			return *v
		}
		var ret ItemPasswordRecipe
		return ret
	}).(ItemPasswordRecipeOutput)
}

// Use digits [0-9] when generating the password.
func (o ItemPasswordRecipePtrOutput) Digits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ItemPasswordRecipe) *bool {
		if v == nil {
			return nil
		}
		return v.Digits
	}).(pulumi.BoolPtrOutput)
}

// The length of the password to be generated.
func (o ItemPasswordRecipePtrOutput) Length() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ItemPasswordRecipe) *int {
		if v == nil {
			return nil
		}
		return v.Length
	}).(pulumi.IntPtrOutput)
}

// Use letters [a-zA-Z] when generating the password.
func (o ItemPasswordRecipePtrOutput) Letters() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ItemPasswordRecipe) *bool {
		if v == nil {
			return nil
		}
		return v.Letters
	}).(pulumi.BoolPtrOutput)
}

// Use symbols [!@.-_*] when generating the password.
func (o ItemPasswordRecipePtrOutput) Symbols() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ItemPasswordRecipe) *bool {
		if v == nil {
			return nil
		}
		return v.Symbols
	}).(pulumi.BoolPtrOutput)
}

type ItemSection struct {
	// A list of custom fields in the section.
	Fields []ItemSectionField `pulumi:"fields"`
	// A unique identifier for the section.
	Id *string `pulumi:"id"`
	// The label for the section.
	Label string `pulumi:"label"`
}

// ItemSectionInput is an input type that accepts ItemSectionArgs and ItemSectionOutput values.
// You can construct a concrete instance of `ItemSectionInput` via:
//
//          ItemSectionArgs{...}
type ItemSectionInput interface {
	pulumi.Input

	ToItemSectionOutput() ItemSectionOutput
	ToItemSectionOutputWithContext(context.Context) ItemSectionOutput
}

type ItemSectionArgs struct {
	// A list of custom fields in the section.
	Fields ItemSectionFieldArrayInput `pulumi:"fields"`
	// A unique identifier for the section.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The label for the section.
	Label pulumi.StringInput `pulumi:"label"`
}

func (ItemSectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemSection)(nil)).Elem()
}

func (i ItemSectionArgs) ToItemSectionOutput() ItemSectionOutput {
	return i.ToItemSectionOutputWithContext(context.Background())
}

func (i ItemSectionArgs) ToItemSectionOutputWithContext(ctx context.Context) ItemSectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemSectionOutput)
}

// ItemSectionArrayInput is an input type that accepts ItemSectionArray and ItemSectionArrayOutput values.
// You can construct a concrete instance of `ItemSectionArrayInput` via:
//
//          ItemSectionArray{ ItemSectionArgs{...} }
type ItemSectionArrayInput interface {
	pulumi.Input

	ToItemSectionArrayOutput() ItemSectionArrayOutput
	ToItemSectionArrayOutputWithContext(context.Context) ItemSectionArrayOutput
}

type ItemSectionArray []ItemSectionInput

func (ItemSectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ItemSection)(nil)).Elem()
}

func (i ItemSectionArray) ToItemSectionArrayOutput() ItemSectionArrayOutput {
	return i.ToItemSectionArrayOutputWithContext(context.Background())
}

func (i ItemSectionArray) ToItemSectionArrayOutputWithContext(ctx context.Context) ItemSectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemSectionArrayOutput)
}

type ItemSectionOutput struct{ *pulumi.OutputState }

func (ItemSectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemSection)(nil)).Elem()
}

func (o ItemSectionOutput) ToItemSectionOutput() ItemSectionOutput {
	return o
}

func (o ItemSectionOutput) ToItemSectionOutputWithContext(ctx context.Context) ItemSectionOutput {
	return o
}

// A list of custom fields in the section.
func (o ItemSectionOutput) Fields() ItemSectionFieldArrayOutput {
	return o.ApplyT(func(v ItemSection) []ItemSectionField { return v.Fields }).(ItemSectionFieldArrayOutput)
}

// A unique identifier for the section.
func (o ItemSectionOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ItemSection) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The label for the section.
func (o ItemSectionOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v ItemSection) string { return v.Label }).(pulumi.StringOutput)
}

type ItemSectionArrayOutput struct{ *pulumi.OutputState }

func (ItemSectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ItemSection)(nil)).Elem()
}

func (o ItemSectionArrayOutput) ToItemSectionArrayOutput() ItemSectionArrayOutput {
	return o
}

func (o ItemSectionArrayOutput) ToItemSectionArrayOutputWithContext(ctx context.Context) ItemSectionArrayOutput {
	return o
}

func (o ItemSectionArrayOutput) Index(i pulumi.IntInput) ItemSectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ItemSection {
		return vs[0].([]ItemSection)[vs[1].(int)]
	}).(ItemSectionOutput)
}

type ItemSectionField struct {
	Id    *string `pulumi:"id"`
	Label string  `pulumi:"label"`
	// Password for this item.
	PasswordRecipe *ItemSectionFieldPasswordRecipe `pulumi:"passwordRecipe"`
	Purpose        *string                         `pulumi:"purpose"`
	// (Only applies to the database category) The type of database. One of ["db2" "filemaker" "msaccess" "mssql" "mysql" "oracle" "postgresql" "sqlite" "other"]
	Type  *string `pulumi:"type"`
	Value *string `pulumi:"value"`
}

// ItemSectionFieldInput is an input type that accepts ItemSectionFieldArgs and ItemSectionFieldOutput values.
// You can construct a concrete instance of `ItemSectionFieldInput` via:
//
//          ItemSectionFieldArgs{...}
type ItemSectionFieldInput interface {
	pulumi.Input

	ToItemSectionFieldOutput() ItemSectionFieldOutput
	ToItemSectionFieldOutputWithContext(context.Context) ItemSectionFieldOutput
}

type ItemSectionFieldArgs struct {
	Id    pulumi.StringPtrInput `pulumi:"id"`
	Label pulumi.StringInput    `pulumi:"label"`
	// Password for this item.
	PasswordRecipe ItemSectionFieldPasswordRecipePtrInput `pulumi:"passwordRecipe"`
	Purpose        pulumi.StringPtrInput                  `pulumi:"purpose"`
	// (Only applies to the database category) The type of database. One of ["db2" "filemaker" "msaccess" "mssql" "mysql" "oracle" "postgresql" "sqlite" "other"]
	Type  pulumi.StringPtrInput `pulumi:"type"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ItemSectionFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemSectionField)(nil)).Elem()
}

func (i ItemSectionFieldArgs) ToItemSectionFieldOutput() ItemSectionFieldOutput {
	return i.ToItemSectionFieldOutputWithContext(context.Background())
}

func (i ItemSectionFieldArgs) ToItemSectionFieldOutputWithContext(ctx context.Context) ItemSectionFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemSectionFieldOutput)
}

// ItemSectionFieldArrayInput is an input type that accepts ItemSectionFieldArray and ItemSectionFieldArrayOutput values.
// You can construct a concrete instance of `ItemSectionFieldArrayInput` via:
//
//          ItemSectionFieldArray{ ItemSectionFieldArgs{...} }
type ItemSectionFieldArrayInput interface {
	pulumi.Input

	ToItemSectionFieldArrayOutput() ItemSectionFieldArrayOutput
	ToItemSectionFieldArrayOutputWithContext(context.Context) ItemSectionFieldArrayOutput
}

type ItemSectionFieldArray []ItemSectionFieldInput

func (ItemSectionFieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ItemSectionField)(nil)).Elem()
}

func (i ItemSectionFieldArray) ToItemSectionFieldArrayOutput() ItemSectionFieldArrayOutput {
	return i.ToItemSectionFieldArrayOutputWithContext(context.Background())
}

func (i ItemSectionFieldArray) ToItemSectionFieldArrayOutputWithContext(ctx context.Context) ItemSectionFieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemSectionFieldArrayOutput)
}

type ItemSectionFieldOutput struct{ *pulumi.OutputState }

func (ItemSectionFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemSectionField)(nil)).Elem()
}

func (o ItemSectionFieldOutput) ToItemSectionFieldOutput() ItemSectionFieldOutput {
	return o
}

func (o ItemSectionFieldOutput) ToItemSectionFieldOutputWithContext(ctx context.Context) ItemSectionFieldOutput {
	return o
}

func (o ItemSectionFieldOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ItemSectionField) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ItemSectionFieldOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v ItemSectionField) string { return v.Label }).(pulumi.StringOutput)
}

// Password for this item.
func (o ItemSectionFieldOutput) PasswordRecipe() ItemSectionFieldPasswordRecipePtrOutput {
	return o.ApplyT(func(v ItemSectionField) *ItemSectionFieldPasswordRecipe { return v.PasswordRecipe }).(ItemSectionFieldPasswordRecipePtrOutput)
}

func (o ItemSectionFieldOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ItemSectionField) *string { return v.Purpose }).(pulumi.StringPtrOutput)
}

// (Only applies to the database category) The type of database. One of ["db2" "filemaker" "msaccess" "mssql" "mysql" "oracle" "postgresql" "sqlite" "other"]
func (o ItemSectionFieldOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ItemSectionField) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ItemSectionFieldOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ItemSectionField) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ItemSectionFieldArrayOutput struct{ *pulumi.OutputState }

func (ItemSectionFieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ItemSectionField)(nil)).Elem()
}

func (o ItemSectionFieldArrayOutput) ToItemSectionFieldArrayOutput() ItemSectionFieldArrayOutput {
	return o
}

func (o ItemSectionFieldArrayOutput) ToItemSectionFieldArrayOutputWithContext(ctx context.Context) ItemSectionFieldArrayOutput {
	return o
}

func (o ItemSectionFieldArrayOutput) Index(i pulumi.IntInput) ItemSectionFieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ItemSectionField {
		return vs[0].([]ItemSectionField)[vs[1].(int)]
	}).(ItemSectionFieldOutput)
}

type ItemSectionFieldPasswordRecipe struct {
	// Use digits [0-9] when generating the password.
	Digits *bool `pulumi:"digits"`
	// The length of the password to be generated.
	Length *int `pulumi:"length"`
	// Use letters [a-zA-Z] when generating the password.
	Letters *bool `pulumi:"letters"`
	// Use symbols [!@.-_*] when generating the password.
	Symbols *bool `pulumi:"symbols"`
}

// ItemSectionFieldPasswordRecipeInput is an input type that accepts ItemSectionFieldPasswordRecipeArgs and ItemSectionFieldPasswordRecipeOutput values.
// You can construct a concrete instance of `ItemSectionFieldPasswordRecipeInput` via:
//
//          ItemSectionFieldPasswordRecipeArgs{...}
type ItemSectionFieldPasswordRecipeInput interface {
	pulumi.Input

	ToItemSectionFieldPasswordRecipeOutput() ItemSectionFieldPasswordRecipeOutput
	ToItemSectionFieldPasswordRecipeOutputWithContext(context.Context) ItemSectionFieldPasswordRecipeOutput
}

type ItemSectionFieldPasswordRecipeArgs struct {
	// Use digits [0-9] when generating the password.
	Digits pulumi.BoolPtrInput `pulumi:"digits"`
	// The length of the password to be generated.
	Length pulumi.IntPtrInput `pulumi:"length"`
	// Use letters [a-zA-Z] when generating the password.
	Letters pulumi.BoolPtrInput `pulumi:"letters"`
	// Use symbols [!@.-_*] when generating the password.
	Symbols pulumi.BoolPtrInput `pulumi:"symbols"`
}

func (ItemSectionFieldPasswordRecipeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemSectionFieldPasswordRecipe)(nil)).Elem()
}

func (i ItemSectionFieldPasswordRecipeArgs) ToItemSectionFieldPasswordRecipeOutput() ItemSectionFieldPasswordRecipeOutput {
	return i.ToItemSectionFieldPasswordRecipeOutputWithContext(context.Background())
}

func (i ItemSectionFieldPasswordRecipeArgs) ToItemSectionFieldPasswordRecipeOutputWithContext(ctx context.Context) ItemSectionFieldPasswordRecipeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemSectionFieldPasswordRecipeOutput)
}

func (i ItemSectionFieldPasswordRecipeArgs) ToItemSectionFieldPasswordRecipePtrOutput() ItemSectionFieldPasswordRecipePtrOutput {
	return i.ToItemSectionFieldPasswordRecipePtrOutputWithContext(context.Background())
}

func (i ItemSectionFieldPasswordRecipeArgs) ToItemSectionFieldPasswordRecipePtrOutputWithContext(ctx context.Context) ItemSectionFieldPasswordRecipePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemSectionFieldPasswordRecipeOutput).ToItemSectionFieldPasswordRecipePtrOutputWithContext(ctx)
}

// ItemSectionFieldPasswordRecipePtrInput is an input type that accepts ItemSectionFieldPasswordRecipeArgs, ItemSectionFieldPasswordRecipePtr and ItemSectionFieldPasswordRecipePtrOutput values.
// You can construct a concrete instance of `ItemSectionFieldPasswordRecipePtrInput` via:
//
//          ItemSectionFieldPasswordRecipeArgs{...}
//
//  or:
//
//          nil
type ItemSectionFieldPasswordRecipePtrInput interface {
	pulumi.Input

	ToItemSectionFieldPasswordRecipePtrOutput() ItemSectionFieldPasswordRecipePtrOutput
	ToItemSectionFieldPasswordRecipePtrOutputWithContext(context.Context) ItemSectionFieldPasswordRecipePtrOutput
}

type itemSectionFieldPasswordRecipePtrType ItemSectionFieldPasswordRecipeArgs

func ItemSectionFieldPasswordRecipePtr(v *ItemSectionFieldPasswordRecipeArgs) ItemSectionFieldPasswordRecipePtrInput {
	return (*itemSectionFieldPasswordRecipePtrType)(v)
}

func (*itemSectionFieldPasswordRecipePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ItemSectionFieldPasswordRecipe)(nil)).Elem()
}

func (i *itemSectionFieldPasswordRecipePtrType) ToItemSectionFieldPasswordRecipePtrOutput() ItemSectionFieldPasswordRecipePtrOutput {
	return i.ToItemSectionFieldPasswordRecipePtrOutputWithContext(context.Background())
}

func (i *itemSectionFieldPasswordRecipePtrType) ToItemSectionFieldPasswordRecipePtrOutputWithContext(ctx context.Context) ItemSectionFieldPasswordRecipePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItemSectionFieldPasswordRecipePtrOutput)
}

type ItemSectionFieldPasswordRecipeOutput struct{ *pulumi.OutputState }

func (ItemSectionFieldPasswordRecipeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ItemSectionFieldPasswordRecipe)(nil)).Elem()
}

func (o ItemSectionFieldPasswordRecipeOutput) ToItemSectionFieldPasswordRecipeOutput() ItemSectionFieldPasswordRecipeOutput {
	return o
}

func (o ItemSectionFieldPasswordRecipeOutput) ToItemSectionFieldPasswordRecipeOutputWithContext(ctx context.Context) ItemSectionFieldPasswordRecipeOutput {
	return o
}

func (o ItemSectionFieldPasswordRecipeOutput) ToItemSectionFieldPasswordRecipePtrOutput() ItemSectionFieldPasswordRecipePtrOutput {
	return o.ToItemSectionFieldPasswordRecipePtrOutputWithContext(context.Background())
}

func (o ItemSectionFieldPasswordRecipeOutput) ToItemSectionFieldPasswordRecipePtrOutputWithContext(ctx context.Context) ItemSectionFieldPasswordRecipePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ItemSectionFieldPasswordRecipe) *ItemSectionFieldPasswordRecipe {
		return &v
	}).(ItemSectionFieldPasswordRecipePtrOutput)
}

// Use digits [0-9] when generating the password.
func (o ItemSectionFieldPasswordRecipeOutput) Digits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ItemSectionFieldPasswordRecipe) *bool { return v.Digits }).(pulumi.BoolPtrOutput)
}

// The length of the password to be generated.
func (o ItemSectionFieldPasswordRecipeOutput) Length() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ItemSectionFieldPasswordRecipe) *int { return v.Length }).(pulumi.IntPtrOutput)
}

// Use letters [a-zA-Z] when generating the password.
func (o ItemSectionFieldPasswordRecipeOutput) Letters() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ItemSectionFieldPasswordRecipe) *bool { return v.Letters }).(pulumi.BoolPtrOutput)
}

// Use symbols [!@.-_*] when generating the password.
func (o ItemSectionFieldPasswordRecipeOutput) Symbols() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ItemSectionFieldPasswordRecipe) *bool { return v.Symbols }).(pulumi.BoolPtrOutput)
}

type ItemSectionFieldPasswordRecipePtrOutput struct{ *pulumi.OutputState }

func (ItemSectionFieldPasswordRecipePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ItemSectionFieldPasswordRecipe)(nil)).Elem()
}

func (o ItemSectionFieldPasswordRecipePtrOutput) ToItemSectionFieldPasswordRecipePtrOutput() ItemSectionFieldPasswordRecipePtrOutput {
	return o
}

func (o ItemSectionFieldPasswordRecipePtrOutput) ToItemSectionFieldPasswordRecipePtrOutputWithContext(ctx context.Context) ItemSectionFieldPasswordRecipePtrOutput {
	return o
}

func (o ItemSectionFieldPasswordRecipePtrOutput) Elem() ItemSectionFieldPasswordRecipeOutput {
	return o.ApplyT(func(v *ItemSectionFieldPasswordRecipe) ItemSectionFieldPasswordRecipe {
		if v != nil {
			return *v
		}
		var ret ItemSectionFieldPasswordRecipe
		return ret
	}).(ItemSectionFieldPasswordRecipeOutput)
}

// Use digits [0-9] when generating the password.
func (o ItemSectionFieldPasswordRecipePtrOutput) Digits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ItemSectionFieldPasswordRecipe) *bool {
		if v == nil {
			return nil
		}
		return v.Digits
	}).(pulumi.BoolPtrOutput)
}

// The length of the password to be generated.
func (o ItemSectionFieldPasswordRecipePtrOutput) Length() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ItemSectionFieldPasswordRecipe) *int {
		if v == nil {
			return nil
		}
		return v.Length
	}).(pulumi.IntPtrOutput)
}

// Use letters [a-zA-Z] when generating the password.
func (o ItemSectionFieldPasswordRecipePtrOutput) Letters() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ItemSectionFieldPasswordRecipe) *bool {
		if v == nil {
			return nil
		}
		return v.Letters
	}).(pulumi.BoolPtrOutput)
}

// Use symbols [!@.-_*] when generating the password.
func (o ItemSectionFieldPasswordRecipePtrOutput) Symbols() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ItemSectionFieldPasswordRecipe) *bool {
		if v == nil {
			return nil
		}
		return v.Symbols
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetItemSectionInput)(nil)).Elem(), GetItemSectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetItemSectionArrayInput)(nil)).Elem(), GetItemSectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetItemSectionFieldInput)(nil)).Elem(), GetItemSectionFieldArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetItemSectionFieldArrayInput)(nil)).Elem(), GetItemSectionFieldArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItemPasswordRecipeInput)(nil)).Elem(), ItemPasswordRecipeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItemPasswordRecipePtrInput)(nil)).Elem(), ItemPasswordRecipeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItemSectionInput)(nil)).Elem(), ItemSectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItemSectionArrayInput)(nil)).Elem(), ItemSectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItemSectionFieldInput)(nil)).Elem(), ItemSectionFieldArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItemSectionFieldArrayInput)(nil)).Elem(), ItemSectionFieldArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItemSectionFieldPasswordRecipeInput)(nil)).Elem(), ItemSectionFieldPasswordRecipeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItemSectionFieldPasswordRecipePtrInput)(nil)).Elem(), ItemSectionFieldPasswordRecipeArgs{})
	pulumi.RegisterOutputType(GetItemSectionOutput{})
	pulumi.RegisterOutputType(GetItemSectionArrayOutput{})
	pulumi.RegisterOutputType(GetItemSectionFieldOutput{})
	pulumi.RegisterOutputType(GetItemSectionFieldArrayOutput{})
	pulumi.RegisterOutputType(ItemPasswordRecipeOutput{})
	pulumi.RegisterOutputType(ItemPasswordRecipePtrOutput{})
	pulumi.RegisterOutputType(ItemSectionOutput{})
	pulumi.RegisterOutputType(ItemSectionArrayOutput{})
	pulumi.RegisterOutputType(ItemSectionFieldOutput{})
	pulumi.RegisterOutputType(ItemSectionFieldArrayOutput{})
	pulumi.RegisterOutputType(ItemSectionFieldPasswordRecipeOutput{})
	pulumi.RegisterOutputType(ItemSectionFieldPasswordRecipePtrOutput{})
}
