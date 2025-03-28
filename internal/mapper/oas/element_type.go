// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oas

import (
	"fmt"

	"github.com/greatman/terraform-plugin-codegen-openapi/internal/mapper/util"
	"github.com/greatman/terraform-plugin-codegen-spec/schema"
)

func (s *OASSchema) BuildElementType() (schema.ElementType, *SchemaError) {
	switch s.Type {
	case util.OAS_type_string:
		return s.BuildStringElementType()
	case util.OAS_type_integer:
		return s.BuildIntegerElementType()
	case util.OAS_type_number:
		return s.BuildNumberElementType()
	case util.OAS_type_boolean:
		return s.BuildBoolElementType()
	case util.OAS_type_array:
		return s.BuildCollectionElementType()
	case util.OAS_type_object:
		if s.IsMap() {
			return s.BuildMapElementType()
		}
		return s.BuildObjectElementType()

	default:
		return schema.ElementType{}, SchemaErrorFromNode(fmt.Errorf("invalid schema type '%s'", s.Type), s.Schema, Type)
	}
}
