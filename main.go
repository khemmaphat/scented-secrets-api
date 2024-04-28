package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/khemmaphat/scented-secrets-api/handler"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./config/scented-secrets-1958e-firebase-adminsdk-rhyyj-18af5c9d54.json")

	// Create a new Firestore client.
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "scented-secrets-1958e")
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://scented-secrets-1958e.web.app"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE"}
	config.AllowHeaders = []string{"*"}

	log.Print("Starting the service")

	// Get the port from the environment variable or use a default port (e.g., 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.Use(cors.New(config))

	handler.ApplyUserHandler(r, client)
	handler.ApplyPerfumeHandler(r, client)
	handler.ApplyQuestionHandler(r, client)

	// Add "Hello, World!" response to the root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	log.Printf("The service is ready to listen and serve on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
