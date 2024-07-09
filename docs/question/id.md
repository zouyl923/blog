# 分布式ID

## 分布式规则
- 全局唯一：必须保证ID是全局性唯一的，基本要求。
- 高性能：高可用低延时，ID生成响应要块，否则反倒会成为业务瓶颈。
- 要秉着拿来即用的设计原则，在系统设计和实现上要尽可能的简单。
- 趋势递增：最好趋势递增，这个要求就得看具体业务场景了，一般不严格要求。

## 生成方案

### 数据库号段模式（推荐方案）
- 预先生成一段ID存放在缓存中，使用时直接使用缓存数据，
- 优点：可以自定义生成规则
```go
// UuidStep uuid生成配置表
type UuidStep struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	BizType   string    `gorm:"column:biz_type;type:varchar(32);not null;index:,unique;comment:业务类型" json:"biz_type"` // 业务类型
	Step      int32     `gorm:"column:step;type:int;not null;comment:每次生成多少;" json:"step"`                            // 每次生成多少
	MaxId     int64     `gorm:"column:max_id;type:int;not null;comment:上次生成最大ID;" json:"max_id"`                      // 当前最大ID
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt  time.Time `gorm:"column:update_at" json:"update_at"`
}
//
func GenUuid(in *rpc.GenUuidReq) (*rpc.GenUuidRes, error) {
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
    if len(bizId) < 1 {
        info := model.UuidStep{}
        err := l.svcCtx.DB.WithContext(l.ctx).
        Where("biz_type= ?", in.BizType).
        First(&info).Error
        if err != nil {
            return nil, err
        }
        //生成个数
        step := int(info.Step)
        maxId := 0
        if len(first) > 0 {
            //继承上波maxId
            maxId = int(info.MaxId)
        }
        //前缀
		prefixInt, _ := strconv.Atoi(date)
		//遍历生成uuid
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
        l.svcCtx.Cache.SetEx(firstKey, "ok")
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

```


### 雪花算法-Snowflake
- 雪花算法是由Twitter开源的分布式ID生成算法，以划分命名空间的方式将 64-bit位分割成多个部分，每个部分代表不同的含义。
- [参考说明](https://pdai.tech/md/arch/arch-z-id.html#%E9%9B%AA%E8%8A%B1%E7%AE%97%E6%B3%95-snowflake)
- 参考代码
```go
package uuid

import (
	"github.com/pkg/errors"
	"sync"
	"time"
)

// 雪花算法
const (
	workerBits  uint8 = 10                      // 节点数
	seqBits     uint8 = 12                      // 1毫秒内可生成的id序号的二进制位数
	workerMax   int64 = -1 ^ (-1 << workerBits) // 节点ID的最大值，用于防止溢出
	seqMax      int64 = -1 ^ (-1 << seqBits)    // 同上，用来表示生成id序号的最大值
	timeShift   uint8 = workerBits + seqBits    // 时间戳向左的偏移量
	workerShift uint8 = seqBits                 // 节点ID向左的偏移量
	epoch       int64 = 1713955600000           // 开始运行时间
)

type SnowFlake struct {
	// 添加互斥锁 确保并发安全
	mu sync.Mutex
	// 记录时间戳
	timestamp int64
	// 该节点的ID
	workerId int64
	// 当前毫秒已经生成的id序列号(从0开始累加) 1毫秒内最多生成4096个ID
	seq int64
}

// 实例化对象
func NewSnowFlake(workerId int64) (*SnowFlake, error) {
	// 要先检测workerId是否在上面定义的范围内
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	// 生成一个新节点
	return &SnowFlake{
		timestamp: 0,
		workerId:  workerId,
		seq:       0,
	}, nil
}

// 获取一个新ID
func (w *SnowFlake) Next() int64 {
	// 获取id最关键的一点 加锁 加锁 加锁
	w.mu.Lock()
	defer w.mu.Unlock() // 生成完成后记得 解锁 解锁 解锁
	// 获取生成时的时间戳
	now := time.Now().UnixNano() / 1e6 // 纳秒转毫秒
	if w.timestamp == now {
		w.seq = (w.seq + 1) & seqMax
		// 这里要判断，当前工作节点是否在1毫秒内已经生成seqMax个ID
		if w.seq == 0 {
			// 如果当前工作节点在1毫秒内生成的ID已经超过上限 需要等待1毫秒再继续生成
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 如果当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
		w.seq = 0
	}
	w.timestamp = now // 将机器上一次生成ID的时间更新为当前时间
	// 第一段 now - epoch 为该算法目前已经奔跑了xxx毫秒
	// 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
	ID := int64((now-epoch)<<timeShift | (w.workerId << workerShift) | (w.seq))
	return ID
}
```

### 其他方案
- [其他方案](https://pdai.tech/md/arch/arch-z-id.html#uuid)
