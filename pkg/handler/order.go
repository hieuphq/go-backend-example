package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderItem struct {
	ProductID string `json:"productId" binding:"required"`
	Amount    int64  `json:"amount" binding:"gt=0"`
}
type orderReq struct {
	Items []orderItem `json:"items" binding:"gt=0,dive"`
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req orderReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, req)
}
