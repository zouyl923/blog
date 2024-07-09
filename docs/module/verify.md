# 验证模块

## 生成rpc
```bash
# 进入项目目录
# 创建模块
mkdir -p app/verify/rpc
# 进入接口模块
cd app/verify/rpc
# 编辑 user.proto
# 编辑.proto 语法参考@see https://go-zero.dev/docs/tasks/dsl/proto
vim verify.proto 
# 执行.proto文件
# -dir 指定生成目录  
# --style 生成代码风格 goZero 小驼峰写法
# --home 生成代码采用的模板 
# -m 采用分组模式
# 具体参数 @see https://go-zero.dev/docs/tutorials/cli/rpc
goctl rpc protoc ./verify.proto  --go_out=./pb --go-grpc_out=./pb --zrpc_out=.  --style="goZero" --home="../../../template" -m 

```