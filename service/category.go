package service

import (
	"context"
	"sync"

	"gin-mall-backend/repository/db/dao"
	"gin-mall-backend/types"
)

var CategorySrvIns *CategorySrv
var CategorySrvOnce sync.Once

type CategorySrv struct {
}

func GetCategorySrv() *CategorySrv {
	CategorySrvOnce.Do(func() {
		CategorySrvIns = &CategorySrv{}
	})
	return CategorySrvIns
}

// CategoryList 列举分类
func (s *CategorySrv) CategoryList(ctx context.Context, req *types.ListCategoryReq) (resp interface{}, err error) {
	categories, err := dao.NewCategoryDao(ctx).ListCategory()
	if err != nil {
		return
	}
	cResp := make([]*types.ListCategoryResp, 0)
	for _, v := range categories {
		cResp = append(cResp, &types.ListCategoryResp{
			ID:           v.ID,
			CategoryName: v.CategoryName,
			CreatedAt:    v.CreatedAt.Unix(),
		})
	}

	resp = &types.DataListResp{
		Item:  cResp,
		Total: int64(len(cResp)),
	}

	return
}
