package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CategoriesModel = (*customCategoriesModel)(nil)

type (
	// CategoriesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCategoriesModel.
	CategoriesModel interface {
		categoriesModel
		FindByUserId(ctx context.Context, userId int64) ([]*Categories, error)
	}

	customCategoriesModel struct {
		*defaultCategoriesModel
	}
)

// NewCategoriesModel returns a model for the database table.
func NewCategoriesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CategoriesModel {
	return &customCategoriesModel{
		defaultCategoriesModel: newCategoriesModel(conn, c, opts...),
	}
}

// FindByUserId 查找指定用户的所有分类
func (m *customCategoriesModel) FindByUserId(ctx context.Context, userId int64) ([]*Categories, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? order by `created_at` desc", categoriesRows, m.table)
	var resp []*Categories
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return []*Categories{}, nil
	default:
		return nil, err
	}
}
