package handler

import (
	"fmt"
	"github.com/MRskyPG/web-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createOrder(c *gin.Context) {
	var order web.Order

	if err := c.BindJSON(&order); err != nil {
		fmt.Printf("An error occured: failed to bind JSON [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.service.InsertOrder(&order)

	c.JSON(http.StatusOK, map[string]interface{}{
		"order_id": order.OrderID,
	})
}

func (h *Handler) getAllOrders(c *gin.Context) {
	orders := h.service.GetAllOrders()
	c.JSON(http.StatusOK, orders)
}

func (h *Handler) getOrderById(c *gin.Context) {
	order_id, err := strconv.Atoi(c.Param("order_id"))

	if err != nil {
		fmt.Printf("An error occured: failed to convert order_id to int [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var order web.Order
	order, err = h.service.GetOrder(order_id)

	if err != nil {
		fmt.Printf("An error occured: [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) updateOrder(c *gin.Context) {
	order_id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		fmt.Printf("An error occured: failed to convert order_id to int [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var order web.Order
	if err := c.BindJSON(&order); err != nil {
		fmt.Printf("An error occured: failed to bind JSON [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	order.OrderID = order_id
	h.service.UpdateOrder(order_id, order)

	c.JSON(http.StatusOK, map[string]interface{}{
		"order_id": order_id,
	})
}

func (h *Handler) deleteOrder(c *gin.Context) {
	order_id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		fmt.Printf("An error occured: failed to convert order_id to int [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.service.DeleteOrder(order_id)
	c.String(http.StatusOK, "Order deleted")
}

func (h *Handler) getClientByOrderId(c *gin.Context) {
	order_id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		fmt.Printf("An error occured: failed to convert order_id to int [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var order web.Order
	order, err = h.service.GetOrder(order_id)

	if err != nil {
		fmt.Printf("An error occured: [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if order.ClientID == 0 {
		c.String(http.StatusOK, "This order doesn't have the client. It is just order.")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"client_id": order.ClientID,
	})
}
