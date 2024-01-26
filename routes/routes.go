package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //events/1, events/2
	
	//Option 1 by executing Authenticate method first and then redirecting it to other route.
	server.POST("/events", middlewares.Authenticate, createEvent)

	//Option 2 by creating a group then adding routes that group and adding authenticating function in that group that needs to be executed every time. 
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/deregister", cancelRegistration)
	
	server.POST("/signup", signup)
	server.POST("/login", login)
}