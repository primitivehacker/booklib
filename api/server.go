package api

import "net/http"

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/book", s.handleGetBookByID)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetBookByID(w http.ResponseWriter, r *http.Request) {

}
