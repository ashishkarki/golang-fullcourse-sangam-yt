package server

import (
	"04-notes-crud/internal/notes"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db *mongo.Database) *gin.Engine {
	router := gin.Default()

	// Global route so is put here within router.go file
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"status": "pong",
		})
	})

	// this is like app.use('/notes', notesRouter) from expressJs
	notes.RegisterRoutes(router, db)

	return router
}
