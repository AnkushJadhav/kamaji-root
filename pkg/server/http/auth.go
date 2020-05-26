package http

import (
	"context"
	"time"

	"github.com/AnkushJadhav/kamaji-root/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	jwtware "github.com/gofiber/jwt"
)

const jwtSigningMethod = "HS256"

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
		SigningMethod: jwtSigningMethod,
	}))

	return nil
}

func generateJWTSecret() string {
	return utils.GenerateUUID()
}

func (srv *Server) getJWTForUser(id string) (string, error) {
	claims := jwt.StandardClaims{
		Issuer:   "kamaji-root",
		IssuedAt: time.Now().Unix(),
		Subject:  id,
		
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(jwtSigningMethod), claims)
	return token.SignedString(srv.config.JWTSecret)
}
