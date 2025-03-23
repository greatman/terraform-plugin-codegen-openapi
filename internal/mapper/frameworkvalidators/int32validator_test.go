// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frameworkvalidators_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/greatman/terraform-plugin-codegen-spec/code"
	"github.com/greatman/terraform-plugin-codegen-spec/schema"

	"github.com/greatman/terraform-plugin-codegen-openapi/internal/mapper/frameworkvalidators"
)

func TestInt32ValidatorAtLeast(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		min      int32
		expected *schema.CustomValidator
	}{
		"test": {
			min: 123,
			expected: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Path: "github.com/hashicorp/terraform-plugin-framework-validators/int32validator",
					},
				},
				SchemaDefinition: "int32validator.AtLeast(123)",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := frameworkvalidators.Int32ValidatorAtLeast(testCase.min)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestInt32ValidatorAtMost(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		max      int32
		expected *schema.CustomValidator
	}{
		"test": {
			max: 123,
			expected: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Path: "github.com/hashicorp/terraform-plugin-framework-validators/int32validator",
					},
				},
				SchemaDefinition: "int32validator.AtMost(123)",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := frameworkvalidators.Int32ValidatorAtMost(testCase.max)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestInt32ValidatorBetween(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		min      int32
		max      int32
		expected *schema.CustomValidator
	}{
		"test": {
			min: 123,
			max: 456,
			expected: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Path: "github.com/hashicorp/terraform-plugin-framework-validators/int32validator",
					},
				},
				SchemaDefinition: "int32validator.Between(123, 456)",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := frameworkvalidators.Int32ValidatorBetween(testCase.min, testCase.max)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestInt32ValidatorOneOf(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		values   []int32
		expected *schema.CustomValidator
	}{
		"nil": {
			values:   nil,
			expected: nil,
		},
		"empty": {
			values:   []int32{},
			expected: nil,
		},
		"one": {
			values: []int32{1},
			expected: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Path: "github.com/hashicorp/terraform-plugin-framework-validators/int32validator",
					},
				},
				SchemaDefinition: "int32validator.OneOf(\n1,\n)",
			},
		},
		"multiple": {
			values: []int32{1, 2},
			expected: &schema.CustomValidator{
				Imports: []code.Import{
					{
						Path: "github.com/hashicorp/terraform-plugin-framework-validators/int32validator",
					},
				},
				SchemaDefinition: "int32validator.OneOf(\n1,\n2,\n)",
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := frameworkvalidators.Int32ValidatorOneOf(testCase.values)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
