package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"walletapi/pkg/database"
	"walletapi/pkg/errors"
)

// AdminCreateUser createUser?username=xxx&password=yyy
func AdminCreateUser(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	if username == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errors.ErrNoUsernameProvided.Error()})
		return
	}

	if password == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errors.ErrNoPasswordProvided.Error()})
		return
	}

	admin, err := database.CreateUser(username, password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Successfully created a new user" ,"createdUser": admin})
}

// AdminSetBalance setBalance?username=xxx?balance=100
func AdminSetBalance(c *gin.Context) {
	username := c.Request.FormValue("username")
	balanceStr := c.Request.FormValue("balance")

	if username == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errors.ErrNoUsernameProvided.Error()})
		return
	}

	if balanceStr == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errors.ErrNoBalanceProvided.Error()})
		return
	}


	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Can't parse numbers from the balance value"})
		return
	}

	targetUser, err := database.GetUserByUsername(username)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": errors.ErrAccountNotExist.Error()})
		return
	}

	oldBalance, newBalance := targetUser.SetBalance(balance)
	c.IndentedJSON(http.StatusAccepted, gin.H{"oldBalance": oldBalance, "newBalance": newBalance})
}


func AdminGetUser(c *gin.Context) {
	username := c.Request.FormValue("username")

	user, err := database.GetUserByUsername(username)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": errors.ErrAccountNotExist.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user.)
}


func initHandlersAdmin() {
	group := server.Group("/admin") 
	{
		group.GET("/getUser", AdminGetUser)
		group.GET("/createUser", AdminCreateUser)
		group.GET("/setBalance", AdminSetBalance)

	}
}