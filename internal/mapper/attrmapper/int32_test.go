// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package attrmapper_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/greatman/terraform-plugin-codegen-spec/datasource"
	"github.com/greatman/terraform-plugin-codegen-spec/resource"
	"github.com/greatman/terraform-plugin-codegen-spec/schema"

	"github.com/greatman/terraform-plugin-codegen-openapi/internal/explorer"
	"github.com/greatman/terraform-plugin-codegen-openapi/internal/mapper/attrmapper"
)

func TestResourceInt32Attribute_Merge(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		targetAttribute   attrmapper.ResourceInt32Attribute
		mergeAttribute    attrmapper.ResourceAttribute
		expectedAttribute attrmapper.ResourceAttribute
	}{
		"mismatch type - no merge": {
			targetAttribute: attrmapper.ResourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
				},
			},
			mergeAttribute: &attrmapper.ResourceStringAttribute{
				Name: "string_attribute",
				StringAttribute: resource.StringAttribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("string description"),
				},
			},
			expectedAttribute: &attrmapper.ResourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
				},
			},
		},
		"populated description - no merge": {
			targetAttribute: attrmapper.ResourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("old Int32 description"),
				},
			},
			mergeAttribute: &attrmapper.ResourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.ComputedOptional,
					Description:              pointer("new Int32 description"),
				},
			},
			expectedAttribute: &attrmapper.ResourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("old Int32 description"),
				},
			},
		},
		"nil description - merge": {
			targetAttribute: attrmapper.ResourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
				},
			},
			mergeAttribute: &attrmapper.ResourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.ComputedOptional,
					Description:              pointer("new Int32 description"),
				},
			},
			expectedAttribute: &attrmapper.ResourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("new Int32 description"),
				},
			},
		},
		"empty description - merge": {
			targetAttribute: attrmapper.ResourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer(""),
				},
			},
			mergeAttribute: &attrmapper.ResourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.ComputedOptional,
					Description:              pointer("new Int32 description"),
				},
			},
			expectedAttribute: &attrmapper.ResourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("new Int32 description"),
				},
			},
		},
	}
	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, _ := testCase.targetAttribute.Merge(testCase.mergeAttribute)

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("Unexpected diagnostics (-got, +expected): %s", diff)
			}
		})
	}
}

func TestResourceInt32Attribute_ApplyOverride(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		attribute         attrmapper.ResourceInt32Attribute
		override          explorer.Override
		expectedAttribute attrmapper.ResourceAttribute
	}{
		"override description": {
			attribute: attrmapper.ResourceInt32Attribute{
				Name: "test_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("old description"),
				},
			},
			override: explorer.Override{
				Description: "new description",
			},
			expectedAttribute: &attrmapper.ResourceInt32Attribute{
				Name: "test_attribute",
				Int32Attribute: resource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("new description"),
				},
			},
		},
	}
	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, _ := testCase.attribute.ApplyOverride(testCase.override)

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("Unexpected diagnostics (-got, +expected): %s", diff)
			}
		})
	}
}

func TestDataSourceInt32Attribute_Merge(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		targetAttribute   attrmapper.DataSourceInt32Attribute
		mergeAttribute    attrmapper.DataSourceAttribute
		expectedAttribute attrmapper.DataSourceAttribute
	}{
		"mismatch type - no merge": {
			targetAttribute: attrmapper.DataSourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
				},
			},
			mergeAttribute: &attrmapper.DataSourceStringAttribute{
				Name: "string_attribute",
				StringAttribute: datasource.StringAttribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("string description"),
				},
			},
			expectedAttribute: &attrmapper.DataSourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
				},
			},
		},
		"populated description - no merge": {
			targetAttribute: attrmapper.DataSourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("old Int32 description"),
				},
			},
			mergeAttribute: &attrmapper.DataSourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.ComputedOptional,
					Description:              pointer("new Int32 description"),
				},
			},
			expectedAttribute: &attrmapper.DataSourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("old Int32 description"),
				},
			},
		},
		"nil description - merge": {
			targetAttribute: attrmapper.DataSourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
				},
			},
			mergeAttribute: &attrmapper.DataSourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.ComputedOptional,
					Description:              pointer("new Int32 description"),
				},
			},
			expectedAttribute: &attrmapper.DataSourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("new Int32 description"),
				},
			},
		},
		"empty description - merge": {
			targetAttribute: attrmapper.DataSourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer(""),
				},
			},
			mergeAttribute: &attrmapper.DataSourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.ComputedOptional,
					Description:              pointer("new Int32 description"),
				},
			},
			expectedAttribute: &attrmapper.DataSourceInt32Attribute{
				Name: "Int32_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("new Int32 description"),
				},
			},
		},
	}
	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, _ := testCase.targetAttribute.Merge(testCase.mergeAttribute)

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("Unexpected diagnostics (-got, +expected): %s", diff)
			}
		})
	}
}

func TestDataSourceInt32Attribute_ApplyOverride(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		attribute         attrmapper.DataSourceInt32Attribute
		override          explorer.Override
		expectedAttribute attrmapper.DataSourceAttribute
	}{
		"override description": {
			attribute: attrmapper.DataSourceInt32Attribute{
				Name: "test_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("old description"),
				},
			},
			override: explorer.Override{
				Description: "new description",
			},
			expectedAttribute: &attrmapper.DataSourceInt32Attribute{
				Name: "test_attribute",
				Int32Attribute: datasource.Int32Attribute{
					ComputedOptionalRequired: schema.Required,
					Description:              pointer("new description"),
				},
			},
		},
	}
	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, _ := testCase.attribute.ApplyOverride(testCase.override)

			if diff := cmp.Diff(got, testCase.expectedAttribute); diff != "" {
				t.Errorf("Unexpected diagnostics (-got, +expected): %s", diff)
			}
		})
	}
}
