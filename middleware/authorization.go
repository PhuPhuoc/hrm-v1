package middleware

import (
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm-v1/common"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		role := ctx.Value(ContextKeyRole)
		fmt.Println("User Role:", role)
		if role != "ADMIN" {
			common.WriteJSON(w, common.ErrorResponse_NoPermission())
			return
		}
		next.ServeHTTP(w, r)
	})
}
