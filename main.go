package main

import (
	"net/http"

	_ "leishi1313/Shifter/statik"

	"github.com/labstack/echo"
	"github.com/rakyll/statik/fs"
)

func main() {
	e := echo.New()
	addr := ":8000"

	statikFS, err := fs.New()
	if err != nil {
		e.Logger.Fatal(err)
	}
	h := http.FileServer(statikFS)

	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", h)))

	e.Logger.Fatal(e.Start(addr))
}
