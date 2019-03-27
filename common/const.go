package common

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

var Request = gorequest.New()

func init()  {
	Request.Set("User-Agent", UserAgent)
	Request.Set("Cookie", "JSESSIONID=CDFE641DD40D95153418A4AA3C06A960")
	Request.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	Request.Timeout(TimeOut)
	fmt.Println(Request.Header)
}

const (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"
)


// Validator
const (
	VerifyUrl = "http://httpbin.org/get"
	VerifyHttpsUrl = "https://httpbin.org/get"
	TimeOut   = 10 * time.Second
	VaildTimeOut = 5 * time.Second
)

const(
	NextVaildTime = 5 * time.Minute
)
