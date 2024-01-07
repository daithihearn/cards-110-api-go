// @title Cards 110 API
// @version 8.0.0
// @description An API for playing the card game called 110. 110 is a game based on the game 25 and is played primarily in Ireland
// @BasePath /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
package main

import (
	_ "cards-110-api/docs"
	"cards-110-api/pkg/auth"
	"cards-110-api/pkg/db"
	"cards-110-api/pkg/settings"
	"cards-110-api/pkg/user"
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	// Load .env file if it exists
	_ = godotenv.Load()
}

func main() {
	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancel() // Cancel the context upon receiving the signal

		// Create a new context for the graceful shutdown procedure
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		// Gracefully close the database connection
		if err := db.CloseMongoConnection(shutdownCtx); err != nil {
			// Handle error (e.g., log it)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	// Get the db name and collection name
	dbName := os.Getenv("MONGODB_DB")
	if dbName == "" {
		dbName = "cards-110"
	}

	// Configure collections
	userCol, err := db.GetCollection(ctx, dbName, "appUsers")
	if err != nil {
		cancel()
		log.Fatal("Failed to get appUser collection: ", err)
	}
	//deckCol, err := db.GetCollection(ctx, dbName, "decks")
	//if err != nil {
	//	cancel()
	//	log.Fatal("Failed to get deck collection: ", err)
	//}
	settingsCol, err := db.GetCollection(ctx, dbName, "playerSettings")
	if err != nil {
		cancel()
		log.Fatal("Failed to get playerSettings collection: ", err)
	}

	// Configure services
	userColRec := db.Collection[user.User]{Col: userCol}
	userService := user.Service{Col: userColRec}
	userHandler := user.Handler{S: &userService}
	//deckColRec := deck.Collection{Col: deckCol}
	//deckService := deck.Service{Col: deckColRec}
	settingsColRec := db.Collection[settings.Settings]{Col: settingsCol}
	settingsService := settings.Service{Col: settingsColRec}
	settingsHandler := settings.Handler{S: &settingsService}

	// Set up the API routes.
	router := gin.Default()

	// Configure CORS with custom settings
	// Get the environment variable
	origins := os.Getenv("CORS_ALLOWED_ORIGINS")

	// Check if the environment variable is empty and set a default value
	if origins == "" {
		origins = "http://localhost:888,http://localhost:3000" // Replace with your default list
	}

	config := cors.Config{
		AllowOrigins:  strings.Split(origins, ","),
		AllowMethods:  []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}
	router.Use(cors.New(config))

	// Redirect from root to /swagger/index.html
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	// Configure the routes
	router.GET("/api/v1/profile/has", auth.EnsureValidTokenGin([]string{auth.ReadGame}), userHandler.HasProfile)
	router.GET("/api/v1/profile", auth.EnsureValidTokenGin([]string{auth.ReadGame}), userHandler.GetProfile)
	router.PUT("/api/v1/profile", auth.EnsureValidTokenGin([]string{auth.ReadGame}), userHandler.UpdateProfile)
	router.GET("/api/v1/settings", auth.EnsureValidTokenGin([]string{auth.ReadGame}), settingsHandler.GetSettings)
	router.PUT("/api/v1/settings", auth.EnsureValidTokenGin([]string{auth.ReadGame}), settingsHandler.SaveSettings)

	// Use the generated docs in the docs package.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = router.Run(":" + port)
	if err != nil {
		return
	}

	// Wait for the cancellation of the context (due to signal handling)
	<-ctx.Done()
}
