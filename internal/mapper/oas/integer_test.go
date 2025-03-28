// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oas_test

import (
	"testing"

	"github.com/greatman/terraform-plugin-codegen-spec/code"
	"github.com/greatman/terraform-plugin-codegen-spec/datasource"
	"github.com/greatman/terraform-plugin-codegen-spec/provider"
	"github.com/greatman/terraform-plugin-codegen-spec/resource"
	"github.com/greatman/terraform-plugin-codegen-spec/schema"
	"gopkg.in/yaml.v3"

	"github.com/greatman/terraform-plugin-codegen-openapi/internal/mapper/attrmapper"
	"github.com/greatman/terraform-plugin-codegen-openapi/internal/mapper/oas"

	"github.com/google/go-cmp/cmp"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/orderedmap"
)

func TestBuildIntegerResource(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		schema             *base.Schema
		expectedAttributes attrmapper.ResourceAttributes
	}{
		"int64 attributes": {
			schema: &base.Schema{
				Type:     []string{"object"},
				Required: []string{"int64_prop_required"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_prop": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"integer"},
						Description: "hey there! I'm an int64 type.",
					}),
					"int64_prop_required": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"integer"},
						Description: "hey there! I'm an int64 type, required.",
					}),
				}),
			},
			expectedAttributes: attrmapper.ResourceAttributes{
				&attrmapper.ResourceInt64Attribute{
					Name: "int64_prop",
					Int64Attribute: resource.Int64Attribute{
						ComputedOptionalRequired: schema.ComputedOptional,
						Description:              pointer("hey there! I'm an int64 type."),
					},
				},
				&attrmapper.ResourceInt64Attribute{
					Name: "int64_prop_required",
					Int64Attribute: resource.Int64Attribute{
						ComputedOptionalRequired: schema.Required,
						Description:              pointer("hey there! I'm an int64 type, required."),
					},
				},
			},
		},
		"int64 attributes default": {
			schema: &base.Schema{
				Type:     []string{"object"},
				Required: []string{"int64_prop_required_default_non_zero"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_prop_default_non_zero": base.CreateSchemaProxy(&base.Schema{
						Type:    []string{"integer"},
						Default: &yaml.Node{Kind: yaml.ScalarNode, Value: "123"},
					}),
					"int64_prop_default_zero": base.CreateSchemaProxy(&base.Schema{
						Type:    []string{"integer"},
						Default: &yaml.Node{Kind: yaml.ScalarNode, Value: "0"},
					}),
					"int64_prop_required_default_non_zero": base.CreateSchemaProxy(&base.Schema{
						Type:    []string{"integer"},
						Default: &yaml.Node{Kind: yaml.ScalarNode, Value: "123"},
					}),
				}),
			},
			expectedAttributes: attrmapper.ResourceAttributes{
				&attrmapper.ResourceInt64Attribute{
					Name: "int64_prop_default_non_zero",
					Int64Attribute: resource.Int64Attribute{
						ComputedOptionalRequired: schema.ComputedOptional,
						Default: &schema.Int64Default{
							Static: pointer(int64(123)),
						},
					},
				},
				&attrmapper.ResourceInt64Attribute{
					Name: "int64_prop_default_zero",
					Int64Attribute: resource.Int64Attribute{
						ComputedOptionalRequired: schema.ComputedOptional,
						Default: &schema.Int64Default{
							Static: pointer(int64(0)),
						},
					},
				},
				&attrmapper.ResourceInt64Attribute{
					Name: "int64_prop_required_default_non_zero",
					Int64Attribute: resource.Int64Attribute{
						// Intentionally not required due to default
						ComputedOptionalRequired: schema.ComputedOptional,
						Default: &schema.Int64Default{
							Static: pointer(int64(123)),
						},
					},
				},
			},
		},
		"int64 attributes deprecated": {
			schema: &base.Schema{
				Type: []string{"object"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_prop": base.CreateSchemaProxy(&base.Schema{
						Type:       []string{"integer"},
						Deprecated: pointer(true),
					}),
				}),
			},
			expectedAttributes: attrmapper.ResourceAttributes{
				&attrmapper.ResourceInt64Attribute{
					Name: "int64_prop",
					Int64Attribute: resource.Int64Attribute{
						ComputedOptionalRequired: schema.ComputedOptional,
						DeprecationMessage:       pointer("This attribute is deprecated."),
					},
				},
			},
		},
		"list attributes with int64 element type": {
			schema: &base.Schema{
				Type:     []string{"object"},
				Required: []string{"int64_list_prop_required"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_list_prop": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"array"},
						Description: "hey there! I'm a list of int64s.",
						Items: &base.DynamicValue[*base.SchemaProxy, bool]{
							A: base.CreateSchemaProxy(&base.Schema{
								Type: []string{"integer"},
							}),
						},
					}),
					"int64_list_prop_required": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"array"},
						Description: "hey there! I'm a list of int64s, required.",
						Items: &base.DynamicValue[*base.SchemaProxy, bool]{
							A: base.CreateSchemaProxy(&base.Schema{
								Type: []string{"integer"},
							}),
						},
					}),
				}),
			},
			expectedAttributes: attrmapper.ResourceAttributes{
				&attrmapper.ResourceListAttribute{
					Name: "int64_list_prop",
					ListAttribute: resource.ListAttribute{
						ComputedOptionalRequired: schema.ComputedOptional,
						Description:              pointer("hey there! I'm a list of int64s."),
						ElementType: schema.ElementType{
							Int64: &schema.Int64Type{},
						},
					},
				},
				&attrmapper.ResourceListAttribute{
					Name: "int64_list_prop_required",
					ListAttribute: resource.ListAttribute{
						ComputedOptionalRequired: schema.Required,
						Description:              pointer("hey there! I'm a list of int64s, required."),
						ElementType: schema.ElementType{
							Int64: &schema.Int64Type{},
						},
					},
				},
			},
		},
		"validators": {
			schema: &base.Schema{
				Type:     []string{"object"},
				Required: []string{"int64_prop"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_prop": base.CreateSchemaProxy(&base.Schema{
						Type: []string{"integer"},
						Enum: []*yaml.Node{
							{Kind: yaml.ScalarNode, Value: "1"},
							{Kind: yaml.ScalarNode, Value: "2"},
						},
					}),
				}),
			},
			expectedAttributes: attrmapper.ResourceAttributes{
				&attrmapper.ResourceInt64Attribute{
					Name: "int64_prop",
					Int64Attribute: resource.Int64Attribute{
						ComputedOptionalRequired: schema.Required,
						Validators: []schema.Int64Validator{
							{
								Custom: &schema.CustomValidator{
									Imports: []code.Import{
										{
											Path: "github.com/hashicorp/terraform-plugin-framework-validators/int64validator",
										},
									},
									SchemaDefinition: "int64validator.OneOf(\n1,\n2,\n)",
								},
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			schema := oas.OASSchema{Schema: testCase.schema}
			attributes, err := schema.BuildResourceAttributes()
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(attributes, testCase.expectedAttributes); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestBuildIntegerDataSource(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		schema             *base.Schema
		expectedAttributes attrmapper.DataSourceAttributes
	}{
		"int64 attributes": {
			schema: &base.Schema{
				Type:     []string{"object"},
				Required: []string{"int64_prop_required"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_prop": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"integer"},
						Description: "hey there! I'm an int64 type.",
					}),
					"int64_prop_required": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"integer"},
						Description: "hey there! I'm an int64 type, required.",
					}),
				}),
			},
			expectedAttributes: attrmapper.DataSourceAttributes{
				&attrmapper.DataSourceInt64Attribute{
					Name: "int64_prop",
					Int64Attribute: datasource.Int64Attribute{
						ComputedOptionalRequired: schema.ComputedOptional,
						Description:              pointer("hey there! I'm an int64 type."),
					},
				},
				&attrmapper.DataSourceInt64Attribute{
					Name: "int64_prop_required",
					Int64Attribute: datasource.Int64Attribute{
						ComputedOptionalRequired: schema.Required,
						Description:              pointer("hey there! I'm an int64 type, required."),
					},
				},
			},
		},
		"int64 attributes deprecated": {
			schema: &base.Schema{
				Type: []string{"object"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_prop": base.CreateSchemaProxy(&base.Schema{
						Type:       []string{"integer"},
						Deprecated: pointer(true),
					}),
				}),
			},
			expectedAttributes: attrmapper.DataSourceAttributes{
				&attrmapper.DataSourceInt64Attribute{
					Name: "int64_prop",
					Int64Attribute: datasource.Int64Attribute{
						ComputedOptionalRequired: schema.ComputedOptional,
						DeprecationMessage:       pointer("This attribute is deprecated."),
					},
				},
			},
		},
		"list attributes with int64 element type": {
			schema: &base.Schema{
				Type:     []string{"object"},
				Required: []string{"int64_list_prop_required"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_list_prop": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"array"},
						Description: "hey there! I'm a list of int64s.",
						Items: &base.DynamicValue[*base.SchemaProxy, bool]{
							A: base.CreateSchemaProxy(&base.Schema{
								Type: []string{"integer"},
							}),
						},
					}),
					"int64_list_prop_required": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"array"},
						Description: "hey there! I'm a list of int64s, required.",
						Items: &base.DynamicValue[*base.SchemaProxy, bool]{
							A: base.CreateSchemaProxy(&base.Schema{
								Type: []string{"integer"},
							}),
						},
					}),
				}),
			},
			expectedAttributes: attrmapper.DataSourceAttributes{
				&attrmapper.DataSourceListAttribute{
					Name: "int64_list_prop",
					ListAttribute: datasource.ListAttribute{
						ComputedOptionalRequired: schema.ComputedOptional,
						Description:              pointer("hey there! I'm a list of int64s."),
						ElementType: schema.ElementType{
							Int64: &schema.Int64Type{},
						},
					},
				},
				&attrmapper.DataSourceListAttribute{
					Name: "int64_list_prop_required",
					ListAttribute: datasource.ListAttribute{
						ComputedOptionalRequired: schema.Required,
						Description:              pointer("hey there! I'm a list of int64s, required."),
						ElementType: schema.ElementType{
							Int64: &schema.Int64Type{},
						},
					},
				},
			},
		},
		"validators": {
			schema: &base.Schema{
				Type:     []string{"object"},
				Required: []string{"int64_prop"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_prop": base.CreateSchemaProxy(&base.Schema{
						Type: []string{"integer"},
						Enum: []*yaml.Node{
							{Kind: yaml.ScalarNode, Value: "1"},
							{Kind: yaml.ScalarNode, Value: "2"},
						},
					}),
				}),
			},
			expectedAttributes: attrmapper.DataSourceAttributes{
				&attrmapper.DataSourceInt64Attribute{
					Name: "int64_prop",
					Int64Attribute: datasource.Int64Attribute{
						ComputedOptionalRequired: schema.Required,
						Validators: []schema.Int64Validator{
							{
								Custom: &schema.CustomValidator{
									Imports: []code.Import{
										{
											Path: "github.com/hashicorp/terraform-plugin-framework-validators/int64validator",
										},
									},
									SchemaDefinition: "int64validator.OneOf(\n1,\n2,\n)",
								},
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			schema := oas.OASSchema{Schema: testCase.schema}
			attributes, err := schema.BuildDataSourceAttributes()
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(attributes, testCase.expectedAttributes); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestBuildIntegerProvider(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		schema             *base.Schema
		expectedAttributes attrmapper.ProviderAttributes
	}{
		"int64 attributes": {
			schema: &base.Schema{
				Type:     []string{"object"},
				Required: []string{"int64_prop_required"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_prop": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"integer"},
						Description: "hey there! I'm an int64 type.",
					}),
					"int64_prop_required": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"integer"},
						Description: "hey there! I'm an int64 type, required.",
					}),
				}),
			},
			expectedAttributes: attrmapper.ProviderAttributes{
				&attrmapper.ProviderInt64Attribute{
					Name: "int64_prop",
					Int64Attribute: provider.Int64Attribute{
						OptionalRequired: schema.Optional,
						Description:      pointer("hey there! I'm an int64 type."),
					},
				},
				&attrmapper.ProviderInt64Attribute{
					Name: "int64_prop_required",
					Int64Attribute: provider.Int64Attribute{
						OptionalRequired: schema.Required,
						Description:      pointer("hey there! I'm an int64 type, required."),
					},
				},
			},
		},
		"int64 attributes deprecated": {
			schema: &base.Schema{
				Type: []string{"object"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_prop": base.CreateSchemaProxy(&base.Schema{
						Type:       []string{"integer"},
						Deprecated: pointer(true),
					}),
				}),
			},
			expectedAttributes: attrmapper.ProviderAttributes{
				&attrmapper.ProviderInt64Attribute{
					Name: "int64_prop",
					Int64Attribute: provider.Int64Attribute{
						OptionalRequired:   schema.Optional,
						DeprecationMessage: pointer("This attribute is deprecated."),
					},
				},
			},
		},
		"list attributes with int64 element type": {
			schema: &base.Schema{
				Type:     []string{"object"},
				Required: []string{"int64_list_prop_required"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_list_prop": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"array"},
						Description: "hey there! I'm a list of int64s.",
						Items: &base.DynamicValue[*base.SchemaProxy, bool]{
							A: base.CreateSchemaProxy(&base.Schema{
								Type: []string{"integer"},
							}),
						},
					}),
					"int64_list_prop_required": base.CreateSchemaProxy(&base.Schema{
						Type:        []string{"array"},
						Description: "hey there! I'm a list of int64s, required.",
						Items: &base.DynamicValue[*base.SchemaProxy, bool]{
							A: base.CreateSchemaProxy(&base.Schema{
								Type: []string{"integer"},
							}),
						},
					}),
				}),
			},
			expectedAttributes: attrmapper.ProviderAttributes{
				&attrmapper.ProviderListAttribute{
					Name: "int64_list_prop",
					ListAttribute: provider.ListAttribute{
						OptionalRequired: schema.Optional,
						Description:      pointer("hey there! I'm a list of int64s."),
						ElementType: schema.ElementType{
							Int64: &schema.Int64Type{},
						},
					},
				},
				&attrmapper.ProviderListAttribute{
					Name: "int64_list_prop_required",
					ListAttribute: provider.ListAttribute{
						OptionalRequired: schema.Required,
						Description:      pointer("hey there! I'm a list of int64s, required."),
						ElementType: schema.ElementType{
							Int64: &schema.Int64Type{},
						},
					},
				},
			},
		},
		"validators": {
			schema: &base.Schema{
				Type:     []string{"object"},
				Required: []string{"int64_prop"},
				Properties: orderedmap.ToOrderedMap(map[string]*base.SchemaProxy{
					"int64_prop": base.CreateSchemaProxy(&base.Schema{
						Type: []string{"integer"},
						Enum: []*yaml.Node{
							{Kind: yaml.ScalarNode, Value: "1"},
							{Kind: yaml.ScalarNode, Value: "2"},
						},
					}),
				}),
			},
			expectedAttributes: attrmapper.ProviderAttributes{
				&attrmapper.ProviderInt64Attribute{
					Name: "int64_prop",
					Int64Attribute: provider.Int64Attribute{
						OptionalRequired: schema.Required,
						Validators: []schema.Int64Validator{
							{
								Custom: &schema.CustomValidator{
									Imports: []code.Import{
										{
											Path: "github.com/hashicorp/terraform-plugin-framework-validators/int64validator",
										},
									},
									SchemaDefinition: "int64validator.OneOf(\n1,\n2,\n)",
								},
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			schema := oas.OASSchema{Schema: testCase.schema}
			attributes, err := schema.BuildProviderAttributes()
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(attributes, testCase.expectedAttributes); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGetIntegerValidators(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		schema   oas.OASSchema
		expected []schema.Int64Validator
	}{
		"none": {
			schema: oas.OASSchema{
				Schema: &base.Schema{
					Type: []string{"integer"},
				},
			},
			expected: nil,
		},
		"enum": {
			schema: oas.OASSchema{
				Schema: &base.Schema{
					Type: []string{"integer"},
					Enum: []*yaml.Node{
						{Kind: yaml.ScalarNode, Value: "1"},
						{Kind: yaml.ScalarNode, Value: "2"},
					},
				},
			},
			expected: []schema.Int64Validator{
				{
					Custom: &schema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/hashicorp/terraform-plugin-framework-validators/int64validator",
							},
						},
						SchemaDefinition: "int64validator.OneOf(\n1,\n2,\n)",
					},
				},
			},
		},
		"maximum": {
			schema: oas.OASSchema{
				Schema: &base.Schema{
					Type:    []string{"integer"},
					Maximum: pointer(float64(123)),
				},
			},
			expected: []schema.Int64Validator{
				{
					Custom: &schema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/hashicorp/terraform-plugin-framework-validators/int64validator",
							},
						},
						SchemaDefinition: "int64validator.AtMost(123)",
					},
				},
			},
		},
		"maximum-and-minimum": {
			schema: oas.OASSchema{
				Schema: &base.Schema{
					Type:    []string{"integer"},
					Minimum: pointer(float64(123.2)),
					Maximum: pointer(float64(456.2)),
				},
			},
			expected: []schema.Int64Validator{
				{
					Custom: &schema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/hashicorp/terraform-plugin-framework-validators/int64validator",
							},
						},
						SchemaDefinition: "int64validator.Between(123, 456)",
					},
				},
			},
		},
		"minimum": {
			schema: oas.OASSchema{
				Schema: &base.Schema{
					Type:    []string{"integer"},
					Minimum: pointer(float64(123)),
				},
			},
			expected: []schema.Int64Validator{
				{
					Custom: &schema.CustomValidator{
						Imports: []code.Import{
							{
								Path: "github.com/hashicorp/terraform-plugin-framework-validators/int64validator",
							},
						},
						SchemaDefinition: "int64validator.AtLeast(123)",
					},
				},
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.schema.GetIntegerValidators()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
