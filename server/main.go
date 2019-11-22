package main


import (
	"github.com/gin-gonic/gin"
	
)


func main() {
	server := gin.Default()
	server.POST("/signup", SignUp)
	server.POST("/signein", SignIn)
}
