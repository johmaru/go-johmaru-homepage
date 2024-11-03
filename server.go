package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		cookie, err := c.Cookie("jsEnabled")
		if err == nil && cookie.Value == "false" {
			return c.Redirect(http.StatusFound, "/noscript-detected")
		}
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.POST("/adblock-detected", func(c echo.Context) error {
		fmt.Println("Adblock detected")
		return c.NoContent(http.StatusOK)
	})

	e.GET("/noscript-detected", func(c echo.Context) error {
		return c.String(http.StatusOK, "JavaScript is disabled. Please enable JavaScript to continue.")
	})

	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Logger.Fatal(e.Start(":1323"))
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
