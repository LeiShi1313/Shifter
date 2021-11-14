package route

import (
	"net/http"

	_ "github.com/leishi1313/Shifter/statik"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo"
	echologrus "github.com/plutov/echo-logrus"
	"github.com/rakyll/statik/fs"

	echoMw "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	echologrus.Logger = logrus.New()
	e.Use(echologrus.Hook())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	statikFS, err := fs.New()
	if err != nil {
		e.Logger.Fatal(err)
	}
	h := http.FileServer(statikFS)

	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", h)))

	return e
}
