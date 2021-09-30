package main

import (
	_ "github.com/lib/pq"

	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "net/http"
	user2 "task/backend/user"
)

func main() {
	router := gin.Default()

	// Adding ability to connect html and static files
	router.LoadHTMLGlob("frontend/html/*.html")
	router.Static("/assets/", "frontend/")

	// processing main '/' page which is going to be authorization field.
	router.GET("/", handlerAuthPg)

	// Here we are trying to authorize
	router.POST("/user/authorization", handleAutorizationAttempt)

	router.Run()
}

// function that processes '/'
func handlerAuthPg(c *gin.Context) {

	c.HTML(200, "autorization.html", nil)
}

/* here we bind info from inputs using JSON to struct user. so that later on we can use
it to send data to our database, which we have not created yet.
*/
func handleAutorizationAttempt(c *gin.Context) {
	var user user2.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(501, gin.H{
			"Error": err.Error(),
		})
		return
	}

	/*
	   TD: 1- connect database to the server.
	   	2- make it possible to transfer data.
	   	and remake that function.
	*/

}
