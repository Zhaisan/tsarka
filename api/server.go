package api

import (
	db "github.com/Zhaisan/tsarka_test/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *db.Queries
	router *gin.Engine
}

func NewServer(db *db.Queries) *Server{
	server := &Server{
		db:    db,
	}
	router := gin.Default()

	// task1:
	router.POST("/rest/substr/find", server.findSubstring)
	router.POST("/rest/strings", server.createString)
	router.GET("/rest/strings/:id", server.getString)
	router.GET("/rest/strings", server.listStrings)
	router.PUT("/rest/strings/:id", server.updateString)
	router.DELETE("/rest/strings/:id", server.deleteString)

	// task2:
	router.POST("/rest/email/check", server.checkEmailsHandler)
	router.POST("/rest/iin/check", server.checkIINHandler)

	// task3:
	router.GET("/rest/self/find/handler", server.findIdentifiersHandler)


	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
