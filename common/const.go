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
	ValidTimeOut   = 5 * time.Second
)

const (
	NextValidTime = 5 * time.Minute
)
