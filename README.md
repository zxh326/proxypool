# ProxyPool

> 基于Go语言实现的代理池

## Design
- Api 
    * 使用 net/http 提供服务
    * `/`  Api列表
    * `/api/proxy?protocol=[http, https]` 所有proxy
    * `/api/proxy/count` proxy数量
    * `/api/proxy/random` 随机一个proxy
- common
    * 静态参数
- proxy
    * provider 代理供应
    * proxy model
- schedule
    * 控制爬取策略以及验证


## Const
```go
// Validator
const (
	VerifyUrl      = "http://httpbin.org/get"
	VerifyHttpsUrl = "https://httpbin.org/get"
	TimeOut        = 10 * time.Second
	ValidTimeOut   = 10 * time.Second
	ValidCount     = 3
)

// 系统每5分钟会重新验证一次数据库中的代理，并更延迟，
// 当池中代理小于 MinPoolNum 时才会重新爬取
const NextValidTime = 5 * time.Minute
const MinPoolNum = 50

// API port
const APIPORT = ":8888"

```

## Valid
系统会以设定time定时检查代理，不通过直接抛弃

latency `=` 三次验证的平均值


## Custom provider
* 实现一个方法向提供的channel写入代理即可
```go
func YourProvider(ch chan<- *Proxy) {
    ...
    
	ch <- &proxy
    ...
}
```
* 注册provider

`schedule/pool.go:52 CrawlerJob`
```go
...
providers := []func(ch chan<- *proxy.Proxy){
	proxy.Data5uProvider,
	proxy.A2uProvider,
    proxy.LiuLiuProvider,
    YourProvider,
}
...
```

## TODO
- [ ] 代理的级别控制
- [ ] 可视化界面
- [ ] 更完整的验证策略
- [ ] 自动感知provider
- [ ] 命令行的参数控制

## Thanks
* [Xorm](https://github.com/go-xorm/xorm)
* [go-sqlite3](github.com/mattn/go-sqlite3)
* [goquery](github.com/PuerkitoBio/goquery)


## License
Apache License 2.0. For more details, please read the
[LICENSE](https://github.com/zxh326/proxypool/blob/master/LICENSE) file.