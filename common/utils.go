package common

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

// handle write json when response to client
func WriteJSON(wr http.ResponseWriter, response any) error {
	wr.Header().Set("Content-Type", "application/json")

	if res, ok := response.(*error_response); ok {
		wr.WriteHeader(res.StatusCode)
	}

	if _, ok := response.(*success_response); ok {
		wr.WriteHeader(http.StatusOK)
	}
	return json.NewEncoder(wr).Encode(response)
}

// create hash password
func GenerateHash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

// check account password is match vs passwordhash or not
func CompareHash(password string, hash string) bool {
	newPasswordHash := GenerateHash(password)
	return newPasswordHash == hash
}
