// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frameworkvalidators

import (
	"strconv"
	"strings"

	"github.com/greatman/terraform-plugin-codegen-spec/code"
	"github.com/greatman/terraform-plugin-codegen-spec/schema"
)

const (
	// MapValidatorPackage is the name of the set validation package in
	// the framework validators module.
	MapValidatorPackage = "mapvalidator"
)

var (
	// MapValidatorCodeImport is a single allocation of the framework
	// validators module mapvalidator package import.
	MapValidatorCodeImport code.Import = CodeImport(MapValidatorPackage)
)

// MapValidatorSizeAtLeast returns a custom validator mapped to the
// Mapvalidator package SizeAtLeast function.
func MapValidatorSizeAtLeast(minimum int64) *schema.CustomValidator {
	var schemaDefinition strings.Builder

	schemaDefinition.WriteString(MapValidatorPackage)
	schemaDefinition.WriteString(".SizeAtLeast(")
	schemaDefinition.WriteString(strconv.FormatInt(minimum, 10))
	schemaDefinition.WriteString(")")

	return &schema.CustomValidator{
		Imports: []code.Import{
			MapValidatorCodeImport,
		},
		SchemaDefinition: schemaDefinition.String(),
	}
}

// MapValidatorSizeAtMost returns a custom validator mapped to the
// Mapvalidator package SizeAtMost function.
func MapValidatorSizeAtMost(maximum int64) *schema.CustomValidator {
	var schemaDefinition strings.Builder

	schemaDefinition.WriteString(MapValidatorPackage)
	schemaDefinition.WriteString(".SizeAtMost(")
	schemaDefinition.WriteString(strconv.FormatInt(maximum, 10))
	schemaDefinition.WriteString(")")

	return &schema.CustomValidator{
		Imports: []code.Import{
			MapValidatorCodeImport,
		},
		SchemaDefinition: schemaDefinition.String(),
	}
}

// MapValidatorSizeBetween returns a custom validator mapped to the
// Mapvalidator package SizeBetween function.
func MapValidatorSizeBetween(minimum, maximum int64) *schema.CustomValidator {
	var schemaDefinition strings.Builder

	schemaDefinition.WriteString(MapValidatorPackage)
	schemaDefinition.WriteString(".SizeBetween(")
	schemaDefinition.WriteString(strconv.FormatInt(minimum, 10))
	schemaDefinition.WriteString(", ")
	schemaDefinition.WriteString(strconv.FormatInt(maximum, 10))
	schemaDefinition.WriteString(")")

	return &schema.CustomValidator{
		Imports: []code.Import{
			MapValidatorCodeImport,
		},
		SchemaDefinition: schemaDefinition.String(),
	}
}
