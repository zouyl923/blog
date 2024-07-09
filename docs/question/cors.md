# 跨域解决方案

## nginx配置跨域
```conf
location / {  
    add_header Access-Control-Allow-Origin *;
    add_header Access-Control-Allow-Methods 'GET, POST, OPTIONS';
    add_header Access-Control-Allow-Headers 'DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization';

    if ($request_method = 'OPTIONS') {
        return 204;
    }
} 
```

## 设置header防止跨域
```go
//在入口函数中加入
var c config.Config
conf.MustLoad(*configFile, &c)
server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(func(header http.Header) {
    //自定义header
	header.Set("Access-Control-Allow-Origin", "*")
    header.Set("Access-Control-Allow-Headers", "*")
}, nil, "*"))
defer server.Stop()
```