package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/gomaglev/microshop/pkg/icontext"

	"github.com/gomaglev/protos/pkg/proto/common"

	"gorm.io/gorm"
)

const (
	maxPageSize = 1000
)

// TransFunc
type TransFunc func(context.Context) error

// ExecTrans
func ExecTrans(ctx context.Context, db *gorm.DB, fn TransFunc) error {
	transModel := &Trans{DB: db}
	return transModel.Exec(ctx, fn)
}

// ExecTransWithLock
func ExecTransWithLock(ctx context.Context, db *gorm.DB, fn TransFunc) error {
	if !icontext.FromTransLock(ctx) {
		ctx = icontext.NewTransLock(ctx)
	}
	return ExecTrans(ctx, db, fn)
}

// WrapPageQuery
func WrapPageQuery(ctx context.Context,
	db *gorm.DB,
	pp *common.PaginationParam,
	opt *common.QueryOptions,
	out interface{},
) (*common.PaginationResult, error) {
	if pp == nil {
		pp = &common.PaginationParam{
			Pagination: false,
			Page:       1,
			PageSize:   maxPageSize,
		}
	}
	if pp.CountOnly {
		var count int64
		err := db.Count(&count).Error
		if err != nil {
			return nil, err
		}
		return &common.PaginationResult{Total: count}, nil
	}

	total, err := FindPage(ctx, db, pp, opt, out)
	if err != nil {
		return nil, err
	}

	return &common.PaginationResult{
		Total:    total,
		Page:     int32(pp.Page),
		PageSize: int32(pp.PageSize),
		Cursor:   pp.Cursor,
	}, nil
}

// FindPage
func FindPage(ctx context.Context,
	db *gorm.DB,
	pp *common.PaginationParam,
	opt *common.QueryOptions,
	out interface{}) (int64, error) {
	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	} else if count == 0 {
		return count, nil
	}

	if len(opt.OrderByFields) > 0 {
		db = db.Order(ParseOrder(opt.OrderByFields))
	}

	current, pageSize := pp.GetPage(), pp.GetPageSize()
	if current > 0 && pageSize > 0 {
		if len(opt.SelectFields) > 0 {
			db = db.Select(opt.SelectFields).Offset(int((current - 1) * pageSize)).Limit(int(pageSize))
		} else {
			db = db.Offset(int((current - 1) * pageSize)).Limit(int(pageSize))
		}
	} else if pageSize > 0 {
		if len(opt.SelectFields) > 0 {
			db = db.Select(opt.SelectFields).Limit(int(pageSize))
		} else {
			db = db.Limit(int(pageSize))
		}
	}

	err = db.Find(out).Error
	return count, err
}

// FindOne
func FindOne(ctx context.Context, db *gorm.DB, out interface{}) (bool, error) {
	result := db.First(out)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Check
func Check(ctx context.Context, db *gorm.DB) (bool, error) {
	var count int64
	result := db.Count(&count)
	if err := result.Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// OrderByFieldFunc
type OrderByFieldFunc func(string) string

// ParseOrder
func ParseOrder(items []*common.OrderByField) string {
	orders := make([]string, len(items))

	for i, item := range items {
		key := item.Key
		direction := "ASC"
		if item.Direction == common.OrderBy_DESC {
			direction = "DESC"
		}
		orders[i] = fmt.Sprintf("%s %s", key, direction)
	}

	return strings.Join(orders, ",")
}

// NewOrderByField
func NewOrderByField(key string, d common.OrderBy) *common.OrderByField {
	return &common.OrderByField{
		Key:       key,
		Direction: d,
	}
}

// UnitStrToSet
func UnitStrToSet(units string) map[interface{}]bool {
	targetUnits := make(map[interface{}]bool)
	if units != "" {
		for _, u := range strings.Split(units, ",") {
			targetUnits[u] = true
		}
	}
	return targetUnits
}

// GetQueryOption
func GetQueryOption(opts ...*common.QueryOptions) *common.QueryOptions {
	var opt *common.QueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = &common.QueryOptions{}
	}
	return opt
}
