package api

import (
	"context"
	"log"

	"github.com/Sarvesh-shinde23/microservice_gin_framework/config"
	"github.com/gin-gonic/gin"
	"github.com/stytchauth/stytch-go/v13/stytch/consumer/passwords"
)

type LoginRequest struct {
	Email    string
	Password string
}

func Login(c *gin.Context) {
	client := config.InitializeStytch()
	var lr LoginRequest
	c.BindJSON(&lr)

	params := &passwords.AuthenticateParams{
		Email:    lr.Email,
		Password: lr.Password,
	}

	resp, err := client.Passwords.Authenticate(context.Background(), params)
	if err != nil {
		log.Fatalf("error in method call: %v", err)
	}

	log.Println(resp)
}

