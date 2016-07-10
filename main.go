package main

import (
	"net/http"

	. "github.com/knmkr/wgx/lib"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	addr := ":1323"

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		record := GetGenotype("test/data/test.vcf42.vcf.gz", "20", "14369", "14370")
		return c.String(http.StatusOK, string(record))
	})
	e.Run(standard.New(addr))
}
