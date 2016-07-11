package main

import (
	"net/http"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/knmkr/wgx/lib"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func doRunServer(c *cli.Context) error {
	addr := c.String("addr")
	if addr == "" {
		addr = "localhost:1323"
	}

	log.WithFields(log.Fields{"addr": addr, "wgx_version": Version}).Info("Running wgx server")

	e := echo.New()
	e.Use(middleware.Logger())

	// e.GET("/", Redirect())
	e.GET("/v1/genomes/:genome_id/genotypes", getGenotypes)
	e.Run(standard.New(addr))

	return nil
}

func getGenotypes(c echo.Context) error {
	// id := c.Param("genome_id")
	fileName := "test/data/test.vcf42.vcf.gz"

	queries := strings.Split(c.QueryParam("locations"), ",")

	q := strings.Split(queries[0], ":")

	s, err := strconv.Atoi(q[1])
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	e, err := strconv.Atoi(q[1])
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	location := wgx.NewLocation(q[0], s, e+1)
	record := wgx.QueryGenotypes(fileName, location)

	return c.String(http.StatusOK, string(record))
}
