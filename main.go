package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/idtoken"
	"net/http"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "root")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health")
	})

	e.GET("/oauthservice/operation", func(c echo.Context) error {
		e.Logger.Print("operation")
		token := ""

		for k, v := range c.Request().Header {
			for i := range v {
				if k == "X-Goog-Iap-Jwt-Assertion" {
					token = v[i]
				}
				e.Logger.Printf("key :%s --> value : %s \n", k, v[i])
			}
		}

		audience := "63785118804-tkq80964jqq266r4mufm6kl9u40an90c.apps.googleusercontent.com"
		validate, err := idtoken.Validate(c.Request().Context(), token, audience)
		if err != nil {
			e.Logger.Print(err)
			return err
		}
		e.Logger.Print("validate", validate.Claims)

		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	e.Logger.Fatal(e.Start(":9001"))
}
