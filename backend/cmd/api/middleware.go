package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Mateusz2734/wdai-project/backend/internal/response"

	"github.com/pascaldekloe/jwt"
	"github.com/tomasen/realip"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				app.serverError(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) logAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mw := response.NewMetricsResponseWriter(w)
		next.ServeHTTP(mw, r)

		var (
			ip     = realip.FromRequest(r)
			method = r.Method
			url    = r.URL.String()
			proto  = r.Proto
		)

		userAttrs := slog.Group("user", "ip", ip)
		requestAttrs := slog.Group("request", "method", method, "url", url, "proto", proto)
		responseAttrs := slog.Group("response", "status", mw.StatusCode, "size", mw.BytesCount)

		app.logger.Info("access", userAttrs, requestAttrs, responseAttrs)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader != "" {
			headerParts := strings.Split(authorizationHeader, " ")

			if len(headerParts) == 2 && headerParts[0] == "Bearer" {
				token := headerParts[1]

				if app.blacklist.CheckToken(token) {
					app.invalidAuthenticationToken(w, r)
					return
				}

				claims, err := jwt.HMACCheck([]byte(token), []byte(app.config.jwt.secretKey))
				if err != nil {
					app.invalidAuthenticationToken(w, r)
					return
				}

				if !claims.Valid(time.Now()) {
					app.invalidAuthenticationToken(w, r)
					return
				}

				if claims.Issuer != app.config.baseURL {
					app.invalidAuthenticationToken(w, r)
					return
				}

				if !claims.AcceptAudience(app.config.baseURL) {
					app.invalidAuthenticationToken(w, r)
					return
				}

				userID, err := strconv.Atoi(claims.Subject)
				if err != nil {
					app.serverError(w, r, err)
					return
				}

				user, err := app.db.GetUserById(r.Context(), int32(userID))
				if err != nil {
					app.serverError(w, r, err)
					return
				}

				if user != nil {
					r = contextSetAuthenticatedUser(r, user)
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authenticatedUser := contextGetAuthenticatedUser(r)

		if authenticatedUser == nil {
			app.authenticationRequired(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAdminPrivileges(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authenticatedUser := contextGetAuthenticatedUser(r)

		if authenticatedUser == nil {
			app.authenticationRequired(w, r)
			return
		}

		if authenticatedUser.Role != "admin" {
			app.adminPrivilegesRequired(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
