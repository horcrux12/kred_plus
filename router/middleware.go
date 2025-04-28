package router

import (
	"kredi-plus.com/be/dto/out"
	"kredi-plus.com/be/lib/constanta"
	"kredi-plus.com/be/lib/exception"
	"kredi-plus.com/be/lib/helper"
	"kredi-plus.com/be/service"
	"net/http"
	"strings"
)

func PublicMiddleware(next http.Handler) http.Handler {
	return handlerStandardMiddleware(next)
}

func PrivateMiddleware(next http.Handler) http.Handler {
	return handlerStandardMiddleware(handlePrivateMiddleware(next))
}

func handlePrivateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx         = r.Context()
			authService = service.NewAuth()
		)

		authorization := r.Header.Get(constanta.HeaderAuthorizationConstanta)
		if strings.TrimSpace(authorization) == "" {
			helper.WriteErrorResponse(w, exception.ForbiddenAccess, nil)
			return
		}

		// todo check JWT
		splitToken := strings.Split(authorization, "Bearer ")

		if len(splitToken) < 2 {
			helper.WriteErrorResponse(w, exception.ForbiddenAccess, nil)
			return
		}

		token := splitToken[1]
		claims, err := helper.ValidateJWT(token)
		if err != nil {
			helper.WriteErrorResponse(w, exception.ForbiddenAccess, nil)
			return
		}

		userSession, err := authService.GetUserById(ctx, claims.UserID)
		if err != nil {
			helper.WriteErrorResponse(w, err, nil)
			return
		}

		ctx = helper.SetAuthSessionModel(ctx, userSession)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handlerStandardMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// handle CORS Origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept, authorization, x-nextoken, CLIENT_ID, RESOURCE_USER_ID")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
		w.Header().Set("Access-Control-Max-Age", "1209600")

		defer func() {
			if rec := recover(); rec != nil {
				helper.WriteToResponseBody(w, out.WebResponse{
					Status: out.WebStatus{
						Message: "Internal Server Error",
						Detail:  rec,
					},
				}, http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
