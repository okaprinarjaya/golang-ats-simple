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
	createdByName     string
	updatedBy         string
	updatedByName     string
	deletedBy         string
	deletedByName     string
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
	if baseData.CreatedByName != "" {
		baseEnt.createdByName = baseData.CreatedByName
	}
	if baseData.UpdatedBy != "" {
		baseEnt.updatedBy = baseData.UpdatedBy
	}
	if baseData.UpdatedByName != "" {
		baseEnt.updatedByName = baseData.UpdatedByName
	}
	if baseData.DeletedBy != "" {
		baseEnt.deletedBy = baseData.DeletedBy
	}
	if baseData.DeletedByName != "" {
		baseEnt.deletedByName = baseData.DeletedByName
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

func (baseEnt *BaseEntity) CreatedByName() string {
	return baseEnt.createdByName
}

func (baseEnt *BaseEntity) UpdatedByName() string {
	return baseEnt.updatedByName
}

func (baseEnt *BaseEntity) DeletedByName() string {
	return baseEnt.deletedByName
}

func (baseEnt *BaseEntity) SetUpdatedAt(updatedAt time.Time) {
	baseEnt.updatedAt = updatedAt
}

func (baseEnt *BaseEntity) SetDeleteddAt(deletedAt time.Time) {
	baseEnt.deletedAt = deletedAt
}

func (baseEnt *BaseEntity) SetUpdatedBy(updatedBy string) {
	baseEnt.updatedBy = updatedBy
}

func (baseEnt *BaseEntity) SetUpdatedByName(updatedByName string) {
	baseEnt.updatedByName = updatedByName
}

func (baseEnt *BaseEntity) SetDeletedBy(deletedBy string) {
	baseEnt.deletedBy = deletedBy
}

func (baseEnt *BaseEntity) SetDeletedByName(deletedByName string) {
	baseEnt.deletedByName = deletedByName
}
