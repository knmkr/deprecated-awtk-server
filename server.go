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

func doRunServer(c *cli.Context) {
	addr := c.String("addr")
	if addr == "" {
		addr = "localhost:1323"
	}

	log.WithFields(log.Fields{"addr": addr, "wgx_version": Version}).Info("Running wgx server")

	e := echo.New()
	e.Use(middleware.Logger())

	// e.GET("/", Redirect())
	// e.GET("/v1/genomes/:genome_id", getGenomes)
	e.GET("/v1/genomes/:genome_id/genotypes", getGenotypes)
	e.Run(standard.New(addr))
}

// func postGenomes(c echo.Context) error {
// }

// func getGenomes(c echo.Context) error {
// }

func getGenotypes(c echo.Context) error {
	// TODO: get genome id
	// id := c.Param("genome_id")
	fileName := "test/data/test.vcf42.vcf.gz"

	// TODO: ?ids=<snp_id, ...>

	queries := strings.Split(c.QueryParam("locations"), ",")

	var locs []wgx.Location
	for i := range queries {
		q := strings.Split(queries[i], "-")
		if len(q) != 2 {
			return c.String(http.StatusBadRequest, "")
		}

		pos, err := strconv.Atoi(q[1])
		if err != nil {
			return c.String(http.StatusBadRequest, "")
		}
		loc := wgx.NewLocation(q[0], pos-1, pos) // 1-based to 0-based
		locs = append(locs, loc)
	}

	record, err := wgx.QueryGenotypes(fileName, locs)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	return c.String(http.StatusOK, string(record))
}
