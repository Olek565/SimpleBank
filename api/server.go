package api

import (
	db "github.com/Olek565/SimpleBank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server that servers the HTTP request for the bank service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new Server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	//adding routes to the router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start runs the HTTP server at the specified address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Method used to parse the error as JSON
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
