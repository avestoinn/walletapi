package server

import "github.com/gin-gonic/gin"


var server = gin.Default()


func Run() error {
	// Initializing handlers before start the server
	initHandlers_Admin()

	err := server.Run("localhost:8000")
	if err != nil {
		return err
	}
	return nil
}
