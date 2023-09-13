package handler

import (
	"net/http"

	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	response "github.com/bootcamp-go/desafio-go-web/pkg/web"
	"github.com/gin-gonic/gin"
)

type HandlerTicket struct {
	service tickets.ServiceTicket
}

func NewHandlerTicket(service tickets.ServiceTicket) *HandlerTicket {
	return &HandlerTicket{
		service: service,
	}
}

func (h *HandlerTicket) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		destination := c.Param("dest")
		tickets, err := h.service.GetTotalTickets(c, destination)

		if err != nil {
			response.Failure(c, http.StatusNotFound, nil)
			return
		}

		response.Success(c, http.StatusOK, tickets)
		return
	}
}

func (h *HandlerTicket) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {
		destination := c.Param("dest")
		avg, err := h.service.AverageDestination(c, destination)

		if err != nil {
			response.Failure(c, http.StatusNotFound, nil)
			return
		}

		response.Success(c, http.StatusOK, avg)
		return
	}
}
