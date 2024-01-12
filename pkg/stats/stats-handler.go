package stats

import (
	"cards-110-api/pkg/api"
	"cards-110-api/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	S ServiceI
}

// GetStats @Summary Get the user's stats
// @Description Returns stats for the current user
// @Tags Stats
// @ID get-stats
// @Produce json
// @Security Bearer
// @Success 200 {object} PlayerStats
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /stats [get]
func (h *Handler) GetStats(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the stats from the database
	stats, err := h.S.GetStats(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, stats)
}

// GetStatsForPlayer @Summary Get the stats for a player
// @Description Returns stats for a player
// @Tags Stats
// @ID get-stats-for-player
// @Produce json
// @Param playerId path string true "Player ID"
// @Success 200 {object} PlayerStats
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /stats/{playerId} [get]
func (h *Handler) GetStatsForPlayer(c *gin.Context) {
	// Check the user is correctly authenticated
	_, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the player ID from the request
	playerId := c.Param("playerId")

	// Get the stats from the database
	stats, err := h.S.GetStats(ctx, playerId)
	if err != nil {
		c.JSON(http.StatusOK, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, stats)
}
