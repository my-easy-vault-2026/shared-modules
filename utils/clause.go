package utils

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Clause interface {
	Scope() func(db *gorm.DB) *gorm.DB
}

type IgnoreClause struct {
}

var _ Clause = (*IgnoreClause)(nil)

func (IgnoreClause) Scope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.Insert{Modifier: "IGNORE"})
	}
}
