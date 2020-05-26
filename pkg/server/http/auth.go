package http

import (
	"context"
	"time"

	"github.com/AnkushJadhav/kamaji-root/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	jwtware "github.com/gofiber/jwt"
)

// EnableJWTAuthentication enables JWT authentication on all routes added to the server
// after calling this function and persists the secret in the store
func (srv *Server) EnableJWTAuthentication() error {
	jwtSecret := generateJWTSecret()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.config.StorageDriver.SetJWTToken(ctx, jwtSecret); err != nil {
		return err
	}
	srv.config.JWTSecret = jwtSecret
	srv.app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		// Signing method should always be explicitly mentioned
		SigningMethod: jwt.SigningMethodHS256.Alg(),
	}))

	return nil
}

func generateJWTSecret() string {
	return utils.GenerateUUID()
}

func (srv *Server) getJWTForUser(id string) (string, error) {
	claims := jwt.StandardClaims{
		Issuer:    "kamaji-root",
		Audience:  "kamaji-root",
		IssuedAt:  time.Now().Unix(),
		Subject:   id,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(srv.config.JWTSecret))
}
