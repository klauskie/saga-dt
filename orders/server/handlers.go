package server

import (
	"github.com/gin-gonic/gin"
	"github.com/klauskie/saga-dt/orders/config"
	"github.com/klauskie/saga-dt/orders/models"
	"github.com/klauskie/saga-dt/orders/service"
	"log"
	"net/http"
)

// Controllers or Handlers

// health godoc
func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "server is up"})
}

// submitOrder godoc
func submitOrder(env config.Env) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var order models.Order
		if err := c.BindJSON(&order); err != nil {
			setErrorMessage(c, http.StatusBadRequest, "unable to parse the given body")
			return
		}

		orderService := service.NewOrderService(env)
		if err := orderService.Submit(order); err != nil {
			log.Printf("error subbmiting order: %v", err)
			setErrorMessage(c, http.StatusInternalServerError, "unable to submit order")
			return
		}

		c.JSON(http.StatusOK, apiMessage{
			Status:  "OK",
			Message: "order submitted",
		})
	}
}

/* ERROR handlers */

func setErrorMessage(c *gin.Context, status int, errMsg string) {
	if status < 400 {
		status = http.StatusBadRequest
	}

	c.JSON(status, apiMessage{
		Status:  "error",
		Message: errMsg,
	})
}
