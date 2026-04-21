package core_server

import "net/http"

type HTTPServer struct {
	mux    *http.ServeMux
	config Config
}

func NewHttpServer(
	config Config,
) *HTTPServer {
	return &HTTPServer{
		mux:    http.NewServeMux(),
		config: config,
	}
}
