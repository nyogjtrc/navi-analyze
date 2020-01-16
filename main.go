package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	client "github.com/influxdata/influxdb1-client/v2"
)

// ReqNavi _
type ReqNavi struct {
	Start            int64                  `json:"start"`
	NavigationTiming map[string]interface{} `json:"navigation_timing"`
}

func navi(c *gin.Context) {

	tags := map[string]string{
		"ip":   c.ClientIP(),
		"path": c.FullPath(),
	}

	req := ReqNavi{
		NavigationTiming: map[string]interface{}{},
	}

	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
	}

	timeNavi := time.Unix(req.Start/1000, (req.Start%1000)*1000000)

	point, err := client.NewPoint("navi", tags, req.NavigationTiming, timeNavi)
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

	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		c.AbortWithError(500, err)
	}
	defer cli.Close()
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

	r.StaticFile("/", "./dist")

	log.Fatal(r.Run())
}
