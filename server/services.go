package main


import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type SignUpReq struct {
	User     string `json:user`
	Password string `json:password`
}

type SignInReq struct {
	User     string `json:user`
	Password string `json:password`
}


func SignUp(c *gin.Context) {
	var json SignUpReq
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are signed up."})
}


func SignIn(c *gin.Context) {
	var json SignInReq
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"status": "you are signed in."})
}



