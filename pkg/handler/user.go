package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signup create a user
func (h *Handler) Signup(c *gin.Context) {
	var req signupRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, req)
}
