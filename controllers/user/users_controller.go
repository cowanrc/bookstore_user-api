package user

import (
	"bookstore_user-api/domain/users"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	log.Printf("Printing user %v:", user)
	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println(err)
	}

	log.Println(string(bytes))
	c.String(http.StatusNotImplemented, "implement me!")

}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }
