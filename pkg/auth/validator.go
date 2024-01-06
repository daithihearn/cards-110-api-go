package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/golang-jwt/jwt/v4"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	validator.CustomClaims
	validator.RegisteredClaims
	Scope string `json:"scope"`
}

// Valid implements the jwt.Claims interface
func (c CustomClaims) Valid() error {
	// Here you can add additional validation for your claims if needed.
	// For example, you can check if the scope is valid, if the token has expired based on a custom expiry claim, etc.
	// If everything is okay, return nil. If there is an issue with the claims, return an error.
	return nil
}

func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidTokenGin is a Gin middleware that will check the validity of our JWT.
func EnsureValidTokenGin(scopes []string) gin.HandlerFunc {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)

	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(func(w http.ResponseWriter, r *http.Request, err error) {
			// Log the error
			log.Printf("Encountered error while validating JWT: %v", err)
			// Don't write to the response here, just log the error
		}),
	)

	return func(c *gin.Context) {
		// Wrap the response writer so we can inspect the status code
		rw := newResponseWriter(c.Writer)
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// This is where your validation takes place
			c.Request = r
			c.Next()
		})

		// Now run the middleware
		handler := middleware.CheckJWT(next)

		// Extract the Authorization header from the request
		authHeader := c.GetHeader("Authorization")
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			// Handle the error - the token is not in the correct format
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}

		// Parse the token
		tokenString := splitToken[1]
		token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &CustomClaims{})
		if err != nil {
			// Handle the error - the token could not be parsed
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Check the token's claims are valid and set them into the Gin context
		if claims, ok := token.Claims.(*CustomClaims); ok {

			// Validate the scopes
			err := validateScopes(claims, scopes)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}

			// Check if the token has the required scopes
			c.Set("user", claims)
		} else {
			// Handle the error - the claims are not of type *CustomClaims
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}
		handler.ServeHTTP(rw, c.Request)

	}
}

func validateScopes(claims *CustomClaims, scopes []string) error {
	for _, scope := range scopes {
		if !claims.HasScope(scope) {
			return jwt.ValidationError{Inner: nil, Errors: jwt.ValidationErrorClaimsInvalid}
		}
	}

	return nil
}

// newResponseWriter creates a new response writer to capture the status code
func newResponseWriter(w gin.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

type responseWriter struct {
	gin.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
