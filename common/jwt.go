package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type JWT struct {
	Header    map[string]interface{}
	Payload   map[string]interface{}
	Signature string
}

func toJSON(data interface{}) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func computeSignature(header, payload, secret string) string {
	data := strings.Join([]string{header, payload}, ".")
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func CreateJWT(payload map[string]interface{}, secret string) (string, error) {
	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}

	encodedHeader := base64.URLEncoding.EncodeToString([]byte(toJSON(header)))
	encodedPayload := base64.URLEncoding.EncodeToString([]byte(toJSON(payload)))

	signature := computeSignature(encodedHeader, encodedPayload, secret)

	return strings.Join([]string{encodedHeader, encodedPayload, signature}, "."), nil
}

func verifyJWT(token, secret string) (map[string]interface{}, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	signature := computeSignature(parts[0], parts[1], secret)
	if signature != parts[2] {
		return nil, fmt.Errorf("invalid signature")
	}

	var payload map[string]interface{}
	decoded, err := base64.URLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(decoded, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func ValidateTokenFromRequest(handlerFunc http.HandlerFunc) http.HandlerFunc {
	log.Printf("calling this middleware when router is called...")

	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("calling JWT auth middleware...")
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			WriteJSON(rw, ErrorResponse_Unauthorized())
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		_, err := verifyJWT(tokenString, "hehelalamama")
		if err != nil {
			WriteJSON(rw, ErrorResponse_Unauthorized())
		}
		handlerFunc(rw, r)
	}
}
