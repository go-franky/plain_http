package web

import "net/http"

func (s *Server) routes() {
	router := http.DefaultServeMux
	router.HandleFunc("/", s.log(s.helloWorld()))
	router.HandleFunc("/health", s.log(s.health()))
	s.Handler = router
}
