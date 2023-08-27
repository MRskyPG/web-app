package handler

import (
	"github.com/MRskyPG/web-app/internal/service"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	service service.Service
}

func New(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		wPositions := api.Group("/working_positions")
		{
			wPositions.POST("", h.createPosition)
			wPositions.GET("", h.getAllPositions)
			wPositions.GET("/:position_id", h.getPositionById)
			wPositions.PUT("/:position_id", h.updatePosition)
			wPositions.DELETE("/:position_id", h.deletePosition)

			staff := wPositions.Group(":position_id/staff")
			{
				staff.POST("", h.createStaff)
				staff.GET("/:staff_id", h.getStaffById)
				staff.PUT("/:staff_id", h.updateStaff)
				staff.DELETE("/:staff_id", h.deleteStaff)
			}
		}
		staff := api.Group("/staff")
		{
			staff.GET("", h.getAllStaff)
		}

		clients := api.Group("/clients")
		{
			clients.POST("", h.createClient)
			clients.GET("", h.getAllClients)
			clients.GET("/:id", h.getClientById)
			clients.PUT("/:id", h.updateClient)
			clients.DELETE("/:id", h.deleteClient)

			order := clients.Group(":id/order")
			{
				order.GET("", h.getOrderByClientId)
			}
		}

		orders := api.Group("/orders")
		{
			orders.POST("", h.createOrder)
			orders.GET("", h.getAllOrders)
			orders.GET("/:order_id", h.getOrderById)
			orders.PUT("/:order_id", h.updateOrder)
			orders.DELETE("/:order_id", h.deleteOrder)

			client := orders.Group(":order_id/client")
			{
				client.GET("", h.getClientByOrderId)
			}
		}
	}

	return router
}
