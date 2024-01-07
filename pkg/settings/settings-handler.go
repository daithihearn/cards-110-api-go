package settings

import (
	"cards-110-api/pkg/api"
	"cards-110-api/pkg/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	S ServiceI
}

// GetSettings @Summary Get the user's settings
// @Description Returns the user's settings.
// @Tags Settings
// @ID get-settings
// @Produce json
// @Security Bearer
// @Success 200 {object} Settings
// @Failure 400 {object} api.ErrorResponse
// @Failure 404 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /settings [get]
func (h *Handler) GetSettings(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the user from the database
	settings, has, err := h.S.Get(ctx, id)

	if err != nil {
		c.JSON(http.StatusOK, api.ErrorResponse{Message: err.Error()})
		return
	}
	if !has {
		log.Printf("No settings found for user, creating new settings")
		settings = Settings{
			ID:      id,
			AutoBuy: true,
		}
	}

	c.IndentedJSON(http.StatusOK, settings)
}

// SaveSettings @Summary Save the user's settings
// @Description Saves the user's settings.
// @Tags Settings
// @ID save-settings
// @Accept json
// @Produce json
// @Security Bearer
// @Param settings body Settings true "Settings"
// @Success 200 {object} Settings
// @Failure 400 {object} api.ErrorResponse
// @Failure 404 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /settings [put]
func (h *Handler) SaveSettings(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the user from the database
	var settings Settings
	if err := c.ShouldBindJSON(&settings); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	settings.ID = id

	log.Printf("Saving settings for user %s", id)

	if err := h.S.Save(ctx, settings); err != nil {
		c.JSON(http.StatusOK, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, settings)
}
