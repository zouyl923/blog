# goctl模版工具

## 初始化模板
```bash
goctl template init
```

## 复制模板到项目
```bash
# 先获取版本号  
goctl template init 
# 上述命令可以看出版本号为1.6.6

# 复制到项目
# 进入此文档目录的同级目录
cp -R ~/.goctl/1.6.6/* ./1.6.6/
```

## 自定义模板
- 修改goctl/1.6.6中的代码可以自定义模版

## 使用模板
```bash
# 获取goctl模板地址
pwd 
# 得到地址/Users/zouyl/www/go/bim/goctl
# 在执行api命令时带上模板参数，即可使用自定义模板，否则使用
goctl api api.api --home="/Users/zouyl/www/go/bim/goctl/1.6.6" 
```

