package app

import (
	"bookstore_user-api/controllers/ping"
	"bookstore_user-api/controllers/user"
)

func mapUrls() {
	r.GET("/ping", ping.Ping)

	r.POST("/users", user.Create)
	r.GET("/users/:userId", user.Get)
	r.PUT("/users/:userId", user.Update)
	r.PATCH("/users/:userId", user.Update)
	r.DELETE("/users/:userId", user.Delete)
}
