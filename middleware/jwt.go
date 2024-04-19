package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PhuPhuoc/hrm-v1/common"
	"github.com/joho/godotenv"
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

func CreateJWT(payload map[string]interface{}) (string, error) {

	err_env := godotenv.Load(".env")
	if err_env != nil {
		log.Fatal("(From jwt_create) Error loading .env file:", err_env)
	}
	secret := os.Getenv("SECRET_KEY")

	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}

	encodedHeader := base64.URLEncoding.EncodeToString([]byte(toJSON(header)))
	encodedPayload := base64.URLEncoding.EncodeToString([]byte(toJSON(payload)))

	signature := computeSignature(encodedHeader, encodedPayload, secret)

	return strings.Join([]string{encodedHeader, encodedPayload, signature}, "."), nil
}

func verifyJWT(token string) (map[string]interface{}, error) {

	err_env := godotenv.Load(".env")
	if err_env != nil {
		log.Fatal("(From jwt_verify) Error loading .env file:", err_env)
	}
	secret := os.Getenv("SECRET_KEY")

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

func ValidateTokenMiddleware(next http.Handler) http.Handler {
	excludedURIs := map[string]bool{
		"/api/v1/register": true,
		"/api/v1/login":    true,
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if excludedURIs[r.URL.Path] || strings.HasPrefix(r.URL.Path, "/swagger") {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			common.WriteJSON(w, common.ErrorResponse_Unauthorized())
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token_payload, err := verifyJWT(tokenString)
		if err != nil {
			common.WriteJSON(w, common.ErrorResponse_Unauthorized())
			return
		}

		if value, ok := token_payload["exp_date"]; ok {
			currentUnixTime := time.Now().Unix()
			expDateFloat, ok_int := value.(float64)
			if !ok_int {
				common.WriteJSON(w, common.ErrorResponse_Unauthorized())
				return
			}
			expDateUnix := int64(expDateFloat)
			if currentUnixTime < expDateUnix {
				next.ServeHTTP(w, r)
			} else {
				common.WriteJSON(w, common.ErrorResponse_TokenExpired())
				return
			}
		} else {
			common.WriteJSON(w, common.ErrorResponse_Unauthorized())
			return
		}
	})
}
