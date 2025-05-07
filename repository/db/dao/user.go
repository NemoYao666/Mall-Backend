package dao

import (
	"context"

	"gorm.io/gorm"

	"gin-mall-backend/pkg/utils/log"
	"gin-mall-backend/repository/db/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// FollowUser userId 关注了 followerId, 添加一条{followerId, uId}
func (dao *UserDao) FollowUser(uId, followerId uint) (err error) {
	u, f := new(model.User), new(model.User)
	dao.DB.Model(&model.User{}).Where(`id = ?`, uId).First(&u)
	dao.DB.Model(&model.User{}).Where(`id = ?`, followerId).First(&f)
	err = dao.DB.Model(&u).Association(`Relations`).
		Append([]model.User{*f})
	if err != nil {
		log.LogrusObj.Error(err)
		return err
	}

	return
}

// UnFollowUser 不再关注
func (dao *UserDao) UnFollowUser(uId, followerId uint) (err error) {
	u, f := new(model.User), new(model.User)
	dao.DB.Model(&model.User{}).Where(`id = ?`, uId).First(&u)
	dao.DB.Model(&model.User{}).Where(`id = ?`, followerId).First(&f)
	err = dao.DB.Model(&u).Association(`Relations`).Delete(f)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	return
}

// ListFollowing 用户关注的人（关注）
func (dao *UserDao) ListFollowing(userId uint, start, limit int) (f []*model.User, err error) {
	err = dao.DB.Table("user").
		Select("user.*").
		Joins("INNER JOIN relation ON relation.relation_id = user.id").
		Where("relation.user_id = ?", userId).
		Offset(start).Limit(limit).
		Find(&f).Error

	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	return
}

// ListFollower 关注用户的人（粉丝）
func (dao *UserDao) ListFollower(userId uint, start, limit int) (f []*model.User, err error) {

	err = dao.DB.Table("user").
		Select("user.*").
		Joins("INNER JOIN relation ON relation.user_id = user.id"). // 这里应该是粉丝的 `user_id`
		Where("relation.relation_id = ?", userId).                  // 这里是 `relation_id`，表示被关注的人
		Offset(start).Limit(limit).
		Find(&f).Error

	if err != nil {
		log.LogrusObj.Errorf("ListFollower 查询失败: %v", err)
		return nil, err
	}

	return
}

// GetUserById 根据 id 获取用户
func (dao *UserDao) GetUserById(uId uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uId).
		First(&user).Error
	return
}

// UpdateUserById 根据 id 更新用户信息
func (dao *UserDao) UpdateUserById(uId uint, user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(&user).Error
}

// ExistOrNotByUserName 根据username判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}
