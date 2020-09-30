package handler

import (
	"net/http"

	"github.com/dwarvesf/gerr"
	"github.com/gin-gonic/gin"
	"github.com/hieuphq/backend-example/pkg/errors"
)

type orderItem struct {
	ProductID string `json:"productId" binding:"required"`
	Amount    int64  `json:"amount" binding:"gt=0"`
}
type orderReq struct {
	Items []orderItem `json:"items" binding:"gt=0,dive"`
}

type orderRes struct {
	Items []orderItem `json:"items" binding:"gt=0,dive"`
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req orderReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	res, err := h.doCreateOrder(req)
	if err != nil {
		h.handleError(c, gerr.E(500, gerr.Trace(err)))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) doCreateOrder(req orderReq) (*orderReq, error) {
	errs := []error{}
	for _, itm := range req.Items {
		err := h.doValidateCreateOrder(itm)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, gerr.E(1004, "call service is failed", errs)
	}

	return &req, nil
}

func (h *Handler) doValidateCreateOrder(itm orderItem) error {

	if itm.ProductID == "invalid" {
		return errors.ErrInvalidProduct.Err(gerr.Target("product_id"))
	}

	return nil

}
