package notes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (handler *Handler) CreateNote(ginCtx *gin.Context) {
	var request CreateNoteRequest

	if err := ginCtx.ShouldBindJSON(&request); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().UTC()

	newNoteToCreate := Note{
		ID:        primitive.NewObjectID(),
		Title:     request.Title,
		Content:   request.Content,
		Pinned:    request.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}

	newlyCreatedNote, err := handler.repo.CreateNote(ginCtx.Request.Context(), newNoteToCreate)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ginCtx.JSON(http.StatusCreated, newlyCreatedNote)
}
