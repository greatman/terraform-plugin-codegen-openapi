// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oas

import (
	"github.com/greatman/terraform-plugin-codegen-openapi/internal/mapper/attrmapper"
	"github.com/greatman/terraform-plugin-codegen-openapi/internal/mapper/frameworkvalidators"
	"github.com/greatman/terraform-plugin-codegen-openapi/internal/mapper/util"
	"github.com/greatman/terraform-plugin-codegen-spec/datasource"
	"github.com/greatman/terraform-plugin-codegen-spec/provider"
	"github.com/greatman/terraform-plugin-codegen-spec/resource"
	"github.com/greatman/terraform-plugin-codegen-spec/schema"
)

func (s *OASSchema) BuildNumberResource(name string, computability schema.ComputedOptionalRequired) (attrmapper.ResourceAttribute, *SchemaError) {
	if s.Format == util.OAS_format_double || s.Format == util.OAS_format_float {
		result := &attrmapper.ResourceFloat64Attribute{
			Name: name,
			Float64Attribute: resource.Float64Attribute{
				ComputedOptionalRequired: computability,
				DeprecationMessage:       s.GetDeprecationMessage(),
				Description:              s.GetDescription(),
			},
		}

		if s.Schema.Default != nil {
			var staticDefault float64
			if err := s.Schema.Default.Decode(&staticDefault); err == nil {
				if computability == schema.Required {
					result.ComputedOptionalRequired = schema.ComputedOptional
				}

				result.Default = &schema.Float64Default{
					Static: &staticDefault,
				}
			}
		}

		if computability != schema.Computed {
			result.Validators = s.GetFloatValidators()
		}

		return result, nil
	}

	return &attrmapper.ResourceNumberAttribute{
		Name: name,
		NumberAttribute: resource.NumberAttribute{
			ComputedOptionalRequired: computability,
			DeprecationMessage:       s.GetDeprecationMessage(),
			Description:              s.GetDescription(),
		},
	}, nil
}

func (s *OASSchema) BuildNumberDataSource(name string, computability schema.ComputedOptionalRequired) (attrmapper.DataSourceAttribute, *SchemaError) {
	if s.Format == util.OAS_format_double || s.Format == util.OAS_format_float {
		result := &attrmapper.DataSourceFloat64Attribute{
			Name: name,
			Float64Attribute: datasource.Float64Attribute{
				ComputedOptionalRequired: computability,
				DeprecationMessage:       s.GetDeprecationMessage(),
				Description:              s.GetDescription(),
			},
		}

		if computability != schema.Computed {
			result.Validators = s.GetFloatValidators()
		}

		return result, nil
	}
	result := &attrmapper.DataSourceNumberAttribute{
		Name: name,
		NumberAttribute: datasource.NumberAttribute{
			ComputedOptionalRequired: computability,
			DeprecationMessage:       s.GetDeprecationMessage(),
			Description:              s.GetDescription(),
		},
	}

	return result, nil
}

func (s *OASSchema) BuildNumberProvider(name string, optionalOrRequired schema.OptionalRequired) (attrmapper.ProviderAttribute, *SchemaError) {
	if s.Format == util.OAS_format_double || s.Format == util.OAS_format_float {
		result := &attrmapper.ProviderFloat64Attribute{
			Name: name,
			Float64Attribute: provider.Float64Attribute{
				OptionalRequired:   optionalOrRequired,
				DeprecationMessage: s.GetDeprecationMessage(),
				Description:        s.GetDescription(),
				Validators:         s.GetFloatValidators(),
			},
		}

		return result, nil
	}
	result := &attrmapper.ProviderNumberAttribute{
		Name: name,
		NumberAttribute: provider.NumberAttribute{
			OptionalRequired:   optionalOrRequired,
			DeprecationMessage: s.GetDeprecationMessage(),
			Description:        s.GetDescription(),
		},
	}

	return result, nil
}

func (s *OASSchema) BuildNumberElementType() (schema.ElementType, *SchemaError) {
	if s.Format == util.OAS_format_double || s.Format == util.OAS_format_float {
		return schema.ElementType{
			Float64: &schema.Float64Type{},
		}, nil
	}

	return schema.ElementType{
		Number: &schema.NumberType{},
	}, nil
}

func (s *OASSchema) GetFloatValidators() []schema.Float64Validator {
	var result []schema.Float64Validator

	if len(s.Schema.Enum) > 0 {
		var enum []float64

		for _, valueNode := range s.Schema.Enum {
			var value float64
			if err := valueNode.Decode(&value); err != nil {
				// could consider error/panic here to notify developers
				continue
			}

			enum = append(enum, value)
		}

		customValidator := frameworkvalidators.Float64ValidatorOneOf(enum)

		if customValidator != nil {
			result = append(result, schema.Float64Validator{
				Custom: customValidator,
			})
		}
	}

	return result
}
