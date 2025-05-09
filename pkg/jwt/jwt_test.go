package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func setup() *Handler {
	// Setup
	logger := &mockLogger{}
	jwtHandler := NewJWTHandler(logger)
	return jwtHandler
}

func TestGenerateJWT(t *testing.T) {
	jwtHandler := setup()
	// Test GenerateJWT
	t.Run("GenerateJWT", func(t *testing.T) {
		claims := map[string]interface{}{
			"user_id": 123,
			"exp":     time.Now().Add(time.Hour).Unix(),
		}
		secretKey := []byte("secret")
		token, err := jwtHandler.GenerateJWT(claims, secretKey)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})
}

func TestParseJwtWithSecretKey(t *testing.T) {
	jwtHandler := setup()
	t.Run("TestParseJwtWithSecretKey", func(t *testing.T) {
		// Generate test token
		claims := map[string]interface{}{
			"user": "test@example.com",
			"exp":  time.Now().Add(time.Hour).Unix(),
		}
		secretKey := []byte("secret")
		token, _ := jwtHandler.GenerateJWT(claims, secretKey)

		parsed, err := jwtHandler.ParseJWTWithSecretKey(token, secretKey)
		assert.NoError(t, err)
		assert.NotNil(t, parsed)
	})

	// Test ParseJWT with invalid token
	t.Run("ParseJWT invalid token", func(t *testing.T) {
		_, err := jwtHandler.ParseJWTWithSecretKey("invalid", []byte("secret"))
		assert.Error(t, err)
	})

	// Test ParseJWT with expired token
	t.Run("ParseJWT expired token", func(t *testing.T) {
		claims := map[string]interface{}{
			"user": "test@example.com",
			"exp":  time.Now().Add(-time.Hour).Unix(),
		}
		secretKey := []byte("secret")
		token, _ := jwtHandler.GenerateJWT(claims, secretKey)

		_, err := jwtHandler.ParseJWTWithSecretKey(token, secretKey)
		assert.Error(t, err)
	})

	// Test ParseJWT with wrong secret key
	t.Run("ParseJWT wrong secret key", func(t *testing.T) {
		claims := map[string]interface{}{
			"user": "test@example.com",
			"exp":  time.Now().Add(-time.Hour).Unix(),
		}
		secretKey := []byte("secret")
		token, _ := jwtHandler.GenerateJWT(claims, secretKey)

		_, err := jwtHandler.ParseJWTWithSecretKey(token, []byte("wrong-secret"))
		assert.Error(t, err)
	})
}

// func TestParseJwtWithJWKSet(t *testing.T) {
// 	jwtHandler := setup()
// 	// Test ParseJWT
// 	t.Run("ParseJWTWithJWKSet", func(t *testing.T) {
// 		// Generate test token
// 		claims := map[string]interface{}{
// 			"user": "test@example.com",
// 			"exp":  time.Now().Add(time.Hour).Unix(),
// 		}
// 		secretKey := []byte("secret")
// 		token, _ := jwtHandler.GenerateJWT(claims, secretKey)

// 		// Mock JWKS with kid
// 		jwks := `{"keys": [{"kty": "oct", "k": "test", "kid": "test-key"}]}`

// 		parsed, err := jwtHandler.ParseJWTWithJWKSet(token, jwks)
// 		assert.NoError(t, err)
// 		assert.NotNil(t, parsed)
// 	})
// }

func TestParseJwtUnverified(t *testing.T) {
	jwtHandler := setup()
	// Test ParseJWT
	t.Run("TestParseJwtUnverified", func(t *testing.T) {
		// Generate test token
		claims := map[string]interface{}{
			"user": "test@example.com",
			"exp":  time.Now().Add(time.Hour).Unix(),
		}
		secretKey := []byte("secret")
		token, _ := jwtHandler.GenerateJWT(claims, secretKey)

		parsed, err := jwtHandler.ParseJWTUnverified(token)
		assert.NoError(t, err)
		assert.NotNil(t, parsed)
	})
}

func TestGetClaimsValuesByKey(t *testing.T) {
	jwtHandler := setup()
	// Test GetClaimValue
	t.Run("GetClaimValue", func(t *testing.T) {
		claims := map[string]interface{}{
			"user": "test@example.com",
			"exp":  time.Now().Add(time.Hour).Unix(),
		}
		secretKey := []byte("secret")
		token, _ := jwtHandler.GenerateJWT(claims, secretKey)

		parsed, _ := jwtHandler.ParseJWTWithSecretKey(token, secretKey)
		value := parsed.Claims.(jwt.MapClaims)["user"]
		assert.Equal(t, "test@example.com", value)
	})
}

type mockLogger struct{}

func (m *mockLogger) Error(...interface{}) {}
func (m *mockLogger) Info(...interface{})  {}
func (m *mockLogger) Debug(...interface{}) {}
func (m *mockLogger) Warn(...interface{})  {}
func (m *mockLogger) Fatal(...interface{}) {}
func (m *mockLogger) Flush()               {}
