package handler

import (
	"fmt"
	"github.com/MRskyPG/web-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createClient(c *gin.Context) {
	var client web.Client

	if err := c.BindJSON(&client); err != nil {
		fmt.Printf("An error occured: failed to bind JSON [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.service.InsertClient(&client)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": client.ID,
	})
}

func (h *Handler) getAllClients(c *gin.Context) {
	clients := h.service.GetAllClients()
	c.JSON(http.StatusOK, clients)
}

func (h *Handler) getClientById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Printf("An error occured: failed to convert id to int [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var client web.Client
	client, err = h.service.GetClient(id)

	if err != nil {
		fmt.Printf("An error occured: [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (h *Handler) updateClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("An error occured: failed to convert id to int [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var client web.Client
	if err := c.BindJSON(&client); err != nil {
		fmt.Printf("An error occured: failed to bind JSON [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	client.ID = id
	h.service.UpdateClient(id, client)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("An error occured: failed to convert id to int [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.service.DeleteClient(id)
	c.String(http.StatusOK, "Client deleted")
}

func (h *Handler) getOrderByClientId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("An error occured: failed to convert id to int [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var client web.Client
	client, err = h.service.GetClient(id)

	if err != nil {
		fmt.Printf("An error occured: [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if client.OrderID == 0 {
		c.String(http.StatusOK, "This client doesn't have the order.")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"order_id": client.OrderID,
	})
}
