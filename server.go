package main

import (
	"io"
	"kasir/controller"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type M map[string]interface{}

type TemplateRegistry struct {
	templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// const CSRFTokenHeader = "X-CSRF-Token"
	// const CSRFKey = "csrf"

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:" + CSRFTokenHeader,
	// 	ContextKey:  CSRFKey,
	// }))

	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Welcome mvc app using GOLANG")
	})

	//start of diamond route

	e.POST("/insertdiamond", echo.WrapHandler(
		http.HandlerFunc(controller.RouteInsertDiamond),
	))

	e.GET("/adddiamond", echo.WrapHandler(
		http.HandlerFunc(controller.AddDiamondView),
	))

	e.GET("/getalldiamonds", echo.WrapHandler(
		http.HandlerFunc(controller.GetAllDiamonds),
	))

	e.GET("/diamond", echo.WrapHandler(
		http.HandlerFunc(controller.DiamondView),
	))

	e.POST("/deletediamondbyid", echo.WrapHandler(
		http.HandlerFunc(controller.DeleteDiamondById),
	))

	e.POST("/updatediamondprocess", echo.WrapHandler(
		http.HandlerFunc(controller.UpdateDiamondById),
	))

	//end of diamond route

	//start of kurs route
	e.POST("/insertkurs", echo.WrapHandler(
		http.HandlerFunc(controller.RouteInsertKurs),
	))

	e.GET("/getkurs", echo.WrapHandler(
		http.HandlerFunc(controller.GetKurs),
	))
	//end of kurs route

	e.Static("/static", "assets")

	e.Logger.Fatal(e.Start(":8082"))
}
