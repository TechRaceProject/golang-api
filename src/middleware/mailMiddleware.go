package middleware

import (
	"api/src/config"
	"api/src/services"
	"log"

	"github.com/gin-gonic/gin"
)

func SendWelcomeEmailMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Writer.Status() == 201 {
			email, emailExists := c.Get("user_email")
			username, usernameExists := c.Get("username")

			if !emailExists {
				log.Println("user_email not found in context")
				return
			}

			if !usernameExists {
				log.Println("username not found in context")
				return
			}

			log.Printf("Preparing to send welcome email to: %s", email)

			emailRequest := services.EmailRequest{
				To:       email.(string),
				Subject:  "Welcome to TechRace!",
				Template: "welcome_template.html",
				Data: map[string]interface{}{
					"Username": username.(string),
				},
			}

			err := services.SendEmail(emailRequest, cfg)
			if err != nil {
				log.Printf("Failed to send welcome email: %v", err)
				c.JSON(500, gin.H{"error": "Failed to send welcome email"})
			} else {
				log.Println("Welcome email sent successfully")
			}
		} else {
			log.Printf("Response status is not 201: %d", c.Writer.Status())
		}
	}
}
