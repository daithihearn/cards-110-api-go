package profile

import (
	"cards-110-api/pkg/api"
	"cards-110-api/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	S ServiceI
}

// Get @Summary Get the user's profile
// @Description Returns the user's profile.
// @Tags Profile
// @ID get-profile
// @Produce json
// @Security Bearer
// @Param playerId query string false "Player ID"
// @Success 200 {object} Profile
// @Failure 400 {object} api.ErrorResponse
// @Failure 404 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /profile [get]
func (h *Handler) Get(c *gin.Context) {
	// Check the user is correctly authenticated
	authID, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Retrieve the optional playerId parameter
	playerId := c.Query("playerId")

	// Use the provided playerId if it exists, otherwise use the authenticated user's ID
	id := authID
	if playerId != "" {
		id = playerId
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the user from the database
	p, has, err := h.S.Get(ctx, id)

	if err != nil {
		c.JSON(http.StatusOK, api.ErrorResponse{Message: err.Error()})
		return
	}
	if !has {
		c.JSON(http.StatusNotFound, api.ErrorResponse{Message: "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, p)
}

// Update @Summary Update the user's profile
// @Description Updates the user's profile or creates it if it doesn't exist.
// @Tags Profile
// @ID update-profile
// @Produce json
// @Accept json
// @Param profile body UpdateProfileRequest true "Profile"
// @Security Bearer
// @Success 200 {object} Profile
// @Failure 400 {object} api.ErrorResponse
// @Failure 404 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /profile [put]
func (h *Handler) Update(c *gin.Context) {
	// Check the user is correctly authenticated
	id, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the request body
	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	// Update the profile
	p, exists, err := h.S.Get(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	if !exists {
		p = Profile{ID: id}
	}
	p.Name = req.Name
	if req.ForceUpdate || !p.PictureLocked {
		p.Picture = req.Picture
	}

	err = h.S.Save(ctx, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, p)
}

// GetAll @Summary Get all profiles
// @Description Returns all profiles.
// @Tags Profile
// @ID get-all-profiles
// @Produce json
// @Security Bearer
// @Success 200 {object} []Profile
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /profile/all [get]
func (h *Handler) GetAll(c *gin.Context) {
	// Check the user is correctly authenticated
	_, ok := auth.CheckValidated(c)
	if !ok {
		return
	}

	// Get the context from the request
	ctx := c.Request.Context()

	// Get all profiles
	profiles, err := h.S.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, profiles)
}
