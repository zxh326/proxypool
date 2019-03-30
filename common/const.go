package common

import (
	"time"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"
)

// Validator
const (
	VerifyUrl      = "http://httpbin.org/get"
	VerifyHttpsUrl = "https://httpbin.org/get"
	TimeOut        = 10 * time.Second
	ValidTimeOut   = 10 * time.Second
	ValidCount     = 3
)

// 当代理池中可用代理数据量小于此数值时才会触发更新任务
//
const MinPoolNum = 10


// 更新操作的间隔时间，建议不超过此数值
const NextValidTime = 5 * time.Minute


const APIPORT = ":8888"
