package user

import (
	"cards-110-api/pkg/api"
	"cards-110-api/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	S ServiceI
}

// HasProfile @Summary Check if user has profile
// @Description Returns a boolean indicating if the user has a profile or not.
// @Tags Profile
// @ID has-profile
// @Produce  json
// @Security Bearer
// @Success 200 {object} bool
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /profile/has [get]
func (h *Handler) HasProfile(c *gin.Context) {

	// Extracting the token claims
	user, exists := c.Get("user") // Use the key that your JWT middleware uses to store the user information
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token is missing or has not been validated"})
		return
	}
	tokenClaims := user.(*auth.CustomClaims)

	// Check if the token has the required scope
	if !tokenClaims.HasScope("read:game") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient scope."})
	}

	id := tokenClaims.RegisteredClaims.Subject

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the user from the database
	_, has, err := h.S.GetUser(ctx, id)

	if err != nil {
		c.JSON(http.StatusOK, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, has)
}

// GetProfile @Summary Get the user's profile
// @Description Returns the user's profile.
// @Tags Profile
// @ID get-profile
// @Produce json
// @Security Bearer
// @Success 200 {object} bool
// @Failure 400 {object} api.ErrorResponse
// @Failure 404 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	// Extracting the token claims
	user, exists := c.Get("user") // Use the key that your JWT middleware uses to store the user information
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token is missing or has not been validated"})
		return
	}
	tokenClaims := user.(*auth.CustomClaims)

	// Check if the token has the required scope
	if !tokenClaims.HasScope("read:game") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient scope."})
	}

	id := tokenClaims.RegisteredClaims.Subject

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the user from the database
	p, has, err := h.S.GetUser(ctx, id)

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

// UpdateProfile @Summary Update the user's profile
// @Description Updates the user's profile or creates it if it doesn't exist.
// @Tags Profile
// @ID update-profile
// @Produce json
// @Accept json
// @Param profile body UpdateProfileRequest true "Profile"
// @Security Bearer
// @Success 200 {object} User
// @Failure 400 {object} api.ErrorResponse
// @Failure 404 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /profile [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	// Extracting the token claims
	user, exists := c.Get("user") // Use the key that your JWT middleware uses to store the user information
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token is missing or has not been validated"})
		return
	}
	tokenClaims := user.(*auth.CustomClaims)

	// Check if the token has the required scope
	if !tokenClaims.HasScope("read:game") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient scope."})
	}

	id := tokenClaims.RegisteredClaims.Subject

	// Get the context from the request
	ctx := c.Request.Context()

	// Get the request body
	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	// Update the profile
	u, err := h.S.UpdateUser(ctx, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, u)
}
