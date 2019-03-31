package proxy

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"github.com/zxh326/proxypool/common"
	"log"
	"regexp"
	"strconv"
	"strings"
)


func A2uProvider(ch chan<- *Proxy) {
	log.Printf("[Crawler] %s provider crawler begin", "A2u")
	url := "https://proxy.rudnkh.me/txt"
	res, body, errs := gorequest.New().Get(url).Set("User-Agent", common.UserAgent).Timeout(common.TimeOut).End()

	if errs != nil {
		log.Printf("[Crawler] %s provider crawler error: %s", "A2u", errs)
	}
	if res.StatusCode != 200 {
		log.Printf("[Crawler] %s provider retun status code error %s", "A2u", errs)
	}

	f, _ := regexp.Compile("\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}:\\d{2,5}")

	ress := f.FindAllString(string(body), 20)
	for _, value := range ress {
		proxy := Proxy{}
		ip := strings.Split(value, ":")
		proxy.Ip = ip[0]
		proxy.Port = ip[1]
		proxy.Protocol = "http"
		proxy.Refer = "A2u"
		ch <- &proxy
	}
	defer res.Body.Close()
	log.Printf("[Crawler] %s provider crawler done (%d)", "A2u", len(ress))
}

func Data5uProvider(ch chan<- *Proxy) {
	log.Printf("[Crawler] %s provider crawler begin", "Data5u")
	url := "http://www.data5u.com/free/index.html"
	res, _, errs := gorequest.New().Get(url).Set("User-Agent", common.UserAgent).Timeout(common.TimeOut).Retry(3, common.TimeOut).End()
	if errs != nil {
		log.Printf("[Crawler] %s provider crawler error: %s", "Data5u", errs)
	}

	if res.StatusCode != 200 {
		log.Printf("[Crawler] %s provider retun status code error %s", "Data5u", errs)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("[Crawler] %s crawler error : parse html error: %v", "Data5u", err)
		return
	}

	count := 0
	isIP, _ := regexp.Compile("\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}")
	doc.Find("body > div.wlist > ul > li:nth-child(2) > ul").Each(func(i int, s *goquery.Selection) {
		ip := s.Find("ul:nth-child(" + strconv.Itoa(i+1) + ") > span:nth-child(1) > li").Text()
		port := s.Find("ul:nth-child(" + strconv.Itoa(i+1) + ") > span:nth-child(2) > li").Text()
		protocol := s.Find("ul:nth-child(" + strconv.Itoa(i+1) + ") > span:nth-child(4) > li").Text()
		proxy := Proxy{}
		if isIP.MatchString(ip) {
			proxy.Ip = ip
			proxy.Port = port
			proxy.Protocol = strings.ToLower(protocol)
			proxy.Refer = "Data5u"
			ch <- &proxy
			count++
		}
	})

	defer res.Body.Close()
	log.Printf("[Crawler] %s provider crawler done (%d)", "Data5u", count)

}

func LiuLiuProvider(ch chan<- *Proxy) {
	log.Printf("[Crawler] %s provider crawler begin", "66Ip")
	url := "http://www.66ip.cn/mo.php?tqsl=100"
	res, body, errs := gorequest.New().Get(url).Set("User-Agent", common.UserAgent).Timeout(common.TimeOut).End()

	if errs != nil {
		log.Printf("[Crawler] %s provider crawler error: %s", "66Ip", errs)
	}
	if res.StatusCode != 200 {
		log.Printf("[Crawler] %s provider retun status code error %s", "66Ip", errs)
	}

	f, _ := regexp.Compile("\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}:\\d{2,5}")

	ress := f.FindAllString(string(body), 100)
	for _, value := range ress {
		proxy := Proxy{}
		ip := strings.Split(value, ":")
		proxy.Ip = ip[0]
		proxy.Port = ip[1]
		proxy.Protocol = "http"
		proxy.Refer = "66Ip"
		ch <- &proxy
	}
	defer res.Body.Close()
	log.Printf("[Crawler] %s provider crawler done (%d)", "66Ip", len(ress))
}
