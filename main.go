package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	r := gin.Default()
	// get client ip
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "your ip: "+c.ClientIP())
	})

	g := r.Group("/api")
	g.GET("/time", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"time": time.Now().Format(time.DateTime),
		})
	})
	g.GET("/timeSecond", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"time": strconv.Itoa(int(time.Now().Unix())),
		})
	})
	r.GET("/v7/weather/:kind", func(c *gin.Context) {
		location := c.Query("location")
		language := c.Query("lang")
		key := c.Query("key")
		//log.Println(c.FullPath())
		realPath := c.Request.URL.Path

		client := resty.New()
		client.SetDebug(true)
		urlStr := "https://devapi.qweather.com" + realPath
		log.Println("urlStr: ", urlStr)
		resp, err := client.R().SetHeader("X-QW-Api-Key", key).
			SetQueryParams(map[string]string{
				"location": location,
				"lang":     language,
			}).
			Get(urlStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.String(http.StatusOK, resp.String())
	})
	r.Run(":28683")
}
