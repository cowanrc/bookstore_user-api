package app

import (
	"bookstore_user-api/controllers/ping"
	"bookstore_user-api/controllers/user"
	"log"
)

func mapUrls() {
	r.GET("/ping", ping.Ping)

	r.GET("/users/:user_id", user.GetUser)
	// r.GET("/users/search", users.SearchUser)
	log.Printf("Calling Post")
	r.POST("/users", user.CreateUser)
}
