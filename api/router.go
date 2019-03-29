package api

import (
	"log"
	"net/http"
	"proxypool/common"
)

type server struct {

}

func (s *server) listen (port string)  {
	mux := http.NewServeMux()
	mux.Handle("/", loggingHandler(http.HandlerFunc(testHandler)))
	mux.Handle("/api/proxy", loggingHandler(http.HandlerFunc(allProxy)))
	mux.Handle("/api/count", loggingHandler(http.HandlerFunc(countProxy)))
	_ = http.ListenAndServe(port, mux)
}


func Start()  {
	s := server{}
	s.listen(common.APIPORT)
	log.Println("listen server in ", common.APIPORT)
}