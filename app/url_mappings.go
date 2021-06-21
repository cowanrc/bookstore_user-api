package app

import (
	"bookstore_user-api/controllers/ping"
	"bookstore_user-api/controllers/user"
)

func mapUrls() {
	r.GET("/ping", ping.Ping)

	r.POST("/users", user.CreateUser)
	r.GET("/users/:userId", user.GetUser)
	r.PUT("/users/:userId", user.UpdateUser)
	r.PATCH("/users/:userId", user.UpdateUser)
}
