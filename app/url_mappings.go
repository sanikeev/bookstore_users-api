package app

import (
	"github.com/sanikeev/bookstore_users-api/controllers/ping"
	"github.com/sanikeev/bookstore_users-api/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
