package users_test

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"user/pkg/middlewares"
)

var (
	privateKey      *rsa.PrivateKey = generateRSAKey()
	wrongPrivateKey *rsa.PrivateKey = generateRSAKey()
	publicKey       *rsa.PublicKey  = &privateKey.PublicKey

	// Handler with MiddlewareJWT
	// Looks ugly but it's ok.
	handler http.Handler = users.MiddlwareJWT(publicKey)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
)

func createToken(claims jwt.MapClaims, secret *rsa.PrivateKey) string {
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(secret)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return token
}

func TestMiddlewareJWT(t *testing.T) {

	testCases := []struct {
		desc  string
		token string
		want  int // http status code

	}{
		{
			desc:  "valid token",
			token: createToken(jwt.MapClaims{"id": "1", "exp": time.Now().Add(time.Hour).Unix()}, privateKey),
			want:  http.StatusOK,
		},
		{
			desc:  "expired token",
			token: createToken(jwt.MapClaims{"id": "1", "exp": time.Now().Add(-time.Hour).Unix()}, privateKey),
			want:  http.StatusUnauthorized,
		},
		{
			desc:  "wrong coded token",
			token: createToken(jwt.MapClaims{"id": "1", "exp": time.Now().Add(time.Hour).Unix()}, wrongPrivateKey),
			want:  http.StatusUnauthorized,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			resp := httptest.NewRecorder()

			req.Header.Set("Authorization", "Bearer "+tC.token)

			handler.ServeHTTP(resp, req)
			if resp.Code != tC.want {
				t.Errorf("%s: got status code %d, want %d", tC.desc, resp.Code, tC.want)
			}
		})
	}
}

func generateRSAKey() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return privateKey
}
