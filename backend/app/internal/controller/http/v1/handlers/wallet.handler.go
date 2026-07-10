package handlers

import (
	"backend.app/internal/models"
	"github.com/gin-gonic/gin"
)

// @Tags Wallet
// @Summary Get User Wallet with balance
// @Description Get User Wallet with user
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ResponseObject{data=dtos.Wallet} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /wallets [get]
func (h *Handler) GetUserWallet(c *gin.Context) {
	user := c.MustGet("authUser").(models.User) // auth user
	result := h.core.GetUserWallet(c, &user)
	c.JSON(result.Code, result)
}
