package server

func (s *Server) routes() {
	s.Router.HandleFunc("/signup", s.handleSignup()).Methods("POST")
	s.Router.HandleFunc("/login", s.handleLogin()).Methods("POST")
	s.Router.HandleFunc("/logout", s.LogoutHandler)
	s.Router.HandleFunc("/", s.auth(s.handleHome(templatesDir)))
}
