package commonservicelogic

import (
	"blog/app/article/rpc/internal/contants"
	"blog/app/article/rpc/internal/svc"
	"blog/app/article/rpc/pb/rpc"
	"blog/common/helper"
	"blog/common/response/errx"
	"blog/database/model"
	"context"
	"encoding/json"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *rpc.InfoReq) (*rpc.Article, error) {
	info := model.Article{}
	key := contants.ArticleInfoKey + in.Uuid
	cInfo, _ := l.svcCtx.Cache.Get(key)
	if cInfo == "*" {
		//防止 缓存击穿
		return nil, errx.NewCodeError(errx.NotFundError)
	}
	//有真实数据
	if len(cInfo) > 1 {
		//解析缓存数据
		json.Unmarshal([]byte(cInfo), &info)
	} else {
		md := l.svcCtx.DB.WithContext(l.ctx)
		err := md.
			Where("uuid = ?", in.Uuid).
			Where("is_del = ?", 0).
			Where("is_hid = ?", 0).
			Where("state = ?", 1).
			Preload("DetailInfo").
			Preload("CategoryInfo").
			Preload("UserInfo").
			First(&info).Error
		if err != nil {
			//防止击穿 缓存60s 防止一直请求数据库
			l.svcCtx.Cache.Setex(key, "*", 60)
			return nil, errx.NewCodeError(errx.NotFundError)
		}
		var viewNum int
		viewNumCache, _ := l.svcCtx.Cache.Get(contants.ArticleViewNUmKey + in.Uuid)
		if len(viewNumCache) > 0 {
			viewNum, _ = strconv.Atoi(viewNumCache)
			viewNum++
		}
		//缓存数据
		ccInfo, _ := json.Marshal(info)
		l.svcCtx.Cache.Setex(key, string(ccInfo), 24*60*60)
		//更新浏览数
		info.ViewNum = info.ViewNum + int32(viewNum)
		//记录访问数，异步跟更新访问数到数据库
		l.svcCtx.Cache.Incr(contants.ArticleViewNUmKey + in.Uuid)
		//异步到kafka 然后更新数据库

		l.svcCtx.KqViewNumProduct.Push(info.Uuid)
	}
	rInfo := rpc.Article{}
	helper.ExchangeStruct(info, &rInfo)
	return &rInfo, nil
}
