package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	authUserId, found := c.Params.Get("auth_user_id")
	if !found || authUserId == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.JSON(http.StatusOK, StdResponse{
		Status:  "failure",
		Message: "msg",
	})
}

type StdResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}
