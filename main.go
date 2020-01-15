package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	client "github.com/influxdata/influxdb1-client/v2"
)

func navi(c *gin.Context) {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		c.AbortWithError(500, err)
	}
	defer cli.Close()

	tags := map[string]string{
		"ip":   "127.0.0.1",
		"host": "localhost",
		"path": "/index.html",
	}
	fields := map[string]interface{}{
		"navigationStart": 1579093839831,
		"loadEventEnd":    1579093841269,
	}

	point, err := client.NewPoint("navi", tags, fields, time.Now())
	if err != nil {
		c.AbortWithError(500, err)
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: "db0",
	})
	if err != nil {
		c.AbortWithError(500, err)
	}

	bp.AddPoint(point)

	err = cli.Write(bp)
	if err != nil {
		c.AbortWithError(500, err)
	}

	c.JSON(204, nil)
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/navi", navi)

	log.Fatal(r.Run())
}
