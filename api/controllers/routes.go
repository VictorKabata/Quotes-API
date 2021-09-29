package controllers

import "github.com/VictorKabata/quotes-api/api/middlewares"

func (server *Server) InitializeRoutes() {
	server.Router.HandleFunc("/", middlewares.JsonMiddleware(server.HomePage)).Methods("GET")

	server.Router.HandleFunc("/login", middlewares.JsonMiddleware(server.LoginUser)).Methods("POST")
	server.Router.HandleFunc("/register", middlewares.JsonMiddleware(server.CreateUser)).Methods("POST")

	server.Router.HandleFunc("/users", middlewares.JsonMiddleware(server.GetAllUsers)).Methods("GET")
	server.Router.HandleFunc("/user/{id}", middlewares.JsonMiddleware(server.GetUser)).Methods("GET")
	server.Router.HandleFunc("/user/{id}", middlewares.AuthenticationMiddleware(middlewares.JsonMiddleware(server.UpdateUser))).Methods("PUT")
	server.Router.HandleFunc("/user/{id}", middlewares.AuthenticationMiddleware(middlewares.JsonMiddleware(server.DeleteUser))).Methods("DELETE")

	server.Router.HandleFunc("/quote", middlewares.AuthenticationMiddleware(middlewares.JsonMiddleware(server.NewQuote))).Methods("POST")
	server.Router.HandleFunc("/quotes", middlewares.JsonMiddleware(server.GetAllQuotes)).Methods("GET")
	server.Router.HandleFunc("/quote/{id}", middlewares.JsonMiddleware(server.GetQuote)).Methods("GET")
	server.Router.HandleFunc("/quote/{id}", middlewares.AuthenticationMiddleware(middlewares.JsonMiddleware(server.UpdateQuotes))).Methods("PUT")
	server.Router.HandleFunc("/quote/{user_id}/{id}", middlewares.AuthenticationMiddleware(middlewares.JsonMiddleware(server.DeleteQuotes))).Methods("DELETE")
}
