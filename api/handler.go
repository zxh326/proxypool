package api

import (
	"encoding/json"
	"github.com/zxh326/proxypool/proxy"
	"log"
	"net/http"
	"time"
)

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[API] %-4s  %-13s  %13v", r.Method, r.URL.Path, time.Since(start))
	})
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	apis := []string{
		"/",
		"/api/proxy",
		"/api/proxy/count",
	}

	js, err := json.Marshal(apis)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func allProxy(w http.ResponseWriter, r *http.Request) {
	querys := r.URL.Query()

	va, has := querys["protocol"]

	var protocol = "http"
	var res []*proxy.Proxy

	if has {
		if va[0] == "https" {
			protocol = va[0]
		}
		res = proxy.GetAll(protocol)
	} else {
		res = proxy.GetAll()
	}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func countProxy(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(map[string]int64{
		"count": proxy.Count(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
