package api

import (
	"context"
	"log"

	"github.com/Sarvesh-shinde23/microservice_gin_framework/config"
	"github.com/gin-gonic/gin"
	"github.com/stytchauth/stytch-go/v13/stytch/consumer/passwords"
)

type RegisterRequest struct {
	Email    string
	Password string
}

func Register(c *gin.Context) {
	client := config.InitializeStytch()
	var rr RegisterRequest
	c.BindJSON(&rr)

	params := &passwords.CreateParams{
		Email:    rr.Email,
		Password: rr.Password,
	}

	resp, err := client.Passwords.Create(context.Background(), params)
	if err != nil {
		log.Fatalf("error in method call: %v", err)
	}

	log.Println(resp)
}
