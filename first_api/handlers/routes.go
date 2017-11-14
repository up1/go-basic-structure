package handlers

import (
	"net/http"

	"demo/internal/middleware"
	"demo/internal/platform/db"
	"demo/internal/platform/web"
)

func API(masterDB *db.DB, redis) http.Handler {
	// Create the web handler for setting routes and middleware.
	app := web.New(middleware.RequestLogger, middleware.ErrorHandler)

	u := User{masterDB}
	app.Handle("GET", "/v1/users", u.List)
	app.Handle("POST", "/v1/users", u.Create)
	app.Handle("GET", "/v1/users/:id", u.Retrieve)
	app.Handle("PUT", "/v1/users/:id", u.Update)
	app.Handle("DELETE", "/v1/users/:id", u.Delete)

	return app
}
