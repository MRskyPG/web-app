package handler

import (
	"fmt"
	"github.com/MRskyPG/web-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createPosition(c *gin.Context) {
	var workingPosition web.WorkingPosition

	if err := c.BindJSON(&workingPosition); err != nil {
		fmt.Printf("An error occured: failed to bind JSON [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.service.InsertPosition(&workingPosition)

	c.JSON(http.StatusOK, map[string]interface{}{
		"position_id": workingPosition.PositionID,
	})
}

func (h *Handler) getAllPositions(c *gin.Context) {
	positions := h.service.GetAllPositions()
	c.JSON(http.StatusOK, positions)
}

func (h *Handler) getPositionById(c *gin.Context) {
	position_id, err := strconv.Atoi(c.Param("position_id"))

	if err != nil {
		fmt.Printf("An error occured: failed to convert position_id to int [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var workingPosition web.WorkingPosition
	workingPosition, err = h.service.GetPosition(position_id)

	if err != nil {
		fmt.Printf("An error occured: [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, workingPosition)
}

func (h *Handler) updatePosition(c *gin.Context) {
	position_id, err := strconv.Atoi(c.Param("position_id"))
	if err != nil {
		fmt.Printf("An error occured: failed to convert position_id to int [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var position web.WorkingPosition
	if err := c.BindJSON(&position); err != nil {
		fmt.Printf("An error occured: failed to bind JSON [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	position.PositionID = position_id
	h.service.UpdatePosition(position_id, position)

	c.JSON(http.StatusOK, web.WorkingPosition{
		PositionID: position_id,
		Name:       position.Name,
	})
}

func (h *Handler) deletePosition(c *gin.Context) {
	position_id, err := strconv.Atoi(c.Param("position_id"))
	if err != nil {
		fmt.Printf("An error occured: failed to convert position_id to int [%s]", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.service.DeletePosition(position_id)
	c.String(http.StatusOK, "Position deleted")
}

func (h *Handler) createStaff(c *gin.Context) {

}

func (h *Handler) getAllStaff(c *gin.Context) {

}

func (h *Handler) getStaffById(c *gin.Context) {

}

func (h *Handler) updateStaff(c *gin.Context) {

}

func (h *Handler) deleteStaff(c *gin.Context) {

}
