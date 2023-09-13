package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type RouterTicket struct {
	r       *gin.Engine
	handler handler.HandlerTicket
}

func NewRouterTicket(r *gin.Engine, dbPath string) *RouterTicket {
	repo := tickets.NewRepositoryTicket(dbPath)
	service := tickets.NewServiceTicket(repo)
	handler := handler.NewHandlerTicket(service)

	return &RouterTicket{
		r:       r,
		handler: *handler,
	}
}

func (rt *RouterTicket) Run() {
	rt.mapRoutes()

	err := rt.r.Run()

	if err != nil {
		panic(err)
	}
}

func (rt *RouterTicket) mapRoutes() {
	routerGroup := rt.r.Group("/ticket")
	{
		routerGroup.GET("/getByCountry/:dest", rt.handler.GetTicketsByCountry())
		routerGroup.GET("/getAverage/:dest", rt.handler.AverageDestination())
	}
}
