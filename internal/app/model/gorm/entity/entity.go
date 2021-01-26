package entity

import (
	"context"
	"time"

	"github.com/gomaglev/microshop/v1/pkg/icontext"

	"gorm.io/gorm"
)

// Model base model
type Model struct {
	Id        string    `gorm:"column:id;primary_key;size:36;"`
	CreatedAt time.Time `gorm:"column:created_at;index;"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;"`
}

// GetDB ...
func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := icontext.FromTrans(ctx)
	if ok && !icontext.FromNoTrans(ctx) {
		db, ok := trans.(*gorm.DB)
		if ok {
			if icontext.FromTransLock(ctx) {
				db = db.Set("gorm:query_option", "FOR UPDATE")
			}
			return db
		}
	}
	return (*gorm.DB)(defDB)
}

// GetDBWithModel ...
func GetDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	db := GetDB(ctx, defDB)

	type tabler interface {
		TableName() string
	}

	if t, ok := m.(tabler); ok {
		return db.Table(t.TableName())

	}
	return db.Model(m)
}
