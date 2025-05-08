package jwt

import (
	logger "client_server/pkg/logger"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	log logger.Logger
}

func NewJWTHandler(log logger.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) GenerateJWT(claims map[string]interface{}, secretKey []byte) (string, error) {
	mapClaims := jwt.MapClaims{}
	for k, v := range claims {
		mapClaims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (h *Handler) ParseJWTWithSecretKey(tokenString string, secretKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (h *Handler) ParseJWTWithJWKSet(tokenString string, jwks string) (*jwt.Token, error) {
	// Get the JWK Set as JSON.
	jwksJSON := json.RawMessage([]byte(jwks))

	// Create the keyfunc.
	k, err := keyfunc.NewJWKSetJSON(jwksJSON)
	if err != nil {
		h.log.Error(fmt.Sprintf("Failed to create a keyfunc.Keyfunc.\nError: %s", err))
		return nil, err
	}

	// Parse the JWT.
	token, err := jwt.Parse(tokenString, k.Keyfunc)
	if err != nil {
		h.log.Error(fmt.Sprintf("Failed to parse the JWT.\nError: %s", err))
		return nil, err
	}

	// Check if the token is valid.
	if !token.Valid {
		h.log.Error("The token is not valid.")
		return nil, errors.New("the token is not valid")
	}
	return token, nil
}

func (h *Handler) ParseJWTUnverified(tokenString string) (*jwt.Token, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &jwt.MapClaims{})
	if err != nil {
		h.log.Error(fmt.Sprintf("Failed to parse the JWT.\nError: %s", err))
		return nil, err
	}
	return token, nil
}

func (h *Handler) GetClaimsValuesByKey(token *jwt.Token, key string, defaultValue string) interface{} {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		h.log.Info(fmt.Sprintf("token.Claims: %v", claims))
		if claims[key] == nil {
			return defaultValue
		}
		return claims[key]
	}
	return defaultValue
}
