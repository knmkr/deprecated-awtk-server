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
	e.POST("/v1/genomes", postGenomes)
	e.GET("/v1/genomes", getGenomes)
	// e.GET("/v1/genomes/:genome_id", getGenome)
	e.GET("/v1/genomes/:genome_id/genotypes", getGenotypes)

	e.Run(standard.New(addr))
}

// $ curl -X POST --data "filePath=/path/to/genome.vcf.gz" "http://localhost:1323/v1/genomes"
func postGenomes(c echo.Context) error {
	g := new(wgx.Genome)
	if err := c.Bind(g); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	genomes, err := wgx.CreateGenomes(g.FilePath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, genomes)
}

func getGenomes(c echo.Context) error {
	// TODO: get all genomes from db
	return c.JSON(http.StatusOK, "")
}

// func getGenome(c echo.Context) error {
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
