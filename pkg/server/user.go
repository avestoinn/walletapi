package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"walletapi/pkg/database"
	"walletapi/pkg/errors"
)


func UserGenerateToken(c *gin.Context) {
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

	user, err := database.TryAuth(username, password)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": errors.ErrIncorrectPassword.Error()})
		return
	}

	generatedToken, err := user.GenerateToken()
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": generatedToken,
		"msg": "Please, don't share the token with third parties!"})
}


func UserGetWallets(c *gin.Context) {
	token := c.Request.FormValue("token")

	if token == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidToken.Error()})
		return
	}

	user, err := database.GetUserByToken(token)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"totalBalance": user.TotalBalance(),
		"wallets": user.Wallets})
}



func initHandlersUser() {
	group := server.Group("/user")
	{
		group.GET("/genToken", UserGenerateToken)
		group.GET("/wallets", UserGetWallets)

	}
}