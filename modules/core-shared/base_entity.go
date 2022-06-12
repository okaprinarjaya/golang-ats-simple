package core_shared

import "time"

const (
	NONE     = "NONE"
	NEW      = "NEW"
	MODIFIED = "MODIFIED"
	DELETED  = "DELETED"
)

type BaseEntity struct {
	id                string
	createdAt         time.Time
	updatedAt         time.Time
	deletedAt         time.Time
	createdBy         string
	updatedBy         string
	deletedBy         string
	PersistenceStatus string
}

func (baseEnt *BaseEntity) Base(baseData BaseDTO) {
	baseEnt.PersistenceStatus = NONE

	if baseData.Id != "" {
		baseEnt.id = baseData.Id
	}
	if !baseData.CreatedAt.IsZero() {
		baseEnt.createdAt = baseData.CreatedAt
	}
	if !baseData.UpdatedAt.IsZero() {
		baseEnt.updatedAt = baseData.UpdatedAt
	}
	if !baseData.DeletedAt.IsZero() {
		baseEnt.deletedAt = baseData.DeletedAt
	}
	if baseData.CreatedBy != "" {
		baseEnt.createdBy = baseData.CreatedBy
	}
	if baseData.UpdatedBy != "" {
		baseEnt.updatedBy = baseData.UpdatedBy
	}
	if baseData.DeletedBy != "" {
		baseEnt.deletedBy = baseData.DeletedBy
	}
}

func (baseEnt *BaseEntity) Id() string {
	return baseEnt.id
}

func (baseEnt *BaseEntity) CreatedAt() time.Time {
	return baseEnt.createdAt
}

func (baseEnt *BaseEntity) UpdatedAt() time.Time {
	return baseEnt.updatedAt
}

func (baseEnt *BaseEntity) DeletedAt() time.Time {
	return baseEnt.deletedAt
}

func (baseEnt *BaseEntity) CreatedBy() string {
	return baseEnt.createdBy
}

func (baseEnt *BaseEntity) UpdatedBy() string {
	return baseEnt.updatedBy
}

func (baseEnt *BaseEntity) DeletedBy() string {
	return baseEnt.deletedBy
}
