package permission

import (
	"blog/common/helper"
	"blog/database/model"
	"context"

	"blog/app/admin/api/internal/svc"
	"blog/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TreeLogic {
	return &TreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TreeLogic) Tree() (resp []types.Menu, err error) {
	var list []model.AdminMenu
	l.svcCtx.DB.WithContext(l.ctx).
		Where("is_del=?", 0).
		Where("is_hid=?", 0).
		Preload("Permission").
		Order("weight desc").Find(&list)

	var tree []types.Menu
	helper.Swap(list, &tree)
	tree = GetTree(tree, 0)
	return tree, nil
}

func GetTree(list []types.Menu, pid int64) []types.Menu {
	tree := make([]types.Menu, 0)
	for _, v := range list {
		if v.ParentId == pid {
			v.Children = GetTree(list, v.Id)
			tree = append(tree, v)
		}
	}
	return tree
}
