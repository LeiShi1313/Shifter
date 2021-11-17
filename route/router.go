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

func useAPIGroup(group *echo.Group) {
	group.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
}

func handleWeb(e *echo.Echo) echo.HandlerFunc {
	statikFS, err := fs.New()
	if err != nil {
		e.Logger.Fatal(err)
	}
	h := http.FileServer(statikFS)

	return echo.WrapHandler(http.StripPrefix("/", h))
}

func Init() *echo.Echo {

	e := echo.New()

	echologrus.Logger = logrus.New()
	e.Use(echologrus.Hook())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	// if username := viper.GetString("web.username"); username != "" {
	// 	if passwordHash := viper.GetString("web.password"); passwordHash != "" {
	// 		e.Use(echoMw.JWTWithConfig(echoMw.JWTConfig{
	// 			SigningKey:  []byte(username + passwordHash),
	// 			TokenLookup: "query:token",
	// 		}))
	// 	}
	// }

	useAPIGroup(e.Group("/api"))
	e.GET("/*", handleWeb(e))

	return e
}
