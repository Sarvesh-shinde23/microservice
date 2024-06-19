package config

import (
	"log"

	"github.com/stytchauth/stytch-go/v13/stytch/consumer/stytchapi"
)

// InitializeStytch initializes and returns the Stytch client.
func InitializeStytch() *stytchapi.API {
	client, err := stytchapi.NewClient(
		"project-live-c60c0abe-c25a-4472-a9ed-320c6667d317",
		"secret-live-80JASucyk7z_G8Z-7dVwZVGXL5NT_qGAQ2I",
	)
	if err != nil {
		log.Fatalf("error initializing Stytch client: %v", err)
	}
	return client
}
