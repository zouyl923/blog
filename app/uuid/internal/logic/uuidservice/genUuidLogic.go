package uuidservicelogic

import (
	"blog/database/model"
	"context"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"blog/app/uuid/internal/svc"
	"blog/app/uuid/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenUuidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	lock sync.Mutex
}

func NewGenUuidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenUuidLogic {
	return &GenUuidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenUuidLogic) GenUuid(in *rpc.GenUuidReq) (*rpc.GenUuidRes, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	//是否第一次获取
	date := time.Now().Format("20060102")
	firstKey := "uuid:" + in.BizType + ":now:" + date
	key := "uuid:" + in.BizType + ":date:" + date
	first, _ := l.svcCtx.Cache.Get(firstKey)
	if len(first) < 1 {
		//每日第一拉取
		keys, _ := l.svcCtx.Cache.Keys("uuid:article:*")
		//移除先前的缓存
		for _, v := range keys {
			l.svcCtx.Cache.Del(v)
		}
	}
	bizId, _ := l.svcCtx.Cache.Rpop(key)
	if len(bizId) < 1 {
		info := model.UuidStep{}
		err := l.svcCtx.DB.WithContext(l.ctx).
			Where("biz_type= ?", in.BizType).
			First(&info).Error
		if err != nil {
			return nil, err
		}
		step := int(info.Step)
		maxId := 0
		if len(first) > 0 {
			//继承上波maxId
			maxId = int(info.MaxId)
		}
		prefixInt, _ := strconv.Atoi(date)

		for i := 0; i < step; i++ {
			//防止别人猜测uuid给上随机值
			maxId = maxId + rand.Intn(20)
			//格式：日期8为+8位
			//理论 每个业务 每日千万级别数据uuid 需要更多可以追加位数8+16满足基本上所有业务
			uuid := prefixInt*100000000 + maxId
			//放入队列
			l.svcCtx.Cache.Lpush(key, uuid)
		}
		//设置当天第一次更新OK
		l.svcCtx.Cache.Setex(firstKey, "ok", 24*60*60)
		//跟新记录到数据库
		l.svcCtx.DB.WithContext(l.ctx).Model(&info).
			Where("biz_type=?", in.BizType).
			Update("max_id", maxId)
		//返回第一个
		bizId, _ = l.svcCtx.Cache.Rpop(key)
	}
	return &rpc.GenUuidRes{
		BizType: in.BizType,
		BizId:   bizId,
	}, nil
}
