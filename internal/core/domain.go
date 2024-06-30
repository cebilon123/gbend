package core

import (
	"context"
	"fmt"
)

type Validator interface {
	// ValidateEntity validates entity, should return nil if entity is valid.
	ValidateEntity(ctx context.Context, schema *EntitySchema, entity map[string]any) error
}

type DomainEntity struct {
	entity map[string]any
}

type DomainBuilder struct {
	validator Validator
}

func NewDomainBuilder(validator Validator) *DomainBuilder {
	return &DomainBuilder{
		validator: validator,
	}
}

func (d *DomainBuilder) BuildDomainEntity(ctx context.Context, schema *EntitySchema) (*DomainEntity, error) {
	if err := d.validator.ValidateEntity(ctx, *entity); err != nil {
		return nil, fmt.Errorf("validate entity: %w", err)
	}

	return &DomainEntity{
		entity: entity,
	}, nil
}
