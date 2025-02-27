package router

import (
	"asira_borrower/groups"
	"asira_borrower/handlers"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// ignore /api-borrower
	e.Pre(middleware.Rewrite(map[string]string{
		"/api-borrower/*":       "/$1",
		"/api-borrower-devel/*": "/$1",
	}))

	// e.GET("/test", handlers.Test)
	e.GET("/clientauth", handlers.ClientLogin)
	e.GET("/ping", handlers.Ping)

	// files url
	gopath, _ := os.Getwd()
	e.Static("/", gopath+"/assets")

	groups.AdminGroup(e)
	groups.ClientGroup(e)
	groups.BorrowerGroup(e)
	groups.UnverifiedBorrowerGroup(e)

	return e
}
