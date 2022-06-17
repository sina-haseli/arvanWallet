package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
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

		err = ch.service.Wallet.Increase(req.UserID, req.Amount, "Increase by something")
		if err != nil {
			fmt.Println("could not increase credit:", err)
		}
	}
}

func (ch *CreditHandler) HandleGetBalanceRequest() func(c echo.Context) error {
	return func(c echo.Context) error {
		i, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return echo.ErrBadRequest
		}

		b, err := ch.service.Wallet.GetBalance(i)
		if err != nil {
			return echo.ErrInternalServerError
		}

		_ = c.JSON(200, map[string]interface{}{"balance": b})
		return nil
	}
}
