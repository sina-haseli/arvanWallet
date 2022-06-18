package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"wallet/models"
	"wallet/services"
)

type CreditHandler struct {
	service *services.Services
}

func newCreditHandler(service *services.Services) *CreditHandler {
	return &CreditHandler{service: service}
}

func (ch *CreditHandler) HandleIncreaseRequestFromChannel(requestChannel chan string) {
	fmt.Println("Initialize Handler")
	for r := range requestChannel {

		fmt.Println("new request received: ", r)
		var req models.IncreaseRequestModel
		err := json.Unmarshal([]byte(r), &req)
		if err != nil {
			fmt.Println("could not parse json:", err)
		}

		ctx := context.Background()

		err = ch.service.Wallet.Increase(ctx, req.UserID, req.Amount, "Increase by something")
		if err != nil {
			fmt.Println("could not increase credit:", err)
		}
	}
}

// HandleGetBalanceRequest
// @Summary HandleGetBalanceRequest
// @Tags HandleGetBalanceRequest
// @Accept       json
// @Produce json
// @Param user_id path string true "user_id"
// @Success 200 {object} map[string]interface{}
// @Router /api/balance/{user_id} [get]
func (ch *CreditHandler) HandleGetBalanceRequest() func(c echo.Context) error {
	return func(c echo.Context) error {
		i := c.Param("user_id")

		b, err := ch.service.Wallet.GetBalance(c.Request().Context(), i)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(200, map[string]interface{}{"balance": b})
	}
}
