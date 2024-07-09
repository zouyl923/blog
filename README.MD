# 项目架构

## 项目简介
- 基于go-zero的管理后台项目。
- 包含微服务解决方案，高并发解决方案

## 框架选择
- go-zero 微服务框架
- kafka 消息中间件

## 目录说明
- [目录结构](docs/menu.md)

## 功能模块
- [管理后台模块](docs/module/admin.md)
- [文章模块](docs/module/article.md)
- [评论模块](docs/module/comment.md)
- [点赞模块](docs/module/like.md)
- [消息模块](docs/module/message.md)
- [用户模块](docs/module/user.md)
- [分布式ID模块](docs/module/uuid.md)
- [认证模块](docs/module/verify.md)


## 分布式系统
- 全局唯一ID
- 分布式锁
- 分布式事务
- 分布式缓存
- 分布式定时任务
- 分布式登录态
- 分布式日志

## 高并发系统
- 缓存
- 限流
- 降级与熔断
- 负载均衡
- 容灾备份，故障转移

## 解决方案
- [跨域如何解决？](docs/question/cors.md)
- [分布式ID如何设计？](docs/question/id.md)
- [高并发-如何更新阅读数？](docs/question/viewNum.md)
- [高并发-缓存雪崩、击穿、穿透？](docs/question/viewNum.md)