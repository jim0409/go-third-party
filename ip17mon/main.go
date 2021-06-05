package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	loc := LoadIpip(datPath)
	if err := loc.NewIpipObj(); err != nil {
		panic(err)
	}

	router.GET("/check", func(c *gin.Context) {
		ip := c.Query("ip")
		info, err := loc.QueryIP(ip)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "unable to resolve ip loc!",
			})

			c.Abort()
			return
		}

		c.JSON(200, gin.H{
			"Country": info[0],
			"Region":  info[1],
			"City":    info[2],
			"ISP":     info[3],
		})
	})

	router.Run(":8080")
}
