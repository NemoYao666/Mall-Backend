package service

import (
	"context"
	"errors"
	"sync"

	conf "gin-mall-backend/config"
	"gin-mall-backend/consts"
	"gin-mall-backend/pkg/utils/ctl"
	"gin-mall-backend/repository/db/dao"
	"gin-mall-backend/repository/db/model"
	"gin-mall-backend/types"
)

var FavoriteSrvIns *FavoriteSrv
var FavoriteSrvOnce sync.Once

type FavoriteSrv struct {
}

func GetFavoriteSrv() *FavoriteSrv {
	FavoriteSrvOnce.Do(func() {
		FavoriteSrvIns = &FavoriteSrv{}
	})
	return FavoriteSrvIns
}

// FavoriteList 商品收藏夹
func (s *FavoriteSrv) FavoriteList(ctx context.Context, req *types.FavoritesServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	favorites, total, err := dao.NewFavoritesDao(ctx).ListFavoriteByUserId(u.Id, req.PageSize, req.PageNum)
	if err != nil {
		return
	}
	for i := range favorites {
		if conf.Config.System.UploadModel == consts.UploadModelLocal {
			favorites[i].ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + favorites[i].ImgPath
		}
	}

	resp = &types.DataListResp{
		Item:  favorites,
		Total: total,
	}

	return
}

// FavoriteCreate 创建收藏夹
func (s *FavoriteSrv) FavoriteCreate(ctx context.Context, req *types.FavoriteCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	fDao := dao.NewFavoritesDao(ctx)
	exist, _ := fDao.FavoriteExistOrNot(req.ProductId, u.Id)
	if exist {
		err = errors.New("已经存在了")
		return
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(u.Id)
	if err != nil {
		return
	}

	bossDao := dao.NewUserDaoByDB(userDao.DB)
	boss, err := bossDao.GetUserById(req.BossId)
	if err != nil {
		return
	}

	product, err := dao.NewProductDao(ctx).GetProductById(req.ProductId)
	if err != nil {
		return
	}

	favorite := &model.Favorite{
		UserID:    u.Id,
		User:      *user,
		ProductID: req.ProductId,
		Product:   *product,
		BossID:    req.BossId,
		Boss:      *boss,
	}
	err = fDao.CreateFavorite(favorite)
	if err != nil {
		return
	}

	return
}

// FavoriteDelete 删除收藏夹
func (s *FavoriteSrv) FavoriteDelete(ctx context.Context, req *types.FavoriteDeleteReq) (resp interface{}, err error) {
	favoriteDao := dao.NewFavoritesDao(ctx)
	err = favoriteDao.DeleteFavoriteById(req.Id)
	if err != nil {
		return
	}

	return
}
