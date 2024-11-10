package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/v7/weather/now", func(c *gin.Context) {
		location := c.Query("location")
		language := c.Query("lang")
		key := c.Query("key")

		client := resty.New()
		client.SetDebug(true)
		resp, err := client.R().SetHeader("X-QW-Api-Key", key).
			SetQueryParams(map[string]string{
				"location": location,
				"lang":     language,
			}).
			Get("https://devapi.qweather.com/v7/weather/now")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(resp.Request.URL)
		c.String(http.StatusOK, resp.String())
	})
	r.Run(":28083") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
