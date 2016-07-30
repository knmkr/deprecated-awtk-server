package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/knmkr/wgx/lib"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/urfave/cli"
)

func doRunServer(c *cli.Context) {
	wgx.InitDatabase()

	addr := c.String("addr")
	if addr == "" {
		addr = "localhost:1323"
	}

	log.WithFields(log.Fields{"addr": addr, "wgx_version": Version}).Info("Running wgx server")

	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/v1/genomes", postGenomes)
	e.GET("/v1/genomes/:genome_id", getGenomes)
	e.GET("/v1/genomes/:genome_id/genotypes", getGenotypes)
	e.GET("/v1/evidences/:evidence_id", getEvidences)

	e.Run(standard.New(addr))
}

// postGenomes creates genomes records by filePath
// $ curl -X POST --data "filePath=test/data/test.vcf41.vcf.gz" "http://localhost:1323/v1/genomes"
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

// getGenomes returns genome record by id
func getGenomes(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("genome_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	genome, err := wgx.GetGenome(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, genome)
}

// getGenotypes returns genotypes records of genome by locations
// $ curl "localhost:1323/v1/genomes/1/genotypes?locations=20-14370"
func getGenotypes(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("genome_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	genome, err := wgx.GetGenome(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	queries := strings.Split(c.QueryParam("locations"), ",")

	var locs []wgx.Location
	for i := range queries {
		q := strings.Split(queries[i], "-")
		if len(q) != 2 {
			err = &wgx.GenomeError{fmt.Sprintf("%s", "Invalid locations")}
			return c.JSON(http.StatusBadRequest, err)
		}

		pos, err := strconv.Atoi(q[1])
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		loc := wgx.NewLocation(q[0], pos-1, pos) // 1-based to 0-based
		locs = append(locs, loc)
	}

	genotypes, err := wgx.QueryGenotypes(genome.FilePath, genome.SampleIndex, locs)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.String(http.StatusOK, string(genotypes))
}

//
func getEvidences(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("evidence_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	evidence, err := wgx.GetEvidence(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.String(http.StatusOK, string(evidence))
}
