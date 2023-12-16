package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PageData struct {
	YouTubeID    string
	ThumbnailURL string
}

func main() {
	e := echo.New()

	// Route to handle the redirect
	e.GET("/:videoID", func(c echo.Context) error {
		videoID := c.Param("videoID")
		pageData := PageData{
			YouTubeID:    videoID,
			ThumbnailURL: "http://img.youtube.com/vi/" + videoID + "/maxresdefault.jpg",
		}

		return c.Render(http.StatusOK, "redirect.html", pageData)
	})

	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Logger.Fatal(e.Start(":8080"))
}

// TemplateRenderer is a custom HTML template renderer
type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
