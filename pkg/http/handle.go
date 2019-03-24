package http

import (
	"log"
	"net/http"
)

func HttpHandle(url, pname string) *http.Response {
	client := &http.Client{}

	var req *http.Request

	req, _ = http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36)")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Printf("[%s]: crawler error : status code error: %d %s", pname, res.StatusCode, res.Status)

	}
	return res
}
