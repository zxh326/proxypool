package proxy

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"log"
	"proxypool/pkg/http"
	"regexp"
	"strconv"
	"strings"
)

var request = gorequest.New()


func A2uProvider(ch chan<- *Proxy) {
	log.Printf("[%s]: provider crawler begin", "A2u")
	url := "https://proxy.rudnkh.me/txt"
	res, body, errs := request.Get(url).End()

	if errs != nil {
		log.Fatalf("[%s]: provider crawler error: %s", "A2u", errs)
	}

	defer res.Body.Close()
	f, _ := regexp.Compile("\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}:\\d{2,5}")

	ress := f.FindAllString(string(body), 20)
	for _, value := range ress {
		proxy := Proxy{}
		ip := strings.Split(value, ":")
		proxy.Ip = ip[0]
		proxy.Port = ip[1]
		proxy.Protocol = "http"
		proxy.refer = "A2u"
		ch <- &proxy
	}
	log.Printf("[%s] provider crawler done", "A2u")
}

func Data5uProvider(ch chan<- *Proxy) {
	url := "http://www.data5u.com/free/index.html"
	log.Printf("[%s]: provider crawler begin", "Data5u")
	res := http.HttpHandle(url, "Data5u")
	doc, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Printf("[%s]: crawler error : parse html error: %v", "Data5u", err)
		return
	}

	isIP,_ := regexp.Compile("\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}")
	doc.Find("body > div.wlist > ul > li:nth-child(2) > ul").Each(func(i int, s *goquery.Selection) {
		ip := s.Find("ul:nth-child(" + strconv.Itoa(i+1) + ") > span:nth-child(1) > li").Text()
		port := s.Find("ul:nth-child(" + strconv.Itoa(i+1) + ") > span:nth-child(2) > li").Text()
		protocol := s.Find("ul:nth-child(" + strconv.Itoa(i+1) + ") > span:nth-child(4) > li").Text()
		proxy := Proxy{}
		if isIP.MatchString(ip){
			proxy.Ip = ip
			proxy.Port = port
			proxy.Protocol = protocol
			proxy.refer = "Data5u"
			ch <- &proxy
		}
	})
	log.Printf("[%s]: provider crawler done", "Data5u")

}
