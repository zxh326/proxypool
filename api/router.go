package api

import (
	"github.com/zxh326/proxypool/common"
	"log"
	"net/http"
)

type server struct {
}

func (s *server) listen(port string) {
	mux := http.NewServeMux()
	mux.Handle("/", loggingHandler(http.HandlerFunc(testHandler)))
	mux.Handle("/api/proxy", loggingHandler(http.HandlerFunc(allProxy)))
	mux.Handle("/api/proxy/count", loggingHandler(http.HandlerFunc(countProxy)))
	err := http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatalf("%v", err)
	}
}

func Start() {
	s := server{}
	s.listen(common.APIPORT)
	log.Println("listen server in ", common.APIPORT)
}
