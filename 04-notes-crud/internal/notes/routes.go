package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(r *gin.Engine, db *mongo.Database) {
	// create new repo and new handler once at the startup
	repo := NewRepo(db)
	handler := NewHandler(repo)

	notesGroup := r.Group("/notes")
	{
		notesGroup.POST("", handler.CreateNote)
	}
}
