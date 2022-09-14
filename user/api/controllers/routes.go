package controllers

import "github.com/richard-here/imx-ilv-land-hooks/user/api/middlewares"

func (s *Server) initializeRoutes() {
	// Login
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// Users
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")

	// Subscriptions
	s.Router.HandleFunc("/subscriptions", middlewares.SetMiddlewareJSON(s.CreateSubscription)).Methods("POST")
	s.Router.HandleFunc("/subscriptions", middlewares.SetMiddlewareJSON(s.GetAllSubscriptions)).Methods("GET")
	s.Router.HandleFunc("/subscriptions/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteSubscription)).Methods("DELETE")
}
